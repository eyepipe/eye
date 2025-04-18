package encryption_service

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPutByte(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		value := byte(10)
		buff := bytes.NewBuffer(nil)
		err := PutByte(buff, value)
		require.NoError(t, err)

		var got byte
		err = GetByte(buff, &got)
		require.NoError(t, err)
		require.Equal(t, value, got)
	})
}
