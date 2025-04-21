package validator

import (
	"testing"

	"github.com/eyepipe/eye/internal/pkg/domain"
	"github.com/shlima/oi/null"
	"github.com/stretchr/testify/require"
)

func TestValidator_ValidateWriteLimit(t *testing.T) {
	t.Parallel()

	type args struct {
		desc  string
		error string
		build func() ValidateWriteLimitOpts
	}

	tests := []args{
		{
			desc: "when ok",
			build: func() ValidateWriteLimitOpts {
				opts := MustBuildValidValidateWriteLimitOpts(t)
				return opts
			},
		},
		{
			desc:  "when WriteBytesLimit limit has been ended",
			error: "WrittenBytes",
			build: func() ValidateWriteLimitOpts {
				opts := MustBuildValidValidateWriteLimitOpts(t)
				opts.Limitation.WrittenBytes.Int64 = opts.WriteBytesLimit + 1
				return opts
			},
		},
		{
			desc:  "when WriteCounterLimit limit has been ended",
			error: "WrittenCounter",
			build: func() ValidateWriteLimitOpts {
				opts := MustBuildValidValidateWriteLimitOpts(t)
				opts.Limitation.WrittenCounter.Int64 = opts.WriteCounterLimit + 1
				return opts
			},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			setup, ctx := MustNew(t)
			err := setup.ValidateWriteLimit(ctx, tt.build())
			switch tt.error {
			case "":
				require.NoError(t, err)
			default:
				require.ErrorContains(t, err, tt.error)
			}
		})
	}
}

func MustBuildValidValidateWriteLimitOpts(t *testing.T) ValidateWriteLimitOpts {
	return ValidateWriteLimitOpts{
		Limitation: domain.Limitation{
			WrittenBytes:   null.NewAutoInt64(199),
			WrittenCounter: null.NewAutoInt64(299),
		},
		WriteBytesLimit:   200,
		WriteCounterLimit: 300,
	}
}
