package store

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/eyepipe/eye/internal/lib/pool"
	"github.com/shlima/oi/sequence"
	"github.com/stretchr/testify/require"
)

var Seq = sequence.New(time.Now().UnixNano())
var shards *pool.Pool[*Store]

func TestMain(m *testing.M) {
	db, err := Connect("../../../db/db01.sqlite")
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	shards = pool.New([]*Store{
		New(pool.New([]*sql.DB{db})),
	})
	os.Exit(m.Run())
}

func MustWithNew(t *testing.T, cb func(store *Store, ctx context.Context)) {
	err := shards.GetE(context.Background(), func(store *Store) error {
		cb(store, context.Background())
		return nil
	})

	require.NoError(t, err)
}

func MustTruncate(t *testing.T, store *Store) {
	err := withTx(context.Background(), store.pool, func(ctx context.Context, tx *Tx) error {
		tables := []string{
			"uploads",
		}
		for i := range tables {
			if err := Exec(ctx, tx, Query().Delete(tables[i])); err != nil {
				return err
			}
		}

		return nil
	})

	require.NoError(t, err)
}
