package store

import (
	"context"
	"testing"
	"time"

	"github.com/eyepipe/eye/internal/pkg/domain"
	"github.com/shlima/oi/null"
	"github.com/stretchr/testify/require"
)

func TestStore_IncrementLimitation(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		MustWithNew(t, func(store *Store, ctx context.Context) {
			MustTruncate(t, store)

			limit := MustBuildLimitation(t, store, domain.Limitation{Date: null.NewAutoDate(time.Now())})
			err := store.IncrementLimitation(ctx, limit)
			require.NoError(t, err)

			got, err := store.FindLimitation(ctx, limit.Date.Time)
			require.NoError(t, err)
			MustCompareLimitations(t, limit, got)

			incremented := domain.Limitation{
				Date:           null.NewAutoDate(time.Now()),
				WrittenBytes:   null.NewAutoInt64(1),
				WrittenCounter: null.NewAutoInt64(2),
				ReadBytes:      null.NewAutoInt64(3),
				ReadCounter:    null.NewAutoInt64(4),
			}
			err = store.IncrementLimitation(ctx, incremented)
			require.NoError(t, err)

			limit.WrittenBytes.Int64 += incremented.WrittenBytes.Int64
			limit.WrittenCounter.Int64 += incremented.WrittenCounter.Int64
			limit.ReadBytes.Int64 += incremented.ReadBytes.Int64
			limit.ReadCounter.Int64 += incremented.ReadCounter.Int64
			got, err = store.FindLimitation(ctx, limit.Date.Time)
			require.NoError(t, err)
			MustCompareLimitations(t, limit, got)
		})
	})
}

func MustCompareLimitations(t *testing.T, a domain.Limitation, b domain.Limitation) {
	require.Equal(t, domain.FormatDate(a.Date.Time), domain.FormatDate(b.Date.Time))
	require.Equal(t, a.WrittenBytes.Int64, b.WrittenBytes.Int64)
	require.Equal(t, a.WrittenCounter.Int64, b.WrittenCounter.Int64)
	require.Equal(t, a.ReadBytes.Int64, b.ReadBytes.Int64)
	require.Equal(t, a.ReadCounter.Int64, b.ReadCounter.Int64)
}

func MustCreateLimitation(t *testing.T, store *Store, in domain.Limitation) (out domain.Limitation) {
	out = MustBuildLimitation(t, store, in)
	err := store.createLimitation(context.Background(), out)
	require.NoError(t, err)
	return out
}

func MustBuildLimitation(t *testing.T, store *Store, in domain.Limitation) (out domain.Limitation) {
	out = domain.Limitation{
		Date:           null.NewAutoTime(Seq.RandomTime()),
		WrittenBytes:   null.NewAutoInt64(Seq.NextInt64()),
		WrittenCounter: null.NewAutoInt64(Seq.NextInt64()),
		ReadBytes:      null.NewAutoInt64(Seq.NextInt64()),
		ReadCounter:    null.NewAutoInt64(Seq.NextInt64()),
		CreatedAt:      null.NewAutoTime(time.Now()),
		UpdatedAt:      null.NewAutoTime(time.Now()),
	}
	if in.Date.Valid {
		out.Date = in.Date
	}
	if in.WrittenBytes.Valid {
		out.WrittenBytes = in.WrittenBytes
	}
	if in.WrittenCounter.Valid {
		out.WrittenCounter = in.WrittenCounter
	}
	if in.ReadBytes.Valid {
		out.ReadBytes = in.ReadBytes
	}
	if in.ReadCounter.Valid {
		out.ReadCounter = in.ReadCounter
	}
	if in.CreatedAt.Valid {
		out.CreatedAt = in.CreatedAt
	}
	if in.UpdatedAt.Valid {
		out.UpdatedAt = in.UpdatedAt
	}

	return out
}
