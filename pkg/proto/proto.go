package proto

import "github.com/samber/lo"

const (
	Copyright = "2025 © https://github.com/eyepipe/eye"
	Usage     = "God sees everything, except what's encrypted."
)

const (
	SlugV1SuperSchemeSuper = "/v1/schemes/super.json"
	SlugV1SuperSchemeHigh  = "/v1/schemes/high.json"
	SlugV1AlgoSigner       = "/v1/algo/signer.json"
	SlugV1KeyAgreement     = "/v1/algo/key-agreement.json"
	SlugV1KeyDerivation    = "/v1/algo/key-derivation.json"
	SlugV1BlockCipher      = "/v1/algo/block-cipher.json"
	SlugV1CreateUpload     = "/v1/uploads/create"
	SlugV1ConfirmUpload    = "/v1/uploads/confirm"
)

const (
	// HeaderSigHex header with digital signature decoded in HEX
	HeaderSigHex     = "x-sig-hex"
	HeaderSignerAlgo = "x-signer-algo"
)

type ContractV1 struct {
	SchemeURLs           []string `json:"scheme_urls"`
	SignerAlgoURL        string   `json:"signer_algo_url"`
	KeyAgreementAlgoURL  string   `json:"key_agreement_algo_url"`
	KeyDerivationAlgoURL string   `json:"key_derivation_algo_url"`
	BlockCipherAlgoURL   string   `json:"block_cipher_algo_url"`
	CreateUploadURLs     []string `json:"create_upload_urls"`
}

func (c ContractV1) WithHost(host string) ContractV1 {
	c.SchemeURLs = lo.Map(c.SchemeURLs, func(x string, index int) string {
		return host + x
	})
	c.SignerAlgoURL = host + c.SignerAlgoURL
	c.KeyAgreementAlgoURL = host + c.KeyAgreementAlgoURL
	c.KeyDerivationAlgoURL = host + c.KeyDerivationAlgoURL
	c.BlockCipherAlgoURL = host + c.BlockCipherAlgoURL
	c.CreateUploadURLs = lo.Map(c.CreateUploadURLs, func(x string, index int) string {
		return host + x
	})

	return c
}

type CreateUploadResponseV1 struct {
	Token           string `json:"token"`
	UploadUUID      string `json:"upload_uuid"`
	ConfirmationURL string `json:"confirmation_url"`
}

type ConfirmUploadRequestV1 struct {
	Token     string `json:"token"`
	PubKeyHex string `json:"pub_key_hex"`
	SigHex    string `json:"sig_hex"`
}

type ConfirmUploadResponseV1 struct {
	URL string `json:"url"`
}
