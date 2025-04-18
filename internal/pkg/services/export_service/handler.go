package export_service

import (
	"github.com/eyepipe/eye/internal/pkg/container"
)

type IHandler interface {
	GetPEMBlock() string
	GetIsPrivate() bool
	Export(c *container.Container) ([]byte, error)
	Import(c *container.Container, bytea []byte) error
}

type Handler struct {
	PEMBlock  string
	IsPrivate bool
	ExportFn  func(c *container.Container) ([]byte, error)
	ImportFn  func(c *container.Container, bytea []byte) error
}

func NewHandler(
	PEMBlock string,
	IsPrivate bool,
	exportFn func(c *container.Container) ([]byte, error),
	importFn func(c *container.Container, bytea []byte) error,
) *Handler {
	return &Handler{
		PEMBlock:  PEMBlock,
		IsPrivate: IsPrivate,
		ExportFn:  exportFn,
		ImportFn:  importFn,
	}
}

func (h *Handler) GetPEMBlock() string {
	return h.PEMBlock
}

func (h *Handler) GetIsPrivate() bool {
	return h.IsPrivate
}

func (h *Handler) Export(c *container.Container) ([]byte, error) {
	return h.ExportFn(c)
}

func (h *Handler) Import(c *container.Container, bytea []byte) error {
	return h.ImportFn(c, bytea)
}
