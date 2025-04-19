package web

import (
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
	"github.com/eyepipe/eye/internal/pkg/services/export_service"
	"github.com/eyepipe/eye/pkg/proto"
	"github.com/gofiber/fiber/v3"
	"github.com/shlima/oi/null"
)

func (w *Web) CreateUpload(c fiber.Ctx) error {
	signAlgo := crypto2.SignerAlgo(c.Get("x-signer-algo"))
	if !signAlgo.IsValid() {
		return fmt.Errorf("invalid signing algo: <%s>", signAlgo.String())
	}

	reader := io.LimitReader(c.Request().BodyStream(), w.config.GetServerBodyLimitBytes())
	defer func() {
		_ = c.Request().CloseBodyStream()
	}()

	// TODO: ADD VALIDATION ON EMPTY READER
	// TODO: VALIDATE panic(c.Request().IsBodyStream())

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
		TTL:        null.NewAutoTime(time.Now().AddDate(0, 0, 7)),
	}
	err := store.CreateUpload(c.Context(), upload)
	if err != nil {
		return fmt.Errorf("failed to create upload: %w", err)
	}

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
		Token:      token,
		UploadUUID: upload.UUID.String,
	})
}

func (w *Web) ConfirmUpload(c fiber.Ctx) error {
	req := new(proto.ConfirmUploadRequestV1)
	err := c.Bind().JSON(req)
	if err != nil {
		return fmt.Errorf("failed to bind request: %w", err)
	}

	jwt, err := w.jwt.DecodeUploadVerificationJWT(req.Token)
	if err != nil {
		return fmt.Errorf("failed to decode jwt: %w", err)
	}

	jwtUUID, err := uuidv7.Decode(jwt.UploadUUID)
	if err != nil {
		return fmt.Errorf("failed to decode uuid: %w", err)
	}

	store := w.stores.MustGet(jwtUUID.Shard)
	upload, err := store.FindUpload(c.Context(), jwtUUID.String())
	if err != nil {
		return fmt.Errorf("failed to find upload: %w", err)
	}

	pubBytes, err := hex.DecodeString(req.PubKeyHex)
	if err != nil {
		return fmt.Errorf("failed to decode public key: %w", err)
	}

	pubKey, err := export_service.SPKIToPublicKey(pubBytes)
	if err != nil {
		return fmt.Errorf("failed to import public key: %w", err)
	}

	sig, err := hex.DecodeString(req.SigHex)
	if err != nil {
		return fmt.Errorf("failed to decode sig: %w", err)
	}

	serverHash, err := hex.DecodeString(jwt.ServerHashHex)
	if err != nil {
		return fmt.Errorf("failed to decode serverhash: %w", err)
	}

	signer := upload.GetSignerAlgo().ToSigner()
	valid, err := signer.Verify(serverHash, sig, pubKey)
	switch {
	case err != nil:
		return fmt.Errorf("failed to verify signature: %w", err)
	case !valid:
		return fmt.Errorf("signature is invalid")
	}

	err = store.UpdateUploadSignatureHex(c.Context(), jwt.UploadUUID, hex.EncodeToString(sig))
	if err != nil {
		return fmt.Errorf("failed to update upload signature: %w", err)
	}

	return c.JSON(proto.ConfirmUploadResponseV1{
		URL: fmt.Sprintf("%s/v1/downloads/%s", c.BaseURL(), jwtUUID.String()),
	})
}
