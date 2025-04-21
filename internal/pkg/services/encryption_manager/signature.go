package encryption_manager

import (
	"context"
	"fmt"
)

func (m *Manager) VerifySignature(ctx context.Context, verification, sig []byte) error {
	signer := m.service.GetP().Scheme.SignAlgo.ToSigner()
	err := signer.Verify(verification, sig, m.service.GetP().SignerPair.Public)
	switch {
	case err != nil:
		return fmt.Errorf("failed to signer.Verify: %w", err)
	default:
		return nil
	}
}
