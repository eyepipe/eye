package size_reader

import (
	"fmt"
	"io"
	"sync/atomic"
)

// SizeReader implements io.Reader
type SizeReader struct {
	r     io.Reader
	bytes int64
}

// New returns SizeReader
func New() *SizeReader {
	return &SizeReader{}
}

func (s *SizeReader) Wrap(r io.Reader) *SizeReader {
	s.r = r
	return s
}

func NewSizeReader(r io.Reader) *SizeReader {
	return &SizeReader{r: r}
}

func (s *SizeReader) Read(p []byte) (n int, err error) {
	if s.r == nil {
		return 0, fmt.Errorf("underling sized reader is nil")
	}

	n, err = s.r.Read(p)
	atomic.AddInt64(&s.bytes, int64(n))
	return n, err
}

// GetByteSize return read bytes size
func (s *SizeReader) GetByteSize() int64 {
	return s.bytes
}
