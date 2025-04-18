package crypto2

import (
	"hash"
	"testing"

	"crypto/sha256"
	"github.com/stretchr/testify/require"
)

func TestHashier_New(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		expected := sha256.New()
		hashier := NewHashier(func() hash.Hash {
			return expected
		})

		got := hashier.New()
		require.Equal(t, expected, got)
	})
}

func TestHashier_ByteSize(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		hashier := NewHashier(func() hash.Hash {
			return sha256.New()
		})

		got := hashier.ByteSize()
		require.Equal(t, 256>>3, got)
	})
}
