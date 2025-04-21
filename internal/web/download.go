package web

import (
	"context"
	"fmt"
	"time"

	"github.com/eyepipe/eye/internal/lib/uuidv7"
	"github.com/eyepipe/eye/internal/pkg/domain"
	"github.com/eyepipe/eye/internal/pkg/validator"
	"github.com/eyepipe/eye/pkg/proto"
	"github.com/gofiber/fiber/v3"
	"github.com/shlima/oi/null"
)

const ParamDownloadUUID = "uuid"

func (w *Web) Download(c fiber.Ctx) error {
	uuid, err := uuidv7.Decode(c.Params(ParamDownloadUUID))
	if err != nil {
		return fmt.Errorf("failed to decode uuid: %w", err)
	}

	shardNum := uuid.Shard
	shard, err := w.stores.Get(shardNum)
	if err != nil {
		return fmt.Errorf("failed to get database shard: %w", err)
	}

	limitation, err := shard.FetchLimitation(c.Context(), w.now())
	if err != nil {
		return fmt.Errorf("failed to fetch limitation: %w", err)
	}
	err = w.validator.ValidateReadLimit(c.Context(), validator.ValidateReadLimitOpts{
		Limitation:       limitation,
		ReadBytesLimit:   w.config.ServerShardReadBytesLimit,
		ReadCounterLimit: w.config.ServerShardReadCounterLimit,
	})
	if err != nil {
		return fmt.Errorf("failed to validate read limit: %w", err)
	}

	upload, err := shard.FindNotExpiredUpload(c.Context(), uuid.String(), time.Now())
	if err != nil {
		return fmt.Errorf("failed to find upload: %w", err)
	}

	// TODO: get s3 shard by its name
	url, err := w.s3.MustGet(0).GetPresignedURL(c.Context(), upload.S3Key.String, time.Minute)
	if err != nil {
		return fmt.Errorf("failed to get presigned url: %w", err)
	}

	err = w.incrementReadLimitation(c.Context(), shardNum, upload.ByteSize.Int64)
	if err != nil {
		return fmt.Errorf("failed to increment read limitation: %w", err)
	}

	c.Set(proto.HeaderSigHex, upload.SignatureHex.String)
	return c.Redirect().Status(307).To(url)
}

func (w *Web) incrementReadLimitation(ctx context.Context, shard uint16, byteSize int64) error {
	store := w.stores.MustGet(shard)
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := store.IncrementLimitation(ctx, domain.Limitation{
		Date:        null.NewAutoDate(w.now()),
		ReadBytes:   null.NewAutoInt64(byteSize),
		ReadCounter: null.NewAutoInt64(1),
	})
	if err != nil {
		return fmt.Errorf("failed to increment limit: %w", err)
	}

	return nil
}
