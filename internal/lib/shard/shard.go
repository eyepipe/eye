package shard

import (
	"fmt"

	"github.com/samber/lo"
)

type Shard[T any] struct {
	Number uint16
	Item   T
}

type Shards[T any] struct {
	items []*Shard[T]
}

// NewShard creates new shard
func NewShard[T any](number uint16, item T) *Shard[T] {
	return &Shard[T]{
		Number: number,
		Item:   item,
	}
}

// Add item to shard
func (s *Shards[T]) Add(item T) *Shards[T] {
	s.items = append(s.items, &Shard[T]{
		Number: uint16(len(s.items)),
		Item:   item,
	})

	return s
}

func (s *Shards[T]) All() []T {
	return lo.Map(s.items, func(item *Shard[T], index int) T {
		return item.Item
	})
}

// NewShards creates new shards
func NewShards[T any]() *Shards[T] {
	return &Shards[T]{
		items: make([]*Shard[T], 0),
	}
}

// Sample returns random shard
func (s *Shards[T]) Sample() (T, uint16) {
	shard := lo.Sample(s.items)
	return shard.Item, shard.Number
}

// Get returns the shard by its number
func (s *Shards[T]) Get(ix uint16) (out T, err error) {
	switch {
	case len(s.items) > int(ix):
		return s.items[ix].Item, nil
	default:
		return out, fmt.Errorf("%w: no shard with num <%d>", ErrShardNotFound, ix)
	}
}

// MustGet returns the shard by its number
func (s *Shards[T]) MustGet(num uint16) T {
	got, err := s.Get(num)
	if err != nil {
		panic(fmt.Errorf("must get failed: %w", err))
	}

	return got
}
