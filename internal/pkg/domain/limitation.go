package domain

import (
	"github.com/samber/lo"
	"github.com/shlima/oi/db"
	"github.com/shlima/oi/null"
)

type Limitation struct {
	Date           null.Date  `db:"date"`
	WrittenBytes   null.Int64 `db:"written_bytes"`
	WrittenCounter null.Int64 `db:"written_counter"`
	ReadBytes      null.Int64 `db:"read_bytes"`
	ReadCounter    null.Int64 `db:"read_counter"`
	CreatedAt      null.Time  `db:"created_at"`
	UpdatedAt      null.Time  `db:"updated_at"`
}

// Attributes returns database columns with values
func (l *Limitation) Attributes() db.Attributes {
	date := lo.Ternary(
		l.Date.Valid,
		null.String{Valid: true, String: FormatDate(l.Date.Time)},
		null.String{},
	)

	return db.Attributes{
		"date":            date,
		"written_bytes":   l.WrittenBytes,
		"written_counter": l.WrittenCounter,
		"read_bytes":      l.ReadBytes,
		"read_counter":    l.ReadCounter,
	}
}
