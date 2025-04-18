package keygen_service

import (
	"fmt"
	"github.com/eyepipe/eye/internal/pkg/container"
)

type Service struct {
	c *container.Container
}

func NewService(scheme container.Scheme) *Service {
	return &Service{
		c: &container.Container{
			Scheme: scheme,
		},
	}
}

func (s *Service) GetContainer() *container.Container {
	return s.c
}

func (s *Service) GenerateSignerPair() error {
	pair, err := s.c.Scheme.SignAlgo.ToSigner().Generate()
	if err != nil {
		return fmt.Errorf("failed to generate signer pair: %w", err)
	}

	s.c.SignerPair = pair
	return nil
}

func (s *Service) GenerateKeyAgreementPair() error {
	pair, err := s.c.Scheme.KeyAgreementAlgo.ToKeyAgreement().Generate()
	if err != nil {
		return fmt.Errorf("failed to generate key agreement pair: %w", err)
	}

	s.c.KeyAgreementPair = pair
	return nil
}
