package e2e

import (
	"bytes"
	"encoding/hex"
	"encoding/pem"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	ValidSignatureText   = "valid signature ✅"
	InvalidSignatureText = "invalid signature ❌"
)

func TestCliKeygen(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		cli, _ := MustSetupCli(t)
		stdout := new(strings.Builder)
		err := cli.WithStdout(stdout).Exec("keygen")
		require.NoError(t, err)
		require.Contains(t, stdout.String(), "-----BEGIN SCHEME PROTO-----")
		require.Contains(t, stdout.String(), "-----BEGIN SIGNER PRIVATE KEY-----")
		require.Contains(t, stdout.String(), "-----BEGIN SIGNER PUBLIC KEY-----")
		require.Contains(t, stdout.String(), "-----BEGIN KEY AGREEMENT PRIVATE KEY-----")
		require.Contains(t, stdout.String(), "-----BEGIN KEY AGREEMENT PUBLIC KEY-----")

		var block *pem.Block
		der := []byte(stdout.String())
		for {
			block, der = pem.Decode(der)
			if block == nil {
				break
			}
			require.NotEmpty(t, block.Type)
			require.NotEmpty(t, block.Bytes)
		}
	})
}

func TestCliVerify(t *testing.T) {
	t.Parallel()

	t.Run("it works from sig file", func(t *testing.T) {
		t.Parallel()
		cli, _ := MustSetupCli(t)
		signed := MustReadAll(t, "testdata/super_v1_00.pub")

		stdout := new(strings.Builder)
		err := cli.
			WithStdout(stdout).
			WithStdin(bytes.NewReader(signed)).
			WithStderr(os.Stderr).
			Exec("verify", "-p", "testdata/super_v1_00.pub", "-sig", "testdata/super_v1_00.pub.sig")

		require.NoError(t, err)
		require.Contains(t, stdout.String(), ValidSignatureText)
	})

	t.Run("it works from sig hex", func(t *testing.T) {
		t.Parallel()
		cli, _ := MustSetupCli(t)
		signed := MustReadAll(t, "testdata/super_v1_00.pub")
		signature := MustReadAll(t, "testdata/super_v1_00.pub.sig")

		stdout := new(strings.Builder)
		err := cli.
			WithStdout(stdout).
			WithStdin(bytes.NewReader(signed)).
			WithStderr(os.Stderr).
			Exec("verify", "-p", "testdata/super_v1_00.pub", "-sig-hex", string(signature))

		require.NoError(t, err)
		require.Contains(t, stdout.String(), ValidSignatureText)
	})

	t.Run("it errors if wrong sig hex", func(t *testing.T) {
		t.Parallel()
		cli, _ := MustSetupCli(t)
		signed := MustReadAll(t, "testdata/super_v1_00.pub")

		stderr := new(strings.Builder)
		err := cli.
			WithStdin(bytes.NewReader(signed)).
			WithStderr(stderr).
			Exec("verify", "-p", "testdata/super_v1_00.pub", "-sig-hex", hex.EncodeToString([]byte{1}))

		MustBeAnExistError(t, err)
		require.Contains(t, stderr.String(), InvalidSignatureText)
	})

	t.Run("it errors when sig file not found", func(t *testing.T) {
		t.Parallel()
		cli, _ := MustSetupCli(t)
		signed := MustReadAll(t, "testdata/super_v1_00.pub")

		stderr := new(strings.Builder)
		err := cli.
			WithStdin(bytes.NewReader(signed)).
			WithStderr(stderr).
			Exec("verify", "-p", "testdata/super_v1_00.pub", "-sig", "testdata/foo")

		MustBeAnExistError(t, err)
		require.Contains(t, stderr.String(), "ERR_FILE_FAILED")
	})
}
