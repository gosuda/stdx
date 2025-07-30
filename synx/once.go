package synx

import (
	"sync"
)

// Once is a generic wrapper around sync.Once that can return a value
type Once[T any] struct {
	once   sync.Once
	value  T
	err    error
	hasErr bool
}

// Do executes the function exactly once and returns its result
// Subsequent calls return the cached result
func (o *Once[T]) Do(f func() T) T {
	o.once.Do(func() {
		o.value = f()
	})
	return o.value
}

// DoWithError executes the function exactly once and caches both value and error
// Subsequent calls return the cached result and error
func (o *Once[T]) DoWithError(f func() (T, error)) (T, error) {
	o.once.Do(func() {
		o.value, o.err = f()
		o.hasErr = true
	})
	return o.value, o.err
}

// Done reports whether Do has been called
func (o *Once[T]) Done() bool {
	return o.isDone()
}

// isDone is a helper to check if once.Do has been called
func (o *Once[T]) isDone() bool {
	// Use a dummy function to check if Do has been called
	called := false
	o.once.Do(func() {
		called = true
	})
	return !called
}

// Value returns the cached value if Do has been called, zero value otherwise
func (o *Once[T]) Value() T {
	return o.value
}

// Error returns the cached error if DoWithError has been called
func (o *Once[T]) Error() error {
	if o.hasErr {
		return o.err
	}
	return nil
}

// Reset creates a new Once instance (since sync.Once cannot be reset)
func (o *Once[T]) Reset() {
	var zero T
	o.once = sync.Once{}
	o.value = zero
	o.err = nil
	o.hasErr = false
}

// OnceValue creates a function that calls f only once and caches the result
// This is similar to sync.OnceValue but with better type safety
func OnceValue[T any](f func() T) func() T {
	var once sync.Once
	var value T

	return func() T {
		once.Do(func() {
			value = f()
		})
		return value
	}
}

// OnceValues creates a function that calls f only once and caches both return values
// This is similar to sync.OnceValues but with better type safety
func OnceValues[T, U any](f func() (T, U)) func() (T, U) {
	var once sync.Once
	var value1 T
	var value2 U

	return func() (T, U) {
		once.Do(func() {
			value1, value2 = f()
		})
		return value1, value2
	}
}

// OnceFunc creates a function that calls f only once
// This is similar to sync.OnceFunc but can work with any function signature
func OnceFunc(f func()) func() {
	var once sync.Once
	return func() {
		once.Do(f)
	}
}

// OnceFuncWithArg creates a function that calls f only once with the first argument provided
func OnceFuncWithArg[T any](f func(T)) func(T) {
	var once sync.Once
	var firstArg T
	var called bool

	return func(arg T) {
		if !called {
			firstArg = arg
			called = true
		}
		once.Do(func() {
			f(firstArg)
		})
	}
}

// OnceResult wraps a function to be called only once and caches the result with error
func OnceResult[T any](f func() (T, error)) func() (T, error) {
	var once sync.Once
	var value T
	var err error

	return func() (T, error) {
		once.Do(func() {
			value, err = f()
		})
		return value, err
	}
}

// LazyValue provides lazy initialization of a value
type LazyValue[T any] struct {
	once        sync.Once
	value       T
	init        func() T
	initialized bool
	mu          sync.RWMutex
}

// NewLazyValue creates a new lazy value with an initializer function
func NewLazyValue[T any](init func() T) *LazyValue[T] {
	return &LazyValue[T]{
		init: init,
	}
}

// Get returns the value, initializing it if necessary
func (l *LazyValue[T]) Get() T {
	l.once.Do(func() {
		if l.init != nil {
			l.value = l.init()
		}
		l.mu.Lock()
		l.initialized = true
		l.mu.Unlock()
	})
	return l.value
}

// IsInitialized checks if the value has been initialized
func (l *LazyValue[T]) IsInitialized() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.initialized
}

// Reset resets the lazy value (creates a new instance)
func (l *LazyValue[T]) Reset(init func() T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var zero T
	l.once = sync.Once{}
	l.value = zero
	l.init = init
	l.initialized = false
}
