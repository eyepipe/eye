// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package crypto2

import (
	"errors"
	"fmt"
)

const (
	// BlockCipherAlgoAESCTR128 is a BlockCipherAlgo of type AES_CTR_128.
	BlockCipherAlgoAESCTR128 BlockCipherAlgo = "AES_CTR_128"
	// BlockCipherAlgoAESCTR192 is a BlockCipherAlgo of type AES_CTR_192.
	BlockCipherAlgoAESCTR192 BlockCipherAlgo = "AES_CTR_192"
	// BlockCipherAlgoAESCTR256 is a BlockCipherAlgo of type AES_CTR_256.
	BlockCipherAlgoAESCTR256 BlockCipherAlgo = "AES_CTR_256"
)

var ErrInvalidBlockCipherAlgo = errors.New("not a valid BlockCipherAlgo")

// BlockCipherAlgoValues returns a list of the values for BlockCipherAlgo
func BlockCipherAlgoValues() []BlockCipherAlgo {
	return []BlockCipherAlgo{
		BlockCipherAlgoAESCTR128,
		BlockCipherAlgoAESCTR192,
		BlockCipherAlgoAESCTR256,
	}
}

// String implements the Stringer interface.
func (x BlockCipherAlgo) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x BlockCipherAlgo) IsValid() bool {
	_, err := ParseBlockCipherAlgo(string(x))
	return err == nil
}

var _BlockCipherAlgoValue = map[string]BlockCipherAlgo{
	"AES_CTR_128": BlockCipherAlgoAESCTR128,
	"AES_CTR_192": BlockCipherAlgoAESCTR192,
	"AES_CTR_256": BlockCipherAlgoAESCTR256,
}

// ParseBlockCipherAlgo attempts to convert a string to a BlockCipherAlgo.
func ParseBlockCipherAlgo(name string) (BlockCipherAlgo, error) {
	if x, ok := _BlockCipherAlgoValue[name]; ok {
		return x, nil
	}
	return BlockCipherAlgo(""), fmt.Errorf("%s is %w", name, ErrInvalidBlockCipherAlgo)
}

const (
	// CurveAlgoP256 is a CurveAlgo of type P256.
	CurveAlgoP256 CurveAlgo = "P256"
	// CurveAlgoP384 is a CurveAlgo of type P384.
	CurveAlgoP384 CurveAlgo = "P384"
	// CurveAlgoP521 is a CurveAlgo of type P521.
	CurveAlgoP521 CurveAlgo = "P521"
)

var ErrInvalidCurveAlgo = errors.New("not a valid CurveAlgo")

// CurveAlgoValues returns a list of the values for CurveAlgo
func CurveAlgoValues() []CurveAlgo {
	return []CurveAlgo{
		CurveAlgoP256,
		CurveAlgoP384,
		CurveAlgoP521,
	}
}

// String implements the Stringer interface.
func (x CurveAlgo) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x CurveAlgo) IsValid() bool {
	_, err := ParseCurveAlgo(string(x))
	return err == nil
}

var _CurveAlgoValue = map[string]CurveAlgo{
	"P256": CurveAlgoP256,
	"P384": CurveAlgoP384,
	"P521": CurveAlgoP521,
}

// ParseCurveAlgo attempts to convert a string to a CurveAlgo.
func ParseCurveAlgo(name string) (CurveAlgo, error) {
	if x, ok := _CurveAlgoValue[name]; ok {
		return x, nil
	}
	return CurveAlgo(""), fmt.Errorf("%s is %w", name, ErrInvalidCurveAlgo)
}

const (
	// HashAlgoSHA256 is a HashAlgo of type SHA256.
	HashAlgoSHA256 HashAlgo = "SHA256"
	// HashAlgoSHA384 is a HashAlgo of type SHA384.
	HashAlgoSHA384 HashAlgo = "SHA384"
	// HashAlgoSHA512 is a HashAlgo of type SHA512.
	HashAlgoSHA512 HashAlgo = "SHA512"
)

var ErrInvalidHashAlgo = errors.New("not a valid HashAlgo")

// HashAlgoValues returns a list of the values for HashAlgo
func HashAlgoValues() []HashAlgo {
	return []HashAlgo{
		HashAlgoSHA256,
		HashAlgoSHA384,
		HashAlgoSHA512,
	}
}

// String implements the Stringer interface.
func (x HashAlgo) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x HashAlgo) IsValid() bool {
	_, err := ParseHashAlgo(string(x))
	return err == nil
}

var _HashAlgoValue = map[string]HashAlgo{
	"SHA256": HashAlgoSHA256,
	"SHA384": HashAlgoSHA384,
	"SHA512": HashAlgoSHA512,
}

// ParseHashAlgo attempts to convert a string to a HashAlgo.
func ParseHashAlgo(name string) (HashAlgo, error) {
	if x, ok := _HashAlgoValue[name]; ok {
		return x, nil
	}
	return HashAlgo(""), fmt.Errorf("%s is %w", name, ErrInvalidHashAlgo)
}

const (
	// HashKeyDerivationAlgoHKDFSHA256 is a HashKeyDerivationAlgo of type HKDF-SHA256.
	HashKeyDerivationAlgoHKDFSHA256 HashKeyDerivationAlgo = "HKDF-SHA256"
	// HashKeyDerivationAlgoHKDFSHA384 is a HashKeyDerivationAlgo of type HKDF-SHA384.
	HashKeyDerivationAlgoHKDFSHA384 HashKeyDerivationAlgo = "HKDF-SHA384"
	// HashKeyDerivationAlgoHKDFSHA512 is a HashKeyDerivationAlgo of type HKDF-SHA512.
	HashKeyDerivationAlgoHKDFSHA512 HashKeyDerivationAlgo = "HKDF-SHA512"
)

var ErrInvalidHashKeyDerivationAlgo = errors.New("not a valid HashKeyDerivationAlgo")

// HashKeyDerivationAlgoValues returns a list of the values for HashKeyDerivationAlgo
func HashKeyDerivationAlgoValues() []HashKeyDerivationAlgo {
	return []HashKeyDerivationAlgo{
		HashKeyDerivationAlgoHKDFSHA256,
		HashKeyDerivationAlgoHKDFSHA384,
		HashKeyDerivationAlgoHKDFSHA512,
	}
}

// String implements the Stringer interface.
func (x HashKeyDerivationAlgo) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x HashKeyDerivationAlgo) IsValid() bool {
	_, err := ParseHashKeyDerivationAlgo(string(x))
	return err == nil
}

var _HashKeyDerivationAlgoValue = map[string]HashKeyDerivationAlgo{
	"HKDF-SHA256": HashKeyDerivationAlgoHKDFSHA256,
	"HKDF-SHA384": HashKeyDerivationAlgoHKDFSHA384,
	"HKDF-SHA512": HashKeyDerivationAlgoHKDFSHA512,
}

// ParseHashKeyDerivationAlgo attempts to convert a string to a HashKeyDerivationAlgo.
func ParseHashKeyDerivationAlgo(name string) (HashKeyDerivationAlgo, error) {
	if x, ok := _HashKeyDerivationAlgoValue[name]; ok {
		return x, nil
	}
	return HashKeyDerivationAlgo(""), fmt.Errorf("%s is %w", name, ErrInvalidHashKeyDerivationAlgo)
}

const (
	// HashKeyDerivationPrimitiveHKDF is a HashKeyDerivationPrimitive of type HKDF.
	HashKeyDerivationPrimitiveHKDF HashKeyDerivationPrimitive = "HKDF"
)

var ErrInvalidHashKeyDerivationPrimitive = errors.New("not a valid HashKeyDerivationPrimitive")

// HashKeyDerivationPrimitiveValues returns a list of the values for HashKeyDerivationPrimitive
func HashKeyDerivationPrimitiveValues() []HashKeyDerivationPrimitive {
	return []HashKeyDerivationPrimitive{
		HashKeyDerivationPrimitiveHKDF,
	}
}

// String implements the Stringer interface.
func (x HashKeyDerivationPrimitive) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x HashKeyDerivationPrimitive) IsValid() bool {
	_, err := ParseHashKeyDerivationPrimitive(string(x))
	return err == nil
}

var _HashKeyDerivationPrimitiveValue = map[string]HashKeyDerivationPrimitive{
	"HKDF": HashKeyDerivationPrimitiveHKDF,
}

// ParseHashKeyDerivationPrimitive attempts to convert a string to a HashKeyDerivationPrimitive.
func ParseHashKeyDerivationPrimitive(name string) (HashKeyDerivationPrimitive, error) {
	if x, ok := _HashKeyDerivationPrimitiveValue[name]; ok {
		return x, nil
	}
	return HashKeyDerivationPrimitive(""), fmt.Errorf("%s is %w", name, ErrInvalidHashKeyDerivationPrimitive)
}

const (
	// KeyAgreementAlgoECDHP256 is a KeyAgreementAlgo of type ECDH-P256.
	KeyAgreementAlgoECDHP256 KeyAgreementAlgo = "ECDH-P256"
	// KeyAgreementAlgoECDHP384 is a KeyAgreementAlgo of type ECDH-P384.
	KeyAgreementAlgoECDHP384 KeyAgreementAlgo = "ECDH-P384"
	// KeyAgreementAlgoECDHP521 is a KeyAgreementAlgo of type ECDH-P521.
	KeyAgreementAlgoECDHP521 KeyAgreementAlgo = "ECDH-P521"
)

var ErrInvalidKeyAgreementAlgo = errors.New("not a valid KeyAgreementAlgo")

// KeyAgreementAlgoValues returns a list of the values for KeyAgreementAlgo
func KeyAgreementAlgoValues() []KeyAgreementAlgo {
	return []KeyAgreementAlgo{
		KeyAgreementAlgoECDHP256,
		KeyAgreementAlgoECDHP384,
		KeyAgreementAlgoECDHP521,
	}
}

// String implements the Stringer interface.
func (x KeyAgreementAlgo) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x KeyAgreementAlgo) IsValid() bool {
	_, err := ParseKeyAgreementAlgo(string(x))
	return err == nil
}

var _KeyAgreementAlgoValue = map[string]KeyAgreementAlgo{
	"ECDH-P256": KeyAgreementAlgoECDHP256,
	"ECDH-P384": KeyAgreementAlgoECDHP384,
	"ECDH-P521": KeyAgreementAlgoECDHP521,
}

// ParseKeyAgreementAlgo attempts to convert a string to a KeyAgreementAlgo.
func ParseKeyAgreementAlgo(name string) (KeyAgreementAlgo, error) {
	if x, ok := _KeyAgreementAlgoValue[name]; ok {
		return x, nil
	}
	return KeyAgreementAlgo(""), fmt.Errorf("%s is %w", name, ErrInvalidKeyAgreementAlgo)
}

const (
	// KeyAgreementPrimitiveECDH is a KeyAgreementPrimitive of type ECDH.
	KeyAgreementPrimitiveECDH KeyAgreementPrimitive = "ECDH"
)

var ErrInvalidKeyAgreementPrimitive = errors.New("not a valid KeyAgreementPrimitive")

// KeyAgreementPrimitiveValues returns a list of the values for KeyAgreementPrimitive
func KeyAgreementPrimitiveValues() []KeyAgreementPrimitive {
	return []KeyAgreementPrimitive{
		KeyAgreementPrimitiveECDH,
	}
}

// String implements the Stringer interface.
func (x KeyAgreementPrimitive) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x KeyAgreementPrimitive) IsValid() bool {
	_, err := ParseKeyAgreementPrimitive(string(x))
	return err == nil
}

var _KeyAgreementPrimitiveValue = map[string]KeyAgreementPrimitive{
	"ECDH": KeyAgreementPrimitiveECDH,
}

// ParseKeyAgreementPrimitive attempts to convert a string to a KeyAgreementPrimitive.
func ParseKeyAgreementPrimitive(name string) (KeyAgreementPrimitive, error) {
	if x, ok := _KeyAgreementPrimitiveValue[name]; ok {
		return x, nil
	}
	return KeyAgreementPrimitive(""), fmt.Errorf("%s is %w", name, ErrInvalidKeyAgreementPrimitive)
}

const (
	// SignerAlgoECDSAP256SHA256 is a SignerAlgo of type ECDSA-P256-SHA256.
	SignerAlgoECDSAP256SHA256 SignerAlgo = "ECDSA-P256-SHA256"
	// SignerAlgoECDSAP256SHA384 is a SignerAlgo of type ECDSA-P256-SHA384.
	SignerAlgoECDSAP256SHA384 SignerAlgo = "ECDSA-P256-SHA384"
	// SignerAlgoECDSAP256SHA512 is a SignerAlgo of type ECDSA-P256-SHA512.
	SignerAlgoECDSAP256SHA512 SignerAlgo = "ECDSA-P256-SHA512"
	// SignerAlgoECDSAP384SHA256 is a SignerAlgo of type ECDSA-P384-SHA256.
	SignerAlgoECDSAP384SHA256 SignerAlgo = "ECDSA-P384-SHA256"
	// SignerAlgoECDSAP384SHA384 is a SignerAlgo of type ECDSA-P384-SHA384.
	SignerAlgoECDSAP384SHA384 SignerAlgo = "ECDSA-P384-SHA384"
	// SignerAlgoECDSAP384SHA512 is a SignerAlgo of type ECDSA-P384-SHA512.
	SignerAlgoECDSAP384SHA512 SignerAlgo = "ECDSA-P384-SHA512"
	// SignerAlgoECDSAP521SHA256 is a SignerAlgo of type ECDSA-P521-SHA256.
	SignerAlgoECDSAP521SHA256 SignerAlgo = "ECDSA-P521-SHA256"
	// SignerAlgoECDSAP521SHA384 is a SignerAlgo of type ECDSA-P521-SHA384.
	SignerAlgoECDSAP521SHA384 SignerAlgo = "ECDSA-P521-SHA384"
	// SignerAlgoECDSAP521SHA512 is a SignerAlgo of type ECDSA-P521-SHA512.
	SignerAlgoECDSAP521SHA512 SignerAlgo = "ECDSA-P521-SHA512"
)

var ErrInvalidSignerAlgo = errors.New("not a valid SignerAlgo")

// SignerAlgoValues returns a list of the values for SignerAlgo
func SignerAlgoValues() []SignerAlgo {
	return []SignerAlgo{
		SignerAlgoECDSAP256SHA256,
		SignerAlgoECDSAP256SHA384,
		SignerAlgoECDSAP256SHA512,
		SignerAlgoECDSAP384SHA256,
		SignerAlgoECDSAP384SHA384,
		SignerAlgoECDSAP384SHA512,
		SignerAlgoECDSAP521SHA256,
		SignerAlgoECDSAP521SHA384,
		SignerAlgoECDSAP521SHA512,
	}
}

// String implements the Stringer interface.
func (x SignerAlgo) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x SignerAlgo) IsValid() bool {
	_, err := ParseSignerAlgo(string(x))
	return err == nil
}

var _SignerAlgoValue = map[string]SignerAlgo{
	"ECDSA-P256-SHA256": SignerAlgoECDSAP256SHA256,
	"ECDSA-P256-SHA384": SignerAlgoECDSAP256SHA384,
	"ECDSA-P256-SHA512": SignerAlgoECDSAP256SHA512,
	"ECDSA-P384-SHA256": SignerAlgoECDSAP384SHA256,
	"ECDSA-P384-SHA384": SignerAlgoECDSAP384SHA384,
	"ECDSA-P384-SHA512": SignerAlgoECDSAP384SHA512,
	"ECDSA-P521-SHA256": SignerAlgoECDSAP521SHA256,
	"ECDSA-P521-SHA384": SignerAlgoECDSAP521SHA384,
	"ECDSA-P521-SHA512": SignerAlgoECDSAP521SHA512,
}

// ParseSignerAlgo attempts to convert a string to a SignerAlgo.
func ParseSignerAlgo(name string) (SignerAlgo, error) {
	if x, ok := _SignerAlgoValue[name]; ok {
		return x, nil
	}
	return SignerAlgo(""), fmt.Errorf("%s is %w", name, ErrInvalidSignerAlgo)
}

const (
	// SigningPrimitiveECDSA is a SigningPrimitive of type ECDSA.
	SigningPrimitiveECDSA SigningPrimitive = "ECDSA"
)

var ErrInvalidSigningPrimitive = errors.New("not a valid SigningPrimitive")

// SigningPrimitiveValues returns a list of the values for SigningPrimitive
func SigningPrimitiveValues() []SigningPrimitive {
	return []SigningPrimitive{
		SigningPrimitiveECDSA,
	}
}

// String implements the Stringer interface.
func (x SigningPrimitive) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x SigningPrimitive) IsValid() bool {
	_, err := ParseSigningPrimitive(string(x))
	return err == nil
}

var _SigningPrimitiveValue = map[string]SigningPrimitive{
	"ECDSA": SigningPrimitiveECDSA,
}

// ParseSigningPrimitive attempts to convert a string to a SigningPrimitive.
func ParseSigningPrimitive(name string) (SigningPrimitive, error) {
	if x, ok := _SigningPrimitiveValue[name]; ok {
		return x, nil
	}
	return SigningPrimitive(""), fmt.Errorf("%s is %w", name, ErrInvalidSigningPrimitive)
}
