package domain

import "time"

func FormatDate(t time.Time) string {
	return t.Format(time.DateOnly)
}
