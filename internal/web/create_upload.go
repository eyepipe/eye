package web

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"time"

	"github.com/eyepipe/eye/internal/lib/crypto2"
	"github.com/eyepipe/eye/internal/lib/jwt_token/jwt_claims"
	"github.com/eyepipe/eye/internal/lib/s3_cli"
	"github.com/eyepipe/eye/internal/lib/size_reader"
	"github.com/eyepipe/eye/internal/lib/uuidv7"
	"github.com/eyepipe/eye/internal/pkg/domain"
	"github.com/eyepipe/eye/internal/pkg/validator"
	"github.com/eyepipe/eye/pkg/proto"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/shlima/oi/null"
)

// CreateUpload
//
//	curl -X POST -H 'x-signer-algo: ECDSA-P521-SHA512' --data @README.md http://localhost:3000/v1/uploads/create
func (w *Web) CreateUpload(c fiber.Ctx) error {
	signAlgo := crypto2.SignerAlgo(c.Get("x-signer-algo"))
	if !signAlgo.IsValid() {
		return fmt.Errorf("invalid signing algo: <%s>", signAlgo.String())
	}

	// POST without data body
	if !c.Request().IsBodyStream() {
		return fmt.Errorf("expected body as a stream chunk")
	}

	reader := io.LimitReader(c.Request().BodyStream(), w.config.ServerSingleUploadBytesLimit)
	defer func() {
		_ = c.Request().CloseBodyStream()
	}()

	signer := signAlgo.ToSigner()
	hashier := signer.GetHashier().New()
	reader = io.TeeReader(reader, hashier)
	sizedReader := size_reader.New()
	reader = sizedReader.Wrap(reader)

	store, shard := w.stores.Sample()
	s3, _ := w.s3.Sample()
	uuid := uuidv7.NewWithShard(shard)
	key := GenS3KeyByUUID(uuid)

	upload := &domain.Upload{
		UUID:       null.NewAutoString(uuid.String()),
		SignerAlgo: null.NewAutoString(signAlgo.String()),
		S3Key:      null.NewAutoString(key),
		S3Urn:      null.NewAutoString(s3.GetURN()),
		TTL:        null.NewAutoTime(w.now().AddDate(0, 0, 7)),
	}
	err := store.CreateUpload(c.Context(), upload)
	if err != nil {
		return fmt.Errorf("failed to create upload: %w", err)
	}

	limitation, err := store.FetchLimitation(c.Context(), w.now())
	if err != nil {
		return fmt.Errorf("failed to fetch limitation: %w", err)
	}
	err = w.validator.ValidateWriteLimit(c.Context(), validator.ValidateWriteLimitOpts{
		Limitation:        limitation,
		WriteBytesLimit:   w.config.ServerShardWriteBytesLimit,
		WriteCounterLimit: w.config.ServerShardWriteCounterLimit,
	})
	if err != nil {
		return fmt.Errorf("failed to validate write limit: %w", err)
	}

	defer func() {
		w.incrementWriteLimitationSilent(c.Context(), shard, sizedReader.GetByteSize())
	}()

	_, err = s3.UploadReadAll(
		c.Context(),
		reader,
		s3_cli.ACLPrivate,
		key,
	)
	if err != nil {
		return fmt.Errorf("failed to upload to s3: %w", err)
	}

	err = store.UpdateUploadByteSize(c.Context(), upload.UUID.String, sizedReader.GetByteSize())
	if err != nil {
		return fmt.Errorf("failed to update upload byte size: %w", err)
	}

	token, err := w.jwt.GenerateUploadVerificationJWT(&jwt_claims.GenerateUploadVerificationClaims{
		UploadUUID:    upload.UUID.String,
		ServerHashHex: hex.EncodeToString(hashier.Sum(nil)),
	})
	if err != nil {
		return fmt.Errorf("failed to generate jwt: %w", err)
	}

	return c.JSON(proto.CreateUploadResponseV1{
		Token:           token,
		UploadUUID:      upload.UUID.String,
		ConfirmationURL: c.BaseURL() + proto.SlugV1ConfirmUpload,
	})
}

func (w *Web) incrementWriteLimitationSilent(ctx context.Context, shard uint16, byteSize int64) {
	err := w.incrementWriteLimitation(ctx, shard, byteSize)
	if err != nil {
		log.Errorf("failed to increment write limitation: %v", err)
	}
}

func (w *Web) incrementWriteLimitation(ctx context.Context, shard uint16, byteSize int64) error {
	store := w.stores.MustGet(shard)
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := store.IncrementLimitation(ctx, domain.Limitation{
		Date:           null.NewAutoDate(w.now()),
		WrittenBytes:   null.NewAutoInt64(byteSize),
		WrittenCounter: null.NewAutoInt64(1),
	})
	if err != nil {
		return fmt.Errorf("failed to increment limit: %w", err)
	}

	return nil
}
