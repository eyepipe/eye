package signature_reader_service

import (
	"encoding/hex"
	"os"
	"testing"

	"github.com/shlima/oi/tmp"
	"github.com/stretchr/testify/require"
)

func TestSignatureReaderService_ReadHex(t *testing.T) {
	t.Parallel()

	t.Run("it works with param", func(t *testing.T) {
		t.Parallel()
		service, ctx := MustNew(t)

		input := []byte{1, 2, 3}
		got, err := service.ReadHex(ctx, ReadOpts{HEX: hex.EncodeToString(input)})
		require.NoError(t, err)
		require.Equal(t, input, got)

		// broken input
		_, err = service.ReadHex(ctx, ReadOpts{HEX: "1"})
		require.ErrorIs(t, err, ErrHexDecode)
	})

	t.Run("it works with filename", func(t *testing.T) {
		t.Parallel()
		service, ctx := MustNew(t)

		temp := tmp.New()
		t.Cleanup(func() { _ = temp.Close() })

		ok, err := temp.File("ok.sig")
		require.NoError(t, err)

		broken, err := temp.File("broken.sig")
		require.NoError(t, err)

		// ok
		input := []byte{1, 2, 3}
		err = os.WriteFile(ok.Name(), []byte(hex.EncodeToString(input)), 0666)
		require.NoError(t, err)

		got, err := service.ReadHex(ctx, ReadOpts{Filename: ok.Name()})
		require.NoError(t, err)
		require.Equal(t, input, got)

		// broken
		err = os.WriteFile(broken.Name(), []byte{1}, 0666)
		require.NoError(t, err)

		_, err = service.ReadHex(ctx, ReadOpts{Filename: broken.Name()})
		require.ErrorIs(t, err, ErrHexDecode)

		// file not found
		_, err = service.ReadHex(ctx, ReadOpts{Filename: "foo"})
		require.ErrorIs(t, err, ErrFileFailed)
	})

	t.Run("it errors with empty input", func(t *testing.T) {
		t.Parallel()

		service, ctx := MustNew(t)
		_, err := service.ReadHex(ctx, ReadOpts{})
		require.ErrorIs(t, err, ErrEmptySignatureInput)
	})
}
