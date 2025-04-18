package signature_reader_service

import (
	"context"
	"testing"
)

func MustNew(t *testing.T) (*SignatureReaderService, context.Context) {
	return NewService(), context.Background()
}
