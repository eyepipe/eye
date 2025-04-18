package crypto2

import (
	"crypto/hkdf"
	"crypto/rand"
	"fmt"
	"io"
)

type HashKeyDerivationHKDF struct {
	hashier IHashier
	rand    io.Reader
}

func NewHashKeyDerivationHKDF(hashier IHashier) *HashKeyDerivationHKDF {
	return &HashKeyDerivationHKDF{
		hashier: hashier,
		rand:    rand.Reader,
	}
}

func (h *HashKeyDerivationHKDF) Derive(secret, salt []byte, keyByteSize int) (*KeyDerivationResult, error) {
	requiredSaltSize := h.hashier.ByteSize()
	if len(salt) != requiredSaltSize {
		salt = make([]byte, h.hashier.ByteSize())
		_, err := rand.Read(salt)
		if err != nil {
			return nil, fmt.Errorf("failed to generate salt: %w", err)
		}
	}

	key, err := hkdf.Key(h.hashier.New, secret, salt, "", keyByteSize)
	if err != nil {
		return nil, fmt.Errorf("failed to derive key: %e", err)
	}

	return &KeyDerivationResult{
		Key:  key,
		Salt: salt,
	}, nil
}
