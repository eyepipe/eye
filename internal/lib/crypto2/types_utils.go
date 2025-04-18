package crypto2

import (
	"crypto/elliptic"
	"crypto/sha256"
	"crypto/sha512"
)

func (c CurveAlgo) ToEllipticCurve() elliptic.Curve {
	switch c {
	case CurveAlgoP384:
		return elliptic.P384()
	case CurveAlgoP256:
		return elliptic.P256()
	case CurveAlgoP521:
		return elliptic.P521()
	default:
		return nil
	}
}

func (h HashAlgo) ToHashier() IHashier {
	switch h {
	case HashAlgoSHA256:
		return NewHashier(sha256.New)
	case HashAlgoSHA384:
		return NewHashier(sha512.New384)
	case HashAlgoSHA512:
		return NewHashier(sha512.New)
	default:
		return nil
	}
}

// ToSigner
// input example "ECDSA-P256-SHA256"
func (s SignerAlgo) ToSigner() ISigner {
	p := NewStringPart(s.String(), "-")

	switch SigningPrimitive(p.Get(0)) {
	case SigningPrimitiveECDSA:
		curve, hash := CurveAlgo(p.Get(1)), HashAlgo(p.Get(2))
		return NewSignerECDSA(curve.ToEllipticCurve(), hash.ToHashier())
	default:
		return nil
	}
}

func (h HashKeyDerivationAlgo) ToDerivation() IHashKeyDerivation {
	p := NewStringPart(h.String(), "-")

	switch HashKeyDerivationPrimitive(p.Get(0)) {
	case HashKeyDerivationPrimitiveHKDF:
		hash := HashAlgo(p.Get(1))
		return NewHashKeyDerivationHKDF(hash.ToHashier())
	default:
		return nil
	}
}

func (k KeyAgreementAlgo) ToKeyAgreement() IKeyAgreement {
	p := NewStringPart(k.String(), "-")

	switch KeyAgreementPrimitive(p.Get(0)) {
	case KeyAgreementPrimitiveECDH:
		curve := CurveAlgo(p.Get(1))
		return NewKeyAgreementECDH(curve.ToEllipticCurve())
	default:
		return nil
	}
}

func (b BlockCipherAlgo) ToCipher() IBlockCipher {
	switch b {
	case BlockCipherAlgoAESCTR128:
		return NewBlockCipherAesCtr(128 >> 3)
	case BlockCipherAlgoAESCTR192:
		return NewBlockCipherAesCtr(192 >> 3)
	case BlockCipherAlgoAESCTR256:
		return NewBlockCipherAesCtr(256 >> 3)
	default:
		return nil
	}
}
