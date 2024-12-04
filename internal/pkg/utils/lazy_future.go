package utils

import (
	"context"
	"errors"
	"sync"
)

// LazyFuture Future, которая начинает вычисляться во время вызова GetResult.
type LazyFuture[T any] struct {
	result T
	err    error
	once   *sync.Once
	fn     func(ctx context.Context) (T, error)
}

// NewLazyFuture ...
func NewLazyFuture[T any](fn func(ctx context.Context) (T, error)) *LazyFuture[T] {
	return &LazyFuture[T]{
		once: &sync.Once{},
		fn:   fn,
	}
}

// GetResult ...
func (f *LazyFuture[T]) GetResult(ctx context.Context) (T, error) {
	if f == nil {
		var r T
		return r, errors.New("future is not initialized")
	}

	f.once.Do(func() {
		f.result, f.err = f.fn(ctx)
	})

	return f.result, f.err
}
