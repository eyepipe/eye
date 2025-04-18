package signature_reader_service

import (
	"fmt"
	"os"

	"context"
	"encoding/hex"
)

type SignatureReaderService struct {
}

func NewService() *SignatureReaderService {
	return &SignatureReaderService{}
}

// ReadHex reads signature in HEX format from CLI param or file
func (s *SignatureReaderService) ReadHex(ctx context.Context, opts ReadOpts) (sig []byte, err error) {
	switch {
	case len(opts.HEX) > 0:
		return s.readHex(ctx, opts.HEX)
	case len(opts.Filename) > 0:
		return s.readFile(ctx, opts.Filename)
	default:
		return nil, fmt.Errorf("%w: no hex signature file nor param specified", ErrEmptySignatureInput)
	}
}

func (s *SignatureReaderService) readHex(ctx context.Context, hexSig string) (sig []byte, err error) {
	sig, err = hex.DecodeString(hexSig)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrHexDecode, err)
	}

	return sig, nil
}

func (s *SignatureReaderService) readFile(ctx context.Context, filename string) (sig []byte, err error) {
	sig, err = os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFileFailed, err)
	}

	return s.readHex(ctx, string(sig))
}
