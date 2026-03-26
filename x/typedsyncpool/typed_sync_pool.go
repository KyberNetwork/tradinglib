package typedsyncpool

import "sync"

type Pool[T any] struct {
	p *sync.Pool
}

func New[T any](fn func() T) *Pool[T] {
	return &Pool[T]{
		p: &sync.Pool{
			New: func() any {
				return fn()
			},
		},
	}
}

func (p *Pool[T]) Put(x T) {
	p.p.Put(x)
}

func (p *Pool[T]) Get() T {
	v, _ := p.p.Get().(T)
	return v
}
