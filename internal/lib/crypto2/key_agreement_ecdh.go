package crypto2

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/samber/lo"
)

type KeyAgreementECDH struct {
	curve elliptic.Curve
	rand  io.Reader
}

func NewKeyAgreementECDH(curve elliptic.Curve) *KeyAgreementECDH {
	return &KeyAgreementECDH{
		curve: curve,
		rand:  rand.Reader,
	}
}

func (k *KeyAgreementECDH) Generate() (*KeyPair, error) {
	private, err := ecdsa.GenerateKey(k.curve, k.rand)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		Private: private,
		Public:  lo.ToPtr(private.PublicKey),
	}, nil
}

func (k *KeyAgreementECDH) ComputeSharedPrivatePublic(privateKey, publicKey any) ([]byte, error) {
	private, privateOk := privateKey.(*ecdsa.PrivateKey)
	if !privateOk {
		return nil, fmt.Errorf("wrong private key format: %T", privateKey)
	}

	public, publicOk := publicKey.(*ecdsa.PublicKey)
	if !publicOk {
		return nil, fmt.Errorf("wrong public key format: %T", privateKey)
	}

	x, _ := public.Curve.ScalarMult(public.X, public.Y, private.D.Bytes())
	if x == nil {
		return nil, fmt.Errorf("failed to compute shared x, y")
	}

	return x.Bytes(), nil
}
