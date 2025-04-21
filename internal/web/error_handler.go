package web

import (
	"errors"
	"net/http"

	"github.com/eyepipe/eye/internal/lib/crypto2"
	"github.com/eyepipe/eye/internal/lib/uuidv7"
	"github.com/eyepipe/eye/internal/pkg/store"
	"github.com/eyepipe/eye/internal/pkg/validator"
	"github.com/eyepipe/eye/pkg/proto"
	"github.com/gofiber/fiber/v3"
)

func (w *Web) ErrorHandler(c fiber.Ctx, err error) error {
	var e *fiber.Error
	var code int
	var msg string

	switch {
	case errors.As(err, &e):
		code = e.Code
		msg = e.Message
	case errors.Is(err, store.ErrNotFound):
		code = http.StatusNotFound
		msg = "record not found"
	case validator.IsInvalidErr(err):
		code = http.StatusUnprocessableEntity
		msg = err.Error()
	case errors.Is(err, uuidv7.ErrDecodeFailed):
		code = http.StatusBadRequest
		msg = "invalid UUIDv7"
	case errors.Is(err, crypto2.ErrSignatureInvalid):
		code = http.StatusBadRequest
		msg = "invalid digital signature"
	default:
		code = http.StatusInternalServerError
		msg = "some of the server code is broken"
	}

	return c.Status(code).JSON(proto.ErrorResponse{
		Code:    code,
		Status:  http.StatusText(code),
		Message: msg,
	})
}
