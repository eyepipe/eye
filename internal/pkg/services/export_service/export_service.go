package export_service

import (
	"encoding/pem"
	"fmt"
	"io"

	"github.com/eyepipe/eye/internal/pkg/container"
	"github.com/samber/lo"
)

type ExportService struct {
	c        *container.Container
	handlers []IHandler
}

func NewExportService(c *container.Container) *ExportService {
	return &ExportService{
		c: c,
		handlers: []IHandler{
			NewHandler(PEMSchemeProto, false, ExportScheme, ImportScheme),
			NewHandler(PEMSignerPrivateKey, true, ExportSignerPrivate, ImportSignerPrivate),
			NewHandler(PEMSignerPublicKey, false, ExportSignerPublic, ImportSignerPublic),
			NewHandler(PEMKeyAgreementPrivateKey, true, ExportKeyAgreementPrivate, ImportKeyAgreementPrivate),
			NewHandler(PEMKeyAgreementPublicKey, false, ExportKeyAgreementPublic, ImportKeyAgreementPublic),
		},
	}
}

func (s *ExportService) ExportPrivate(w io.Writer) error {
	return s.export(w, true)
}

func (s *ExportService) ExportPublic(w io.Writer) error {
	return s.export(w, false)
}

func (s *ExportService) Import(input []byte) error {
	index := lo.KeyBy(s.handlers, func(x IHandler) string {
		return x.GetPEMBlock()
	})

	for {
		var block *pem.Block
		block, input = pem.Decode(input)
		if block == nil {
			break
		}

		handler, ok := index[block.Type]
		if !ok {
			return fmt.Errorf("%s: <%s>", ErrPemBlockUnknown, block.Type)
		}

		err := handler.Import(s.c, block.Bytes)
		if err != nil {
			return fmt.Errorf("%w: failed to import <%s> block: %w", ErrPemParse, block.Type, err)
		}
	}

	return nil
}

func (s *ExportService) export(w io.Writer, isPrivate bool) error {
	for i := range s.handlers {
		handler := s.handlers[i]
		if !isPrivate && handler.GetIsPrivate() {
			continue
		}

		bytea, err := handler.Export(s.c)
		if err != nil {
			return fmt.Errorf("%w: <%s> %w", ErrPemExport, handler.GetPEMBlock(), err)
		}

		err = pem.Encode(w, &pem.Block{Type: handler.GetPEMBlock(), Bytes: bytea})
		if err != nil {
			return fmt.Errorf("%w: <%s> %w", ErrPemEncode, handler.GetPEMBlock(), err)
		}
	}

	return nil
}
