package validator

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func IsInvalidErr(err error) bool {
	return errors.As(err, &validation.Errors{})
}
