package validator

import (
	"context"

	"github.com/eyepipe/eye/internal/pkg/domain"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ValidateWriteLimitOpts struct {
	Limitation        domain.Limitation
	WriteBytesLimit   int64
	WriteCounterLimit int64
}

func (v *Validator) ValidateWriteLimit(ctx context.Context, opts ValidateWriteLimitOpts) error {
	return validation.Errors{
		"WrittenBytes": validation.Validate(
			opts.Limitation.WrittenBytes.Int64,
			validation.Max(opts.WriteBytesLimit)),
		"WrittenCounter": validation.Validate(
			opts.Limitation.WrittenCounter.Int64,
			validation.Max(opts.WriteCounterLimit)),
	}.Filter()
}
