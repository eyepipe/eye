package validator

import (
	"errors"
	"fmt"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/require"
)

func TestIsInvalidErr(t *testing.T) {
	t.Parallel()

	type args struct {
		desc string
		err  error
		want bool
	}

	tests := []args{
		{
			desc: "is false for custom error",
			err:  errors.New("custom"),
			want: false,
		}, {
			desc: "is false for nil",
			err:  nil,
			want: false,
		}, {
			desc: "is true for validation.Errors",
			err:  validation.Errors{},
			want: true,
		}, {
			desc: "is true for wrapped validation.Errors",
			err:  fmt.Errorf("wrapped %w", validation.Errors{}),
			want: true,
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()

			got := IsInvalidErr(tt.err)
			require.Equal(t, tt.want, got)
		})
	}
}
