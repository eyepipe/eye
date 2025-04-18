package domain

import (
	"github.com/eyepipe/eye/internal/lib/crypto2"
	"github.com/shlima/oi/db"
	"github.com/shlima/oi/null"
)

type Upload struct {
	UUID         null.String `db:"uuid"`
	SignerAlgo   null.String `db:"signer_algo"`
	S3Key        null.String `db:"s3_key"`
	S3Urn        null.String `db:"s3_urn"`
	ByteSize     null.Int64  `db:"byte_size"`
	TTL          null.Time   `db:"ttl"`
	SignatureHex null.String `db:"signature_hex"`
	CreatedAt    null.Time   `db:"created_at"`
	UpdatedAt    null.Time   `db:"updated_at"`
}

func (u *Upload) GetSignerAlgo() crypto2.SignerAlgo {
	return crypto2.SignerAlgo(u.SignerAlgo.String)
}

// Attributes returns database columns with values
func (u *Upload) Attributes() db.Attributes {
	return db.Attributes{
		"uuid":          u.UUID,
		"signer_algo":   u.SignerAlgo,
		"s3_key":        u.S3Key,
		"s3_urn":        u.S3Urn,
		"byte_size":     u.ByteSize,
		"signature_hex": u.SignatureHex,
		"ttl":           u.TTL,
		"created_at":    u.CreatedAt,
		"updated_at":    u.UpdatedAt,
	}
}
