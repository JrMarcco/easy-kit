package xsync

import "sync"

// Pool is simple packing for sync.Pool.
type Pool[T any] struct {
	p sync.Pool
}

// Get gets a new object from the pool.
func (p *Pool[T]) Get() T {
	return p.p.Get().(T)
}

// Put puts an object back into the pool.
func (p *Pool[T]) Put(t T) {
	p.p.Put(t)
}

// NewPool creates a new pool with a factory function.
// factory must not return nil.
func NewPool[T any](factory func() T) *Pool[T] {
	return &Pool[T]{
		p: sync.Pool{
			New: func() any {
				return factory()
			},
		},
	}
}
