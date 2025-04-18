package input_resolver_service

import "github.com/pkg/errors"

var (
	ErrJSONParse         = errors.New("ERR_JSON_PARSE")
	ErrHttpFailed        = errors.New("ERR_HTTP_FAILED")
	ErrHttpBadStatusCode = errors.New("ERR_HTTP_BAD_STATUS_CODE")
)
