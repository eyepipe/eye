package web

import (
	"fmt"
	"github.com/eyepipe/eye/internal/lib/uuidv7"
)

func GenS3KeyByUUID(uuid *uuidv7.UUIDv7) string {
	text := uuid.String()

	return fmt.Sprintf(
		"%d/%02d/%02d/%s/%s/%s",
		uuid.Time.Year(),
		uuid.Time.Month(),
		uuid.Time.Day(),
		text[24:26],
		text[26:29],
		text,
	)
}
