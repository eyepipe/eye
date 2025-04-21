package store

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/eyepipe/eye/internal/pkg/domain"
	"github.com/samber/lo"
)

type ILimitations interface {
	FindLimitation(ctx context.Context, date time.Time) (out domain.Limitation, err error)
	FetchLimitation(ctx context.Context, date time.Time) (out domain.Limitation, err error)
	IncrementLimitation(ctx context.Context, limitation domain.Limitation) error
}

func (s *Store) FetchLimitation(ctx context.Context, date time.Time) (out domain.Limitation, err error) {
	got, err := s.FindLimitation(ctx, date)
	switch {
	case errors.Is(err, ErrNotFound):
		return out, nil
	case err != nil:
		return out, fmt.Errorf("failed to find limitation: %w", err)
	default:
		return got, nil
	}
}

func (s *Store) FindLimitation(ctx context.Context, date time.Time) (out domain.Limitation, err error) {
	query := Query().
		Select("*").
		From("limitations").
		Where("date = ?", domain.FormatDate(date)).
		Limit(1)

	err = withTx(ctx, s.pool, func(ctx context.Context, tx *Tx) error {
		return Get(ctx, tx, &out, query)
	})
	return
}

func (s *Store) IncrementLimitation(ctx context.Context, limitation domain.Limitation) error {
	if !limitation.Date.Valid {
		return fmt.Errorf("limitation date must be set")
	}

	conflictUpdates := []string{
		"written_bytes = COALESCE(written_bytes, 0) + EXCLUDED.written_bytes",
		"written_counter = COALESCE(written_counter, 0) + EXCLUDED.written_counter",
		"read_bytes = COALESCE(read_bytes, 0) + EXCLUDED.read_bytes",
		"read_counter = COALESCE(read_counter, 0) + EXCLUDED.read_counter",
		"updated_at = CURRENT_TIMESTAMP",
	}

	query := Query().
		Insert("limitations").
		SetMap(lo.Assign(limitation.Attributes(), map[string]any{"created_at": "CURRENT_TIMESTAMP"})).
		Suffix(fmt.Sprintf("ON CONFLICT (date) DO UPDATE SET %s", strings.Join(conflictUpdates, ", ")))

	return withTx(ctx, s.pool, func(ctx context.Context, tx *Tx) error {
		return Exec(ctx, tx, query)
	})
}

func (s *Store) createLimitation(ctx context.Context, limitation domain.Limitation) error {
	query := Query().
		Insert("limitations").
		SetMap(limitation.Attributes())

	return withTx(ctx, s.pool, func(ctx context.Context, tx *Tx) error {
		return Exec(ctx, tx, query)
	})
}
