package crypto2

import (
	"hash"
)

// Hashier implements IHashier
// Hash Function wrapper
type Hashier struct {
	fn func() hash.Hash
}

// NewHashier returns new Hashier
func NewHashier(fn func() hash.Hash) *Hashier {
	return &Hashier{fn: fn}
}

// New returns new hash function instance
func (e *Hashier) New() hash.Hash {
	return e.fn()
}

// ByteSize returns the number of bytes Sum will return.
func (e *Hashier) ByteSize() int {
	return e.New().Size()
}
