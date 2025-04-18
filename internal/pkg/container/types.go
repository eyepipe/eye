package container

import (
	crypto3 "github.com/eyepipe/eye/internal/lib/crypto2"
)

type Scheme struct {
	SignAlgo              crypto3.SignerAlgo            `json:"sign_algo"`
	KeyAgreementAlgo      crypto3.KeyAgreementAlgo      `json:"key_agreement_algo"`
	BlockCipherAlgo       crypto3.BlockCipherAlgo       `json:"block_cipher_algo"`
	HashKeyDerivationAlgo crypto3.HashKeyDerivationAlgo `json:"hash_key_derivation_algo"`
}

// SchemeV1Super
// Superb secure scheme
var SchemeV1Super = Scheme{
	SignAlgo:              crypto3.SignerAlgoECDSAP521SHA512,
	KeyAgreementAlgo:      crypto3.KeyAgreementAlgoECDHP521,
	BlockCipherAlgo:       crypto3.BlockCipherAlgoAESCTR256,
	HashKeyDerivationAlgo: crypto3.HashKeyDerivationAlgoHKDFSHA512,
}

// SchemeV1High
// Highly secure scheme for common use
var SchemeV1High = Scheme{
	SignAlgo:              crypto3.SignerAlgoECDSAP256SHA256,
	KeyAgreementAlgo:      crypto3.KeyAgreementAlgoECDHP256,
	BlockCipherAlgo:       crypto3.BlockCipherAlgoAESCTR192,
	HashKeyDerivationAlgo: crypto3.HashKeyDerivationAlgoHKDFSHA256,
}

// SchemeDefault default scheme is used in CLI keygen command
var SchemeDefault = SchemeV1Super

type Container struct {
	Scheme           Scheme // used crypto scheme
	SignerPair       *crypto3.KeyPair
	KeyAgreementPair *crypto3.KeyPair
}

// NewEmptyContainer returns blank Container object
func NewEmptyContainer() *Container {
	return &Container{
		SignerPair:       new(crypto3.KeyPair),
		KeyAgreementPair: new(crypto3.KeyPair),
	}
}
