package slices

import (
	"errors"

	"github.com/gosuda/stdx/listx"
)

var _ listx.Queue[int] = (*SliceQueue[int])(nil)

// SliceQueue is a slice-based implementation of the Queue interface
type SliceQueue[T any] struct {
	list *SliceList[T]
}

// NewQueue creates a new SliceQueue
func NewQueue[T any]() *SliceQueue[T] {
	return &SliceQueue[T]{
		list: New[T](),
	}
}

// Enqueue adds an element to the back of the queue.
func (q *SliceQueue[T]) Enqueue(element T) {
	q.list.Add(element) // Add to end (back of queue)
}

// Dequeue removes and returns the front element of the queue.
func (q *SliceQueue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errors.New("queue is empty")
	}
	return q.list.Remove(0) // Remove from beginning (front of queue)
}

// Peek returns the front element of the queue without removing it.
func (q *SliceQueue[T]) Peek() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errors.New("queue is empty")
	}
	return q.list.Get(0) // Get from beginning (front of queue)
}

// Size returns the size of the queue.
func (q *SliceQueue[T]) Size() int {
	return q.list.Size()
}

// IsEmpty checks if the queue is empty.
func (q *SliceQueue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

// Clear removes all elements from the queue.
func (q *SliceQueue[T]) Clear() {
	q.list.Clear()
}

// ToSlice returns all elements of the queue as a slice (from front to back).
func (q *SliceQueue[T]) ToSlice() []T {
	return q.list.ToSlice() // Already in front-to-back order
}
