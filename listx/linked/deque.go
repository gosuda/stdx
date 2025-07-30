package linked

import (
	"errors"

	"github.com/gosuda/stdx/listx"
)

var _ listx.Deque[int] = (*LinkedDeque[int])(nil)

// LinkedDeque is a linked list implementation of the Deque interface
type LinkedDeque[T any] struct {
	*LinkedList[T]
}

// NewDeque creates a new LinkedDeque
func NewDeque[T any]() *LinkedDeque[T] {
	return &LinkedDeque[T]{
		LinkedList: New[T](),
	}
}

// AddFirst adds an element to the front of the deque.
func (d *LinkedDeque[T]) AddFirst(element T) {
	d.Insert(0, element)
}

// AddLast adds an element to the back of the deque.
func (d *LinkedDeque[T]) AddLast(element T) {
	d.Add(element)
}

// RemoveFirst removes and returns the first element of the deque.
func (d *LinkedDeque[T]) RemoveFirst() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.Remove(0)
}

// RemoveLast removes and returns the last element of the deque.
func (d *LinkedDeque[T]) RemoveLast() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.Remove(d.Size() - 1)
}

// PeekFirst returns the first element of the deque without removing it.
func (d *LinkedDeque[T]) PeekFirst() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.Get(0)
}

// PeekLast returns the last element of the deque without removing it.
func (d *LinkedDeque[T]) PeekLast() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.Get(d.Size() - 1)
}
