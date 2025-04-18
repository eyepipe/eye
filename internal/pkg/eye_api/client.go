package eye_api

import (
	"context"
	"fmt"
	"github.com/eyepipe/eye/pkg/proto"
	"io"
	"resty.dev/v3"
)

type Client struct {
	r *resty.Client
}

func NewClient() *Client {
	return &Client{
		r: resty.New().SetRedirectPolicy(resty.FlexibleRedirectPolicy(10)),
	}
}

type UploadOpts struct {
	SignerAlgo string
}

func (c *Client) Upload(ctx context.Context, URL string, r io.Reader, opts UploadOpts) (out *proto.CreateUploadResponseV1, err error) {
	out = new(proto.CreateUploadResponseV1)
	res, err := c.r.R().
		SetContext(ctx).
		SetContentLength(false).
		SetResult(out).
		SetHeader("Content-Type", "application/octet-stream").
		SetHeader("Transfer-Encoding", "chunked").
		SetHeader("x-signer-algo", opts.SignerAlgo).
		SetBody(r).
		Post(URL)

	switch {
	case err != nil:
		return nil, err
	case !res.IsSuccess():
		return nil, fmt.Errorf("bad status code: %d, %s", res.StatusCode(), res.String())
	default:
		return out, nil
	}
}

func (c *Client) Confirm(ctx context.Context, URL string, req *proto.ConfirmUploadRequestV1) (out *proto.ConfirmUploadResponseV1, err error) {
	out = new(proto.ConfirmUploadResponseV1)
	res, err := c.r.R().
		SetBody(req).
		SetResult(out).
		SetContext(ctx).
		Post(URL)

	switch {
	case err != nil:
		return nil, err
	case !res.IsSuccess():
		return nil, fmt.Errorf("bad status code: %d, %s", res.StatusCode(), res.String())
	default:
		return out, nil
	}
}
