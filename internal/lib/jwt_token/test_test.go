package jwt_token

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func MustNew(t *testing.T) *Token {
	got, err := New("test", AlgoHS512, []byte{1, 2, 3})
	require.NoError(t, err)
	return got
}
