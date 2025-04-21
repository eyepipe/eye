package encryption_service

import (
	"fmt"
	"hash"
	"io"

	"github.com/eyepipe/eye/internal/lib/crypto2"
	"github.com/eyepipe/eye/internal/pkg/container"
)

type Envelope struct {
	Ver     uint32
	KeySalt []byte
	IV      []byte
}

type Service struct {
	i          *container.Container // identity
	r          *container.Container // recipient
	key        *crypto2.KeyDerivationResult
	signerHash hash.Hash
	envelope   *Envelope // decrypter
}

func NewService(i, r *container.Container) *Service {
	return &Service{
		i:        i,
		r:        r,
		envelope: new(Envelope),
	}
}

func (s *Service) GetI() *container.Container {
	return s.i
}

func (s *Service) GetP() *container.Container {
	return s.r
}

func (s *Service) WrapSigner(writer io.Writer) io.Writer {
	signer := s.i.Scheme.SignAlgo.ToSigner()
	s.signerHash = signer.GetHashier().New()
	return io.MultiWriter(writer, s.signerHash)
}

func (s *Service) WrapVerifier(writer io.Writer) io.Writer {
	signer := s.r.Scheme.SignAlgo.ToSigner()
	s.signerHash = signer.GetHashier().New()
	return io.MultiWriter(writer, s.signerHash)
}

func (s *Service) Signature() ([]byte, error) {
	signer := s.i.Scheme.SignAlgo.ToSigner()
	return signer.Sign(s.signerHash.Sum(nil), s.i.SignerPair.Private)
}

func (s *Service) Verification() []byte {
	return s.signerHash.Sum(nil)
}

func (s *Service) DeriveKey() error {
	shared, err := s.i.Scheme.KeyAgreementAlgo.
		ToKeyAgreement().
		ComputeSharedPrivatePublic(s.i.KeyAgreementPair.Private, s.r.KeyAgreementPair.Public)
	if err != nil {
		return fmt.Errorf("failed to compute shared secret: %w", err)
	}

	cipher := s.r.Scheme.BlockCipherAlgo.ToCipher()
	derivation := s.r.Scheme.HashKeyDerivationAlgo.ToDerivation()
	result, err := derivation.Derive(shared, s.envelope.KeySalt, cipher.GetKeySizeBytes())
	if err != nil {
		return fmt.Errorf("failed to derive key: %w", err)
	}

	s.key = result
	return nil
}

func (s *Service) WarpEncryptor(writer io.Writer) (io.Writer, error) {
	cipher := s.r.Scheme.BlockCipherAlgo.ToCipher()
	out, iv, err := cipher.NewEncryptor(s.key.Key, writer)
	if err != nil {
		return nil, fmt.Errorf("failed to initilize a cipher: %w", err)
	}

	err = PutByte(writer, ByteVer)
	if err != nil {
		return nil, fmt.Errorf("failed to put buildinfo byte: %w", err)
	}
	err = PutUint32(writer, Ver1)
	if err != nil {
		return nil, fmt.Errorf("failed to put buildinfo number: %w", err)
	}

	err = PutByteWithData(writer, ByteSalt, s.key.Salt)
	if err != nil {
		return nil, fmt.Errorf("failed to write salt: %w", err)
	}

	err = PutByteWithData(writer, ByteIV, iv)
	if err != nil {
		return nil, fmt.Errorf("failed to write iv: %w", err)
	}

	err = PutByte(writer, ByteBody)
	if err != nil {
		return nil, fmt.Errorf("failed to put body byte: %w", err)
	}

	return out, nil
}

func (s *Service) ReadEnvelope(reader io.Reader) (err error) {
	var verMarker byte
	err = GetByte(reader, &verMarker)
	switch {
	case err != nil:
		return fmt.Errorf("failed to read buildinfo byte: %w", err)
	case verMarker != ByteVer:
		return fmt.Errorf("expected buildinfo marker, expected <%v> got: <%v>", ByteVer, verMarker)
	}

	var ver uint32
	err = GetUint32(reader, &ver)
	switch {
	case err != nil:
		return fmt.Errorf("failed to read buildinfo: %w", err)
	case ver != Ver1:
		return fmt.Errorf("invalid buildinfo: %d", ver)
	}

	var (
		salt []byte
		iv   []byte
	)
LOOP:
	for {
		var marker byte
		err = GetByte(reader, &marker)
		if err != nil {
			return fmt.Errorf("failed to read marker: %w", err)
		}

		switch marker {
		case ByteSalt:
			salt, err = GetData(reader)
		case ByteIV:
			iv, err = GetData(reader)
		case ByteBody:
			break LOOP
		}
	}

	s.envelope = &Envelope{
		Ver:     ver,
		KeySalt: salt,
		IV:      iv,
	}

	return nil
}

func (s *Service) WarpDecrypter(reader io.Reader) (io.Reader, error) {
	cipher := s.r.Scheme.BlockCipherAlgo.ToCipher()
	out, err := cipher.NewDecrypter(s.key.Key, s.envelope.IV, reader)
	if err != nil {
		return nil, fmt.Errorf("failed to initilize a cipher: %w", err)
	}

	return out, nil
}
