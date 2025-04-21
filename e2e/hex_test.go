package e2e

import (
	"bytes"
	"encoding/hex"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCliHex(t *testing.T) {
	t.Parallel()

	t.Run("it encodes", func(t *testing.T) {
		t.Parallel()
		cli, _ := MustSetupCli(t)

		input := []byte("hello, word")
		stdout := new(strings.Builder)
		err := cli.
			WithStdin(bytes.NewReader(input)).
			WithStdout(stdout).
			Exec("hex")

		require.NoError(t, err)
		require.Equal(t, hex.EncodeToString(input), stdout.String())
	})

	t.Run("it decodes", func(t *testing.T) {
		t.Parallel()
		cli, _ := MustSetupCli(t)

		text := "foo bar"
		input := hex.EncodeToString([]byte(text))
		stdout := new(strings.Builder)

		err := cli.
			WithStdin(strings.NewReader(input)).
			WithStdout(stdout).
			WithStderr(os.Stderr).
			Exec("hex", "--dec")

		require.NoError(t, err)
		require.Equal(t, text, stdout.String())
	})
}
