package crypto2

import (
	"hash"
	"io"
)

//go:generate go-enum --file types.go --values

// KeyAgreementAlgo
// ENUM(
//
//	ECDH-P256
//	ECDH-P384
//	ECDH-P521
//
// )
type KeyAgreementAlgo string

// KeyAgreementPrimitive
// ENUM(
//
//	ECDH
//
// )
type KeyAgreementPrimitive string

// SignerAlgo
// ENUM(
//
//	ECDSA-P256-SHA256
//	ECDSA-P256-SHA384
//	ECDSA-P256-SHA512
//	ECDSA-P384-SHA256
//	ECDSA-P384-SHA384
//	ECDSA-P384-SHA512
//	ECDSA-P521-SHA256
//	ECDSA-P521-SHA384
//	ECDSA-P521-SHA512
//
// )
type SignerAlgo string

// SigningPrimitive
// ENUM(
//
//	ECDSA
//
// )
type SigningPrimitive string

// HashKeyDerivationPrimitive
// ENUM(
//
//	HKDF
//
// )
type HashKeyDerivationPrimitive string

// HashAlgo
// ENUM(
//
//	SHA256
//	SHA384
//	SHA512
//
// )
type HashAlgo string

// CurveAlgo
// ENUM(
//
//	P256
//	P384
//	P521
//
// )
type CurveAlgo string

// HashKeyDerivationAlgo
// ENUM(
//
//	HKDF-SHA256
//	HKDF-SHA384
//	HKDF-SHA512
//
// )
type HashKeyDerivationAlgo string

// BlockCipherAlgo
// ENUM(
//
//	AES_CTR_128
//	AES_CTR_192
//	AES_CTR_256
//
// )
type BlockCipherAlgo string

type KeyPair struct {
	Public  any
	Private any
}

type KeyDerivationResult struct {
	Key  []byte
	Salt []byte
}

// IHashier interface to work with hash function
type IHashier interface {
	New() hash.Hash
	ByteSize() int
}

// ISigner interface to work with digital signatures
type ISigner interface {
	Generate() (*KeyPair, error)
	GetHashier() IHashier
	Sign(data []byte, private any) ([]byte, error)
	Verify(data, signature []byte, public any) error
}

type IKeyAgreement interface {
	Generate() (*KeyPair, error)
	ComputeSharedPrivatePublic(privateKey, publicKey any) ([]byte, error)
}

type IBlockCipher interface {
	GetKeySizeBytes() int
	NewEncryptor(key []byte, writer io.Writer) (w io.Writer, iv []byte, err error)
	NewDecrypter(key, iv []byte, reader io.Reader) (r io.Reader, err error)
}

type IHashKeyDerivation interface {
	Derive(secret, salt []byte, keyByteSize int) (*KeyDerivationResult, error)
}
