package size_reader

import "io"

type ISizeReader interface {
	io.Reader
	GetByteSize() int64
}
