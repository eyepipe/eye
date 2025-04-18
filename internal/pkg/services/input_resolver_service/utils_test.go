package input_resolver_service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsHttpURL(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		require.True(t, IsHttpURL("http://example.com"))
		require.True(t, IsHttpURL("HTTP://EXAMPLE.COM"))
		require.True(t, IsHttpURL("HttP://EXAMPLE.COM/v1?hello"))

		require.False(t, IsHttpURL("http"))
		require.False(t, IsHttpURL("https"))
		require.False(t, IsHttpURL("foo"))
	})
}
