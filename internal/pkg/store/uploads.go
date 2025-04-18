package store

import (
	"context"
	"time"

	"github.com/eyepipe/eye/internal/pkg/domain"
	"github.com/shlima/oi/null"
)

type IUploads interface {
	FindUpload(ctx context.Context, UUID string) (out domain.Upload, err error)
	FindNotExpiredUpload(ctx context.Context, UUID string, now time.Time) (out domain.Upload, err error)
	CreateUpload(ctx context.Context, in *domain.Upload) error
	UpdateUploadByteSize(ctx context.Context, UUID string, size int64) error
	UpdateUploadSignatureHex(ctx context.Context, UUID string, signature string) error
}

func (s *Store) FindNotExpiredUpload(ctx context.Context, UUID string, now time.Time) (out domain.Upload, err error) {
	upload, err := s.FindUpload(ctx, UUID)
	switch {
	case err != nil:
		return out, err
	case now.After(upload.TTL.Time):
		return out, ErrNotFound
	default:
		return upload, nil
	}
}

func (s *Store) FindUpload(ctx context.Context, UUID string) (out domain.Upload, err error) {
	query := Query().
		Select("*").
		From("uploads").
		Where("uuid = ?", UUID).
		Limit(1)

	err = withTx(ctx, s.pool, func(ctx context.Context, tx *Tx) error {
		return Get(ctx, tx, &out, query)
	})

	return
}

func (s *Store) CreateUpload(ctx context.Context, in *domain.Upload) error {
	in.CreatedAt = null.NewAutoTime(time.Now())
	in.UpdatedAt = null.NewAutoTime(time.Now())

	query := Query().
		Insert("uploads").
		SetMap(in.Attributes())

	return withTx(ctx, s.pool, func(ctx context.Context, tx *Tx) error {
		return Exec(ctx, tx, query)
	})
}

func (s *Store) UpdateUploadByteSize(ctx context.Context, UUID string, size int64) error {
	query := Query().
		Update("uploads").Set("byte_size", size).
		Where("uuid = ?", UUID).
		Suffix(`RETURNING uuid`)

	return withTx(ctx, s.pool, func(ctx context.Context, tx *Tx) error {
		return Get(ctx, tx, new(string), query)
	})
}

func (s *Store) UpdateUploadSignatureHex(ctx context.Context, UUID string, signature string) error {
	query := Query().
		Update("uploads").Set("signature_hex", signature).
		Where("uuid = ?", UUID).
		Suffix(`RETURNING uuid`)

	return withTx(ctx, s.pool, func(ctx context.Context, tx *Tx) error {
		return Get(ctx, tx, new(string), query)
	})
}
