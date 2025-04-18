package input_resolver_service

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/eyepipe/eye/internal/pkg/container"
	"github.com/eyepipe/eye/internal/pkg/services/export_service"
)

func (i *InputResolverService) ResolveContainerReaderOrFileOrUrl(ctx context.Context, reader io.Reader, input string) (c *container.Container, err error) {
	var bytea []byte

	if IsHttpURL(input) {
		res, err := i.resty.R().
			SetContext(ctx).
			Get(input)
		switch {
		case err != nil:
			return nil, fmt.Errorf("failed to http GET <%s>: %w", input, err)
		case !res.IsSuccess():
			return nil, fmt.Errorf("bad status code: <%s>", res.Status())
		default:
			bytea = res.Bytes()
		}
	} else if len(input) > 0 {
		bytea, err = os.ReadFile(input)
		if err != nil {
			return nil, fmt.Errorf("failed to os.ReadFile: %w", err)
		}
	} else {
		bytea, err = io.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("failed to io.ReadAll: %w", err)
		}
	}

	c = container.NewEmptyContainer()
	service := export_service.NewExportService(c)
	err = service.Import(bytea)
	if err != nil {
		return nil, fmt.Errorf("failed to import key: %w", err)
	}

	return c, nil
}
