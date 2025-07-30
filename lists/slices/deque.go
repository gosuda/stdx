package slices

import (
	"errors"

	"github.com/gosuda/stdx/lists"
)

var _ lists.Deque[int] = (*SliceDeque[int])(nil)

// SliceDeque is a slice-based implementation of the Deque interface
type SliceDeque[T any] struct {
	*SliceList[T]
}

// NewDeque creates a new SliceDeque
func NewDeque[T any]() *SliceDeque[T] {
	return &SliceDeque[T]{
		SliceList: New[T](),
	}
}

// AddFirst adds an element to the front of the deque.
func (d *SliceDeque[T]) AddFirst(element T) {
	d.Insert(0, element)
}

// AddLast adds an element to the back of the deque.
func (d *SliceDeque[T]) AddLast(element T) {
	d.Add(element)
}

// RemoveFirst removes and returns the first element of the deque.
func (d *SliceDeque[T]) RemoveFirst() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.Remove(0)
}

// RemoveLast removes and returns the last element of the deque.
func (d *SliceDeque[T]) RemoveLast() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.Remove(d.Size() - 1)
}

// PeekFirst returns the first element of the deque without removing it.
func (d *SliceDeque[T]) PeekFirst() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.Get(0)
}

// PeekLast returns the last element of the deque without removing it.
func (d *SliceDeque[T]) PeekLast() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.Get(d.Size() - 1)
}
