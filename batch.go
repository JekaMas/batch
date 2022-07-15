package batch

import (
	"sync"
)

type Batch[T any] struct {
	Values []T
	m      sync.RWMutex
}

func NewBatch[T any](values ...T) *Batch[T] {
	return &Batch[T]{values, sync.RWMutex{}}
}

func (b *Batch[T]) Apply(fns ...func(params ...T)) {
	var wg sync.WaitGroup

	wg.Add(len(fns))

	for _, fn := range fns {
		fn := fn

		go func() {
			defer wg.Done()

			fn(b.Values...)
		}()
	}

	wg.Wait()
}

func (b *Batch[T]) NewWrite(fn func(params ...T)) func(params ...T) {
	return func(params ...T) {
		b.m.Lock()
		defer b.m.Unlock()

		fn(params...)
	}
}

func (b *Batch[T]) NewRead(fn func(params ...T)) func(params ...T) {
	return func(params ...T) {
		b.m.RLock()
		defer b.m.RUnlock()

		fn(params...)
	}
}
