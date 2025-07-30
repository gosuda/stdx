package hash

import (
	"errors"

	"github.com/gosuda/stdx/lists"
)

var _ lists.Deque[int] = (*HashDeque[int])(nil)

// HashDeque is a hash-based implementation of the Deque interface
type HashDeque[T any] struct {
	*HashList[T]
}

// NewDeque creates a new HashDeque
func NewDeque[T any]() *HashDeque[T] {
	return &HashDeque[T]{
		HashList: New[T](),
	}
}

// AddFirst adds an element to the front of the deque.
func (d *HashDeque[T]) AddFirst(element T) {
	d.Insert(0, element)
}

// AddLast adds an element to the back of the deque.
func (d *HashDeque[T]) AddLast(element T) {
	d.Add(element)
}

// RemoveFirst removes and returns the first element of the deque.
func (d *HashDeque[T]) RemoveFirst() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.Remove(0)
}

// RemoveLast removes and returns the last element of the deque.
func (d *HashDeque[T]) RemoveLast() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.Remove(d.Size() - 1)
}

// PeekFirst returns the first element of the deque without removing it.
func (d *HashDeque[T]) PeekFirst() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.Get(0)
}

// PeekLast returns the last element of the deque without removing it.
func (d *HashDeque[T]) PeekLast() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.Get(d.Size() - 1)
}
