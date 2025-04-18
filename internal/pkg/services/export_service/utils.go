package export_service

import (
	"crypto/x509"
	"fmt"
)

func PublicKeyToSPKI(pubKey any) ([]byte, error) {
	derBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to x509.MarshalPKIXPublicKey: %v", err)
	}
	return derBytes, nil
}

func PrivateKeyPKCS8(privateKey any) ([]byte, error) {
	pkcs8Bytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to x509.MarshalPKCS8PrivateKey: %w", err)
	}

	return pkcs8Bytes, nil
}

func SPKIToPublicKey(bytes []byte) (any, error) {
	key, err := x509.ParsePKIXPublicKey(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to x509.ParsePKIXPublicKey: %w", err)
	}

	return key, err
}

func PKCS8ToPrivateKey(bytes []byte) (any, error) {
	key, err := x509.ParsePKCS8PrivateKey(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to x509.ParsePKCS8PrivateKey: %w", err)
	}

	return key, nil
}
