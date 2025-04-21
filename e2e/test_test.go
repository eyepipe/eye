package e2e

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	cliExecPath = "./testdata/eye-cli-test.bin"
)

type Setup struct {
	cmd    string
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

func (s *Setup) WithStdin(r io.Reader) *Setup {
	s.stdin = r
	return s
}

func (s *Setup) WithStdout(w io.Writer) *Setup {
	s.stdout = w
	return s
}

func (s *Setup) WithStderr(w io.Writer) *Setup {
	s.stderr = w
	return s
}

func (s *Setup) Exec(args ...string) error {
	cmd := exec.Command(s.cmd, args...)
	cmd.Stdin = s.stdin
	cmd.Stdout = s.stdout
	cmd.Stdout = s.stdout
	cmd.Stderr = s.stderr
	return cmd.Run()
}

func NewSetup(cmd string) *Setup {
	return &Setup{cmd: cmd}
}

func TestMain(m *testing.M) {
	setup := NewSetup("go").WithStderr(os.Stderr)
	err := setup.Exec("build", "-o", cliExecPath, "../cmd/cli")
	if err != nil {
		panic(fmt.Errorf("failed to build app: %w", err))
	}

	os.Exit(m.Run())
}

func MustSetupCli(t *testing.T) (*Setup, context.Context) {
	return NewSetup(cliExecPath), context.Background()
}

func MustReadAll(t *testing.T, filepath string) []byte {
	got, err := os.ReadFile(filepath)
	require.NoError(t, err)
	return got
}

func MustBeAnExistError(t *testing.T, err error) {
	require.Error(t, err)
	require.IsType(t, &exec.ExitError{}, err)
}
