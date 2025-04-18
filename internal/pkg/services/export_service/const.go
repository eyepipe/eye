package export_service

import "errors"

const (
	PEMSchemeProto            = "SCHEME PROTO"
	PEMSignerPrivateKey       = "SIGNER PRIVATE KEY"
	PEMSignerPublicKey        = "SIGNER PUBLIC KEY"
	PEMKeyAgreementPrivateKey = "KEY AGREEMENT PRIVATE KEY"
	PEMKeyAgreementPublicKey  = "KEY AGREEMENT PUBLIC KEY"
)

var (
	ErrPemEncode       = errors.New("ERR_PEM_ENCODE")
	ErrPemExport       = errors.New("ERR_PEM_EXPORT")
	ErrPemBlockUnknown = errors.New("ERR_PEM_BLOCK_UNKNOWN")
	ErrPemParse        = errors.New("ERR_PEM_PARSE")
)
