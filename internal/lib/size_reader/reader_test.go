package size_reader

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSizeReader_GetByteSize(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		data := make([]byte, 4096)
		buff := bytes.NewBuffer(data)
		reader := NewSizeReader(buff)
		_, err := io.ReadAll(reader)
		require.NoError(t, err)
		require.EqualValues(t, len(data), reader.GetByteSize())
	})
}

func TestSizeReader_Wrap(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		sized := New()
		data := []byte{1, 2, 3}
		buff := bytes.NewBuffer(data)
		reader := sized.Wrap(buff)
		got, err := io.ReadAll(reader)
		require.NoError(t, err)
		require.EqualValues(t, data, got)
		require.EqualValues(t, len(data), reader.GetByteSize())
	})
}
