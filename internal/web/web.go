package web

import (
	"hash"
	"time"

	"github.com/eyepipe/eye/internal/lib/crypto2"
	"github.com/eyepipe/eye/internal/lib/jwt_token"
	"github.com/eyepipe/eye/internal/lib/s3_cli"
	"github.com/eyepipe/eye/internal/lib/shard"
	"github.com/eyepipe/eye/internal/pkg/configuration"
	"github.com/eyepipe/eye/internal/pkg/validator"
	"github.com/gofiber/fiber/v3"
)

type Web struct {
	app       *fiber.App
	stores    *shard.Shards[IStore]
	s3        *shard.Shards[s3_cli.ICli]
	jwt       jwt_token.IToken
	config    *configuration.Configuration
	validator validator.IValidator
	now       func() time.Time
}

func New() *Web {
	return &Web{
		validator: validator.New(),
		now:       time.Now,
	}
}

func (w *Web) NewHash() hash.Hash {
	return crypto2.HashAlgoSHA256.ToHashier().New()
}

func (w *Web) SetS3(shards *shard.Shards[s3_cli.ICli]) {
	w.s3 = shards
}

func (w *Web) SetStores(stores *shard.Shards[IStore]) {
	w.stores = stores
}

func (w *Web) SetJwtToken(jwt jwt_token.IToken) {
	w.jwt = jwt
}

func (w *Web) SetApp(c *fiber.App) {
	w.app = c
}

func (w *Web) SetConfig(config *configuration.Configuration) {
	w.config = config
}
