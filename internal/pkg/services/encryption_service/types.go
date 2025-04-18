package encryption_service

import (
	"github.com/eyepipe/eye/internal/pkg/container"
	"io"
)

type IService interface {
	WrapSigner(writer io.Writer) io.Writer
	WrapVerifier(writer io.Writer) io.Writer
	Signature() ([]byte, error)
	Verification() []byte
	DeriveKey() error
	WarpEncryptor(writer io.Writer) (io.Writer, error)
	ReadEnvelope(reader io.Reader) (err error)
	WarpDecrypter(reader io.Reader) (io.Reader, error)
	GetI() *container.Container
	GetP() *container.Container
}
