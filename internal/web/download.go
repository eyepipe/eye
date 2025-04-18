package web

import (
	"fmt"
	"github.com/eyepipe/eye/internal/lib/uuidv7"
	"time"

	"github.com/eyepipe/eye/pkg/proto"
	"github.com/gofiber/fiber/v3"
)

const ParamDownloadUUID = "uuid"

func (w *Web) Download(c fiber.Ctx) error {
	uuid, err := uuidv7.Decode(c.Params(ParamDownloadUUID))
	if err != nil {
		return fmt.Errorf("failed to decode uuid: %w", err)
	}

	shard, err := w.stores.Get(uuid.Shard)
	if err != nil {
		return fmt.Errorf("failed to get database shard: %w", err)
	}

	upload, err := shard.FindNotExpiredUpload(c.Context(), uuid.String(), time.Now())
	if err != nil {
		return fmt.Errorf("failed to find upload: %w", err)
	}

	// TODO: update upload download size
	// TODO: get s3 shard by its name
	url, err := w.s3.MustGet(0).GetPresignedURL(c.Context(), upload.S3Key.String, time.Minute)
	if err != nil {
		return fmt.Errorf("failed to get presigned url: %w", err)
	}

	c.Set(proto.HeaderSigHex, upload.SignatureHex.String)
	return c.Redirect().Status(307).To(url)
}
