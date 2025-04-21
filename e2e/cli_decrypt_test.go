package e2e

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCliDecrypt(t *testing.T) {
	t.Parallel()

	t.Run("it works (signature from file)", func(t *testing.T) {
		t.Parallel()
		cli, _ := MustSetupCli(t)
		encrypted := MustReadAll(t, "testdata/message_super_v1_00.enc")
		stdout := new(bytes.Buffer)
		stderr := new(strings.Builder)

		err := cli.
			WithStdin(bytes.NewReader(encrypted)).
			WithStdout(stdout).
			WithStderr(stderr).
			Exec("decrypt", "-i", "testdata/super_v1_00.priv", "-p", "testdata/super_v1_00.pub", "--sig", "testdata/message_super_v1_00.enc.sig")
		require.NoError(t, err)
		require.Empty(t, stderr)
		require.Equal(t, "message", stdout.String())
	})

	t.Run("it works (signature from hex)", func(t *testing.T) {
		t.Parallel()
		cli, _ := MustSetupCli(t)
		encrypted := MustReadAll(t, "testdata/message_super_v1_00.enc")
		sig := MustReadAll(t, "testdata/message_super_v1_00.enc.sig")
		stdout := new(bytes.Buffer)
		stderr := new(strings.Builder)

		err := cli.
			WithStdin(bytes.NewReader(encrypted)).
			WithStdout(stdout).
			WithStderr(stderr).
			Exec("decrypt", "-i", "testdata/super_v1_00.priv", "-p", "testdata/super_v1_00.pub", "--sig-hex", string(sig))
		require.NoError(t, err)
		require.Empty(t, stderr)
		require.Equal(t, "message", stdout.String())
	})
}
