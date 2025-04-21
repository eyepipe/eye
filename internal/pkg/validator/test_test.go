package validator

import (
	"context"
	"testing"
)

func MustNew(t *testing.T) (*Validator, context.Context) {
	return New(), context.Background()
}
