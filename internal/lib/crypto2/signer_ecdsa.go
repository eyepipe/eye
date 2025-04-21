package crypto2

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"io"

	"github.com/samber/lo"
)

var ErrSignatureInvalid = errors.New("ERR_SIGNATURE_INVALID")

type SignerECDSA struct {
	curve   elliptic.Curve
	hashier IHashier
	rand    io.Reader
}

func NewSignerECDSA(curve elliptic.Curve, hashier IHashier) *SignerECDSA {
	return &SignerECDSA{
		curve:   curve,
		hashier: hashier,
		rand:    rand.Reader,
	}
}

func (s *SignerECDSA) Generate() (*KeyPair, error) {
	private, err := ecdsa.GenerateKey(s.curve, s.rand)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		Private: private,
		Public:  lo.ToPtr(private.PublicKey),
	}, nil
}

func (s *SignerECDSA) GetHashier() IHashier {
	return s.hashier
}

func (s *SignerECDSA) Sign(data []byte, private any) ([]byte, error) {
	return ecdsa.SignASN1(rand.Reader, private.(*ecdsa.PrivateKey), data)
}

func (s *SignerECDSA) Verify(data, signature []byte, public any) error {
	valid := ecdsa.VerifyASN1(public.(*ecdsa.PublicKey), data, signature)
	switch {
	case valid:
		return nil
	default:
		return ErrSignatureInvalid
	}
}
