package uuidv7

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {
	t.Parallel()

	// @refs https://www.uuid.lol/uuid/decode
	t.Run("it works", func(t *testing.T) {
		t.Parallel()
		got, err := Decode("061cb26a-54b8-7a52-8000-2124e7041024")
		require.NoError(t, err)
		require.Equal(t, time.UnixMilli(6720322163896), got.Time)
	})

	t.Run("it fails", func(t *testing.T) {
		t.Parallel()
		_, err := Decode("061cb26a-54b8-7a52-8000-2124e70410")
		require.ErrorIs(t, err, ErrDecodeFailed)

		_, err = Decode("doo")
		require.ErrorIs(t, err, ErrDecodeFailed)

		_, err = Decode("")
		require.ErrorIs(t, err, ErrDecodeFailed)
	})
}

func TestUUIDv7_String(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		input := "061cb26a-54b8-7a52-8000-2124e7041024"
		got, err := Decode(input)
		require.NoError(t, err)
		require.Equal(t, input, got.String())
	})
}

func TestNewWithShard(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		shard := uint16(3)
		uuid := NewWithShard(shard)
		require.Equal(t, shard, uuid.Shard)

		got, err := Decode(uuid.String())
		require.NoError(t, err)
		require.Equal(t, uuid, got)
	})
}

func TestNewWithTimeShard(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		timestamp := time.Now().Round(TimePrecision)
		shard := uint16(3)
		uuid := NewWithTimeShard(timestamp, shard)
		require.Equal(t, shard, uuid.Shard)
		require.Equal(t, timestamp, uuid.Time)

		got, err := Decode(uuid.String())
		require.NoError(t, err)
		require.Equal(t, uuid, got)
	})
}
