package e2e

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCliEncrypt(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		t.Parallel()
		cli, _ := MustSetupCli(t)
		stdout := new(bytes.Buffer)
		stderr := new(strings.Builder)
		text := "Hello, world"

		// encrypt
		err := cli.
			WithStdout(stdout).
			WithStdin(strings.NewReader(text)).
			WithStderr(stderr).
			Exec("encrypt", "-i", "testdata/super_v1_00.priv", "-p", "testdata/super_v1_00.pub")
		require.NoError(t, err)
		require.NotEmpty(t, stdout, "encrypted data")
		require.NotEmpty(t, stderr, "message signature")

		stdin := bytes.NewReader(stdout.Bytes())
		sigHex := stderr.String()
		stdout = new(bytes.Buffer)
		stderr = new(strings.Builder)

		// decrypt
		err = cli.
			WithStdin(stdin).
			WithStdout(stdout).
			WithStderr(stderr).
			Exec("decrypt", "-i", "testdata/super_v1_00.priv", "-p", "testdata/super_v1_00.pub", "--sig-hex", sigHex)
		require.NoError(t, err)
		require.Empty(t, stderr)
		require.Equal(t, text, stdout.String())
	})
}
