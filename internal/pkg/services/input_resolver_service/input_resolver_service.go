package input_resolver_service

import "resty.dev/v3"

const (
	Limit1MB = 1 << 20
)

type InputResolverService struct {
	resty *resty.Client
}

func NewService() *InputResolverService {
	return &InputResolverService{
		resty: resty.New().
			SetRedirectPolicy(resty.FlexibleRedirectPolicy(5)).
			SetResponseBodyLimit(Limit1MB),
	}
}
