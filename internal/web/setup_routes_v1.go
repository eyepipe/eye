package web

import (
	"fmt"
	"github.com/eyepipe/eye/internal/lib/crypto2"
	"github.com/eyepipe/eye/internal/pkg/container"

	"github.com/eyepipe/eye/pkg/proto"
	"github.com/gofiber/fiber/v3"
)

func (w *Web) SetupRoutesV1(app *fiber.App) {
	app.Get(proto.SlugV1SuperSchemeSuper, w.RenderConstJSONFn(container.SchemeV1Super))
	app.Get(proto.SlugV1SuperSchemeHigh, w.RenderConstJSONFn(container.SchemeV1High))
	app.Get(proto.SlugV1AlgoSigner, w.RenderConstJSONFn(crypto2.SignerAlgoValues()))
	app.Get(proto.SlugV1KeyAgreement, w.RenderConstJSONFn(crypto2.KeyAgreementAlgoValues()))
	app.Get(proto.SlugV1KeyDerivation, w.RenderConstJSONFn(crypto2.HashKeyDerivationAlgoValues()))
	app.Get(proto.SlugV1BlockCipher, w.RenderConstJSONFn(crypto2.BlockCipherAlgoValues()))
	app.Get("/v1", w.RenderContractJSONFn(proto.ContractV1{
		SchemeURLs:           []string{proto.SlugV1SuperSchemeSuper, proto.SlugV1SuperSchemeHigh},
		SignerAlgoURL:        proto.SlugV1AlgoSigner,
		KeyAgreementAlgoURL:  proto.SlugV1KeyAgreement,
		KeyDerivationAlgoURL: proto.SlugV1KeyDerivation,
		BlockCipherAlgoURL:   proto.SlugV1BlockCipher,
		CreateUploadURL:      proto.SlugV1CreateUpload,
		ConfirmUploadURL:     proto.SlugV1ConfirmUpload,
	}))
	app.Post(proto.SlugV1CreateUpload, w.CreateUpload)
	app.Post(proto.SlugV1ConfirmUpload, w.ConfirmUpload)
	app.Get(fmt.Sprintf("/v1/downloads/:%s", ParamDownloadUUID), w.Download)
}
