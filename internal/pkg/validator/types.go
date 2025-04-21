package validator

import "context"

type IValidator interface {
	ValidateWriteLimit(ctx context.Context, opts ValidateWriteLimitOpts) error
	ValidateReadLimit(ctx context.Context, opts ValidateReadLimitOpts) error
}
