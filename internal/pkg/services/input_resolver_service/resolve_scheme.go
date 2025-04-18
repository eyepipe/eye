package input_resolver_service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eyepipe/eye/internal/pkg/container"
)

func (i *InputResolverService) ResolveSchemeURLOrJSON(ctx context.Context, input string) (out container.Scheme, err error) {
	switch IsHttpURL(input) {
	case true:
		return i.resolveSchemeURL(ctx, input)
	default:
		return i.resolveSchemeJSON(ctx, input)
	}
}

func (i *InputResolverService) resolveSchemeURL(ctx context.Context, input string) (out container.Scheme, err error) {
	res, err := i.resty.R().
		SetContext(ctx).
		SetResult(&out).
		Get(input)
	switch {
	case err != nil:
		return out, fmt.Errorf("%w: %w", ErrHttpFailed, err)
	case !res.IsSuccess():
		return out, fmt.Errorf("%w: <%s> %s", ErrHttpBadStatusCode, res.Status(), input)
	default:
		return out, nil
	}
}

func (i *InputResolverService) resolveSchemeJSON(ctx context.Context, input string) (out container.Scheme, err error) {
	err = json.Unmarshal([]byte(input), &out)
	if err != nil {
		return out, fmt.Errorf("%w: %w", ErrJSONParse, err)
	}

	return out, nil
}
