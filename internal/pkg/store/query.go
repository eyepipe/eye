package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/eyepipe/eye/internal/lib/pool"
	"github.com/georgysavva/scany/v2/sqlscan"
)

type (
	Tx      = sql.Tx
	Sqlizer = squirrel.Sqlizer
)

var (
	ErrSqlBuilder = errors.New("ERR_SQL_BUILDER")
	ErrNotFound   = sql.ErrNoRows
)

func Query() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
}

// Get вернет строку, если строки нет то вернет ошибку
func Get(ctx context.Context, tx *Tx, dest interface{}, query Sqlizer) error {
	stmt, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrSqlBuilder, err)
	}

	rows, err := tx.QueryContext(ctx, stmt, args...)
	switch {
	case err != nil:
		return err
	case rows.Err() != nil:
		return rows.Err()
	}

	return sqlscan.ScanOne(dest, rows)
}

func Exec(ctx context.Context, tx *Tx, query Sqlizer) error {
	stmt, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrSqlBuilder, err)
	}

	_, err = tx.ExecContext(ctx, stmt, args...)
	return err
}

func withTx(ctx context.Context, pool *pool.Pool[*sql.DB], cb func(ctx context.Context, tx *Tx) error) error {
	return pool.GetE(ctx, func(db *sql.DB) error {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return fmt.Errorf("failed to begin tx: %w", err)
		}

		defer func() {
			err = tx.Rollback()
		}()

		err = cb(ctx, tx)
		if err != nil {
			_ = tx.Rollback() // ignore rollback error as there is already an error to return
			return err
		}

		return tx.Commit()
	})
}
