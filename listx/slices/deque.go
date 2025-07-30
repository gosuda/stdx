package slices

import (
	"errors"

	"github.com/gosuda/stdx/listx"
	"github.com/gosuda/stdx/option"
	"github.com/gosuda/stdx/result"
)

var _ listx.Deque[int] = (*SliceDeque[int])(nil)

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
func (d *SliceDeque[T]) RemoveFirst() result.Result[T, error] {
	if d.IsEmpty() {
		return result.Err[T, error](errors.New("deque is empty"))
	}
	return d.Remove(0)
}

// RemoveLast removes and returns the last element of the deque.
func (d *SliceDeque[T]) RemoveLast() result.Result[T, error] {
	if d.IsEmpty() {
		return result.Err[T, error](errors.New("deque is empty"))
	}
	return d.Remove(d.Size() - 1)
}

// PeekFirst returns the first element of the deque without removing it.
func (d *SliceDeque[T]) PeekFirst() option.Option[T] {
	if d.IsEmpty() {
		return option.None[T]()
	}
	return d.Get(0)
}

// PeekLast returns the last element of the deque without removing it.
func (d *SliceDeque[T]) PeekLast() option.Option[T] {
	if d.IsEmpty() {
		return option.None[T]()
	}
	return d.Get(d.Size() - 1)
}
