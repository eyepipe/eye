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
var dbPool *pool.Pool[*Store]

func TestMain(m *testing.M) {
	paths := []string{
		"../../../db/databases/db00.sqlite",
		"../../../db/databases/db01.sqlite",
		"../../../db/databases/db02.sqlite",
		"../../../db/databases/db03.sqlite",
		"../../../db/databases/db04.sqlite",
		"../../../db/databases/db05.sqlite",
		"../../../db/databases/db06.sqlite",
		"../../../db/databases/db07.sqlite",
		"../../../db/databases/db08.sqlite",
		"../../../db/databases/db09.sqlite",
		"../../../db/databases/db10.sqlite",
	}

	stores := make([]*Store, 0, len(paths))
	for _, path := range paths {
		db, err := Connect(path)
		if err != nil {
			panic(fmt.Errorf("failed to connect to database <%s>: %w", path, err))
		}
		stores = append(stores, New(pool.New([]*sql.DB{db})))
	}

	dbPool = pool.New(stores)
	os.Exit(m.Run())
}

func MustWithNew(t *testing.T, cb func(store *Store, ctx context.Context)) {
	err := dbPool.GetE(context.Background(), func(store *Store) error {
		cb(store, context.Background())
		return nil
	})

	require.NoError(t, err)
}

func MustTruncate(t *testing.T, store *Store) {
	err := withTx(context.Background(), store.pool, func(ctx context.Context, tx *Tx) error {
		tables := []string{
			"uploads",
			"limitations",
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
