package encryption_manager

import (
	"context"
	"fmt"
	"io"

	"github.com/eyepipe/eye/internal/pkg/services/signature_reader_service"
	"github.com/eyepipe/eye/pkg/proto"
)

func (m *Manager) Decrypt(ctx context.Context, sig []byte, reader io.Reader, writer io.Writer) error {
	reader = io.TeeReader(reader, m.service.WrapVerifier(io.Discard))

	err := m.service.ReadEnvelope(reader)
	if err != nil {
		return fmt.Errorf("failed to read envelope: %w", err)
	}

	err = m.service.DeriveKey()
	if err != nil {
		return fmt.Errorf("failed to derive key: %w", err)
	}

	reader, err = m.service.WarpDecrypter(reader)
	if err != nil {
		return fmt.Errorf("failed to wrap decrypter: %w", err)
	}

	_, err = io.Copy(writer, reader)
	if err != nil {
		return fmt.Errorf("failed to read io.ReadAll: %w", err)
	}

	verification := m.service.Verification()
	return m.VerifySignature(ctx, verification, sig)
}

func (m *Manager) DownloadDecrypt(ctx context.Context, URL string, writer io.Writer) error {
	location, sig, err := m.downloadHeader(ctx, URL)
	if err != nil {
		return fmt.Errorf("failed to download header: %w", err)
	}

	res, err := m.downloader.R().
		SetDoNotParseResponse(true).
		Get(location)
	switch {
	case err != nil:
		return fmt.Errorf("failed to download: %w", err)
	case res.IsError():
		return fmt.Errorf("bad status code: <%d>", res.StatusCode())
	}

	defer res.Body.Close()
	return m.Decrypt(ctx, sig, res.Body, writer)
}

func (m *Manager) downloadHeader(ctx context.Context, URL string) (locationURL string, sig []byte, err error) {
	res, err := m.downloader.R().Get(URL)
	switch {
	case err != nil:
		return "", nil, fmt.Errorf("failed to downloader GET: %w", err)
	case res.IsError():
		return "", nil, fmt.Errorf("bad status code: <%d>", res.StatusCode())
	}

	sigHex := res.Header().Get(proto.HeaderSigHex)
	if len(sigHex) == 0 {
		return "", nil, fmt.Errorf("missing expected signature header: <%s>", proto.HeaderSigHex)
	}

	location := res.Header().Get("location")
	if len(location) == 0 {
		return "", nil, fmt.Errorf("missing location redirect header")
	}

	sig, err = m.signatureReader.ReadHex(ctx, signature_reader_service.ReadOpts{
		HEX: sigHex,
	})
	if err != nil {
		return "", nil, fmt.Errorf("failed to read signature: %w", err)
	}

	return location, sig, nil
}
