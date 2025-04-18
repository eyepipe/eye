package export_service

import (
	"fmt"
	"github.com/eyepipe/eye/internal/pkg/container"
)

func ExportSignerPrivate(c *container.Container) ([]byte, error) {
	return PrivateKeyPKCS8(c.SignerPair.Private)
}

func ExportSignerPublic(c *container.Container) ([]byte, error) {
	return PublicKeyToSPKI(c.SignerPair.Public)
}

func ImportSignerPrivate(c *container.Container, bytea []byte) error {
	key, err := PKCS8ToPrivateKey(bytea)
	if err != nil {
		return fmt.Errorf("failed to parse PKCS8: %w", err)
	}

	c.SignerPair.Private = key
	return nil
}

func ImportSignerPublic(c *container.Container, bytea []byte) error {
	key, err := SPKIToPublicKey(bytea)
	if err != nil {
		return fmt.Errorf("failed to parse SPKI: %w", err)
	}

	c.SignerPair.Public = key
	return nil
}
