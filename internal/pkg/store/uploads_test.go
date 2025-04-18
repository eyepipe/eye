package store

import (
	"context"
	"github.com/eyepipe/eye/internal/lib/uuidv7"
	"testing"
	"time"

	"github.com/eyepipe/eye/internal/pkg/domain"
	"github.com/shlima/oi/null"
	"github.com/stretchr/testify/require"
)

func TestStore_CreateUpload(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		MustWithNew(t, func(store *Store, ctx context.Context) {
			MustTruncate(t, store)
			upload := MustBuildUpload(t, store, domain.Upload{})
			err := store.CreateUpload(ctx, &upload)
			require.NoError(t, err)

			got, err := store.FindUpload(ctx, upload.UUID.String)
			require.NoError(t, err)
			MustCompareUploads(t, upload, got)
		})
	})
}

func MustCompareUploads(t *testing.T, a, b domain.Upload) {
	require.Equal(t, a.UUID, b.UUID)
	require.Equal(t, a.SignerAlgo, b.SignerAlgo)
	require.Equal(t, a.S3Key, b.S3Key)
	require.Equal(t, a.S3Urn, b.S3Urn)
	require.Equal(t, a.ByteSize, b.ByteSize)
	require.Equal(t, a.TTL.Time.UTC(), b.TTL.Time.UTC())
	require.Equal(t, a.SignatureHex, b.SignatureHex)
}

func MustBuildUpload(t *testing.T, store *Store, in domain.Upload) (out domain.Upload) {
	out = domain.Upload{
		UUID:       null.NewAutoString(uuidv7.NewWithShard(0).String()),
		SignerAlgo: null.NewAutoString("SignerAlgo"),
		S3Key:      null.NewAutoString("S3Key"),
		S3Urn:      null.NewAutoString("S3Urn"),
		TTL:        null.NewAutoTime(time.Now().AddDate(0, 0, 1)),
	}
	if in.UUID.Valid {
		out.UUID = in.UUID
	}
	if in.SignerAlgo.Valid {
		out.SignerAlgo = in.SignerAlgo
	}
	if in.S3Key.Valid {
		out.S3Key = in.S3Key
	}
	if in.S3Urn.Valid {
		out.S3Urn = in.S3Urn
	}
	if in.TTL.Valid {
		out.TTL = in.TTL
	}

	return out
}
