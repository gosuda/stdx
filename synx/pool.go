package synx

import (
	"sync"
)

// Pool is a generic wrapper around sync.Pool with type safety
type Pool[T any] struct {
	pool sync.Pool
}

// NewPool creates a new generic Pool with a factory function
func NewPool[T any](factory func() T) *Pool[T] {
	return &Pool[T]{
		pool: sync.Pool{
			New: func() interface{} {
				return factory()
			},
		},
	}
}

// NewPoolWithNew is an alias for NewPool for clarity
func NewPoolWithNew[T any](newFunc func() T) *Pool[T] {
	return NewPool(newFunc)
}

// Get retrieves an object from the pool, with type safety
func (p *Pool[T]) Get() T {
	return p.pool.Get().(T)
}

// Put adds an object back to the pool
func (p *Pool[T]) Put(obj T) {
	p.pool.Put(obj)
}

// GetOrCreate gets an object from the pool or creates one if pool is empty
func (p *Pool[T]) GetOrCreate(factory func() T) T {
	if obj := p.pool.Get(); obj != nil {
		return obj.(T)
	}
	return factory()
}

// TryGet attempts to get an object from the pool without blocking
// Returns the object and true if successful, zero value and false otherwise
func (p *Pool[T]) TryGet() (T, bool) {
	obj := p.pool.Get()
	if obj == nil {
		var zero T
		return zero, false
	}
	return obj.(T), true
}

// Reset clears the pool by creating a new underlying sync.Pool
func (p *Pool[T]) Reset(factory func() T) {
	p.pool = sync.Pool{
		New: func() interface{} {
			if factory != nil {
				return factory()
			}
			return nil
		},
	}
}

// StringPool is a specialized pool for strings
type StringPool struct {
	*Pool[string]
}

// NewStringPool creates a new string pool
func NewStringPool() *StringPool {
	return &StringPool{
		Pool: NewPool(func() string { return "" }),
	}
}

// ByteSlicePool is a specialized pool for byte slices
type ByteSlicePool struct {
	*Pool[[]byte]
}

// NewByteSlicePool creates a new byte slice pool with initial capacity
func NewByteSlicePool(initialCap int) *ByteSlicePool {
	return &ByteSlicePool{
		Pool: NewPool(func() []byte { return make([]byte, 0, initialCap) }),
	}
}

// GetWithCap gets a byte slice and ensures it has at least the specified capacity
func (p *ByteSlicePool) GetWithCap(minCap int) []byte {
	buf := p.Get()
	if cap(buf) < minCap {
		return make([]byte, 0, minCap)
	}
	return buf[:0] // Reset length but keep capacity
}

// PutReset resets the byte slice and puts it back to the pool
func (p *ByteSlicePool) PutReset(buf []byte) {
	p.Put(buf[:0]) // Reset length before putting back
}
