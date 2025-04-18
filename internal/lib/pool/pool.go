package pool

import "context"

// Pool very simple pool implementation
type Pool[T any] struct {
	active chan T
}

// New returns pool with fixed member size
func New[T any](members []T) *Pool[T] {
	active := make(chan T, len(members))
	for i := range members {
		active <- members[i]
	}

	return &Pool[T]{
		active: active,
	}
}

// GetE изымет из пула элемент на время выполнения функции коллбека,
// затем вернет элемент обратно в пул
func (p *Pool[T]) GetE(ctx context.Context, cb func(T) error) error {
	var conn T
	select {
	case conn = <-p.active:
		break
	case <-ctx.Done():
		return ctx.Err()
	}

	defer func() {
		p.active <- conn
	}()

	return cb(conn)
}
