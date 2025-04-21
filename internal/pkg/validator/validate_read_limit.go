package validator

import (
	"context"

	"github.com/eyepipe/eye/internal/pkg/domain"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ValidateReadLimitOpts struct {
	Limitation       domain.Limitation
	ReadBytesLimit   int64
	ReadCounterLimit int64
}

func (v *Validator) ValidateReadLimit(ctx context.Context, opts ValidateReadLimitOpts) error {
	return validation.Errors{
		"ReadBytes": validation.Validate(
			opts.Limitation.ReadBytes.Int64,
			validation.Max(opts.ReadBytesLimit)),
		"ReadCounter": validation.Validate(
			opts.Limitation.ReadCounter.Int64,
			validation.Max(opts.ReadCounterLimit)),
	}.Filter()
}
