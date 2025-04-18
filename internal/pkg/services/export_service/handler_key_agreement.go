package export_service

import (
	"fmt"
	"github.com/eyepipe/eye/internal/pkg/container"
)

func ExportKeyAgreementPrivate(c *container.Container) ([]byte, error) {
	return PrivateKeyPKCS8(c.KeyAgreementPair.Private)
}

func ExportKeyAgreementPublic(c *container.Container) ([]byte, error) {
	return PublicKeyToSPKI(c.KeyAgreementPair.Public)
}

func ImportKeyAgreementPrivate(c *container.Container, bytea []byte) error {
	key, err := PKCS8ToPrivateKey(bytea)
	if err != nil {
		return fmt.Errorf("failed to parse PKCS8: %w", err)
	}

	c.KeyAgreementPair.Private = key
	return nil
}

func ImportKeyAgreementPublic(c *container.Container, bytea []byte) error {
	key, err := SPKIToPublicKey(bytea)
	if err != nil {
		return fmt.Errorf("failed to parse SPKI: %w", err)
	}

	c.KeyAgreementPair.Public = key
	return nil
}
