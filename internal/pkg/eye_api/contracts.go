package eye_api

import (
	"context"
	"fmt"

	"github.com/eyepipe/eye/pkg/proto"
)

func (c *Client) GetContract(ctx context.Context, URL string) (out *proto.ContractV1, err error) {
	out = new(proto.ContractV1)
	res, err := c.r.R().
		SetContext(ctx).
		SetResult(out).
		Get(URL)

	switch {
	case err != nil:
		return nil, err
	case !res.IsSuccess():
		return nil, fmt.Errorf("bad status code: %d, %s", res.StatusCode(), res.String())
	default:
		return out, nil
	}
}
