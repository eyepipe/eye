package web

import (
	"github.com/eyepipe/eye/internal/lib/uuidv7"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGenS3KeyByUUID(t *testing.T) {
	t.Parallel()

	t.Run("it has date prefix", func(t *testing.T) {
		timestamp := time.Date(2009, 1, 2, 3, 0, 0, 0, time.UTC)
		uuid := uuidv7.NewWithTimeShard(timestamp, 0)
		got := GenS3KeyByUUID(uuid)
		require.True(t, strings.HasPrefix(got, "2009/01/02"), got)
	})

	t.Run("it has random prefix", func(t *testing.T) {
		uuid, err := uuidv7.Decode("01964027-88e9-7128-992a-35ea14a0655c")
		require.NoError(t, err)
		got := GenS3KeyByUUID(uuid)
		require.Contains(t, got, "35/ea1/01964027-88e9-7128-992a-35ea14a0655c")
	})
}
