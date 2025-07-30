package hash

import (
	"errors"

	"github.com/gosuda/stdx/listx"
)

var _ listx.Queue[int] = (*HashQueue[int])(nil)

// HashQueue is a hash-based implementation of the Queue interface
type HashQueue[T any] struct {
	list *HashList[T]
}

// NewQueue creates a new HashQueue
func NewQueue[T any]() *HashQueue[T] {
	return &HashQueue[T]{
		list: New[T](),
	}
}

// Enqueue adds an element to the back of the queue.
func (q *HashQueue[T]) Enqueue(element T) {
	q.list.Add(element) // Add to end (back of queue)
}

// Dequeue removes and returns the front element of the queue.
func (q *HashQueue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errors.New("queue is empty")
	}
	return q.list.Remove(0) // Remove from beginning (front of queue)
}

// Peek returns the front element of the queue without removing it.
func (q *HashQueue[T]) Peek() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errors.New("queue is empty")
	}
	return q.list.Get(0) // Get from beginning (front of queue)
}

// Size returns the size of the queue.
func (q *HashQueue[T]) Size() int {
	return q.list.Size()
}

// IsEmpty checks if the queue is empty.
func (q *HashQueue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

// Clear removes all elements from the queue.
func (q *HashQueue[T]) Clear() {
	q.list.Clear()
}

// ToSlice returns all elements of the queue as a slice (from front to back).
func (q *HashQueue[T]) ToSlice() []T {
	return q.list.ToSlice() // Already in front-to-back order
}
