package encryption_manager

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/eyepipe/eye/internal/pkg/eye_api"
	"github.com/eyepipe/eye/internal/pkg/services/export_service"
	"github.com/eyepipe/eye/pkg/proto"
	"golang.org/x/sync/errgroup"
)

func (m *Manager) Encrypt(ctx context.Context, reader io.Reader, writer io.Writer) (sig []byte, err error) {
	err = m.service.DeriveKey()
	if err != nil {
		return nil, fmt.Errorf("failed to derive key: %w", err)
	}

	writer = m.service.WrapSigner(writer)
	writer, err = m.service.WarpEncryptor(writer)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap encryptor: %w", err)
	}

	_, err = io.Copy(writer, reader)
	if err != nil {
		return nil, fmt.Errorf("failed to io.Copy: %w", err)
	}

	sig, err = m.service.Signature()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve signature: %w", err)
	}

	return sig, nil
}

func (m *Manager) SendEncrypt(ctx context.Context, reader io.Reader) (out *proto.ConfirmUploadResponseV1, err error) {
	pr, pw := io.Pipe()
	group, gCtx := errgroup.WithContext(ctx)
	var (
		res *proto.CreateUploadResponseV1
		sig []byte
	)

	contract, err := m.contract(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get contract: %w", err)
	}

	group.Go(func() error {
		defer func() {
			_ = pr.Close()
		}()

		res, err = m.upload(gCtx, contract, pr)
		switch {
		case err != nil:
			return fmt.Errorf("failed to upload: %w", err)
		default:
			return nil
		}
	})

	group.Go(func() error {
		defer func() {
			// notify reader with io.EOF
			_ = pw.Close()
		}()

		// reading original reader and writing
		// to the piped writer (data will be available inside the piped reader)
		// group context is to stop encrypting
		// when HTTP upload has been failed
		sig, err = m.Encrypt(gCtx, reader, pw)
		if err != nil {
			return fmt.Errorf("failed to encrypt: %w", err)
		}

		return nil
	})

	err = group.Wait()
	if err != nil {
		return nil, fmt.Errorf("errgroup has been failed: %w", err)
	}

	// confirm upload
	// (server validates signature)
	out, err = m.confirm(ctx, contract, res.Token, sig)
	if err != nil {
		return nil, fmt.Errorf("failed to confirm upload: %w", err)
	}

	return out, nil
}

func (m *Manager) contract(ctx context.Context) (*proto.ContractV1, error) {
	return m.api.GetContract(ctx, "https://api.eyepipe.pw/v1")
}

func (m *Manager) upload(ctx context.Context, contract *proto.ContractV1, reader io.Reader) (*proto.CreateUploadResponseV1, error) {
	res, err := m.api.Upload(ctx, contract.CreateUploadURL, reader, eye_api.UploadOpts{
		SignerAlgo: m.service.GetI().Scheme.SignAlgo.String(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload: %w", err)
	}

	return res, nil
}

func (m *Manager) confirm(ctx context.Context, contract *proto.ContractV1, token string, sig []byte) (out *proto.ConfirmUploadResponseV1, err error) {
	publicKey, err := export_service.PublicKeyToSPKI(m.service.GetI().SignerPair.Public)
	if err != nil {
		return nil, fmt.Errorf("failed to export public key: %w", err)
	}

	res, err := m.api.Confirm(ctx, contract.ConfirmUploadURL, &proto.ConfirmUploadRequestV1{
		Token:     token,
		SigHex:    hex.EncodeToString(sig),
		PubKeyHex: hex.EncodeToString(publicKey),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to api.Confirm: %w", err)
	}

	return res, nil
}
