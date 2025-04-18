package store

import (
	"fmt"
	"time"

	"database/sql"
	"github.com/eyepipe/eye/internal/lib/pool"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	pool *pool.Pool[*sql.DB]
}

func Connect(filename string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func New(pool *pool.Pool[*sql.DB]) *Store {
	return &Store{pool: pool}
}
