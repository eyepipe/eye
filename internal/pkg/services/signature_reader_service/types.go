package signature_reader_service

import "errors"

var (
	ErrHexDecode           = errors.New("ERR_HEX_DECODE")
	ErrFileFailed          = errors.New("ERR_FILE_FAILED")
	ErrEmptySignatureInput = errors.New("ERR_EMPTY_SIGNATURE_INPUT")
)

type ReadOpts struct {
	Filename string // Filename path to filename with HEX signature
	HEX      string // HEX signature itself
}
