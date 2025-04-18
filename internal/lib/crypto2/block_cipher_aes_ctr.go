package crypto2

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"fmt"
	"io"
)

type BlockCipherAesCtr struct {
	keySizeBytes int
	rand         io.Reader
}

func NewBlockCipherAesCtr(keySizeBytes int) *BlockCipherAesCtr {
	return &BlockCipherAesCtr{
		keySizeBytes: keySizeBytes,
		rand:         rand.Reader,
	}
}

func (b *BlockCipherAesCtr) GetKeySizeBytes() int {
	return b.keySizeBytes
}

func (b *BlockCipherAesCtr) NewEncryptor(key []byte, writer io.Writer) (w io.Writer, iv []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to new aes cipher: %w", err)
	}

	iv = make([]byte, block.BlockSize())
	_, err = rand.Read(iv)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate iv: %w", err)
	}

	stream := cipher.NewCTR(block, iv)
	return &cipher.StreamWriter{S: stream, W: writer}, iv, nil
}

func (b *BlockCipherAesCtr) NewDecrypter(key, iv []byte, reader io.Reader) (r io.Reader, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to new aes cipher: %w", err)
	}

	stream := cipher.NewCTR(block, iv)
	return &cipher.StreamReader{S: stream, R: reader}, nil
}
