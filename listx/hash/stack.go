package hash

import (
	"errors"

	"github.com/gosuda/stdx/listx"
	"github.com/gosuda/stdx/option"
	"github.com/gosuda/stdx/result"
)

var _ listx.Stack[int] = (*HashStack[int])(nil)

// HashStack is a hash-based implementation of the Stack interface
type HashStack[T any] struct {
	list *HashList[T]
}

// NewStack creates a new HashStack
func NewStack[T any]() *HashStack[T] {
	return &HashStack[T]{
		list: New[T](),
	}
}

// Push adds an element to the top of the stack.
func (s *HashStack[T]) Push(element T) {
	s.list.Add(element) // Add to end (top of stack for efficiency)
}

// Pop removes and returns the top element of the stack.
func (s *HashStack[T]) Pop() result.Result[T, error] {
	if s.IsEmpty() {
		return result.Err[T, error](errors.New("stack is empty"))
	}
	return s.list.Remove(s.list.Size() - 1) // Remove from end (top of stack)
}

// Peek returns the top element of the stack without removing it.
func (s *HashStack[T]) Peek() option.Option[T] {
	if s.IsEmpty() {
		return option.None[T]()
	}
	return s.list.Get(s.list.Size() - 1) // Get from end (top of stack)
}

// Size returns the size of the stack.
func (s *HashStack[T]) Size() int {
	return s.list.Size()
}

// IsEmpty checks if the stack is empty.
func (s *HashStack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

// Clear removes all elements from the stack.
func (s *HashStack[T]) Clear() {
	s.list.Clear()
}

// ToSlice returns all elements of the stack as a slice (from top to bottom).
func (s *HashStack[T]) ToSlice() []T {
	slice := s.list.ToSlice()
	// Reverse the slice to show top-to-bottom order
	result := make([]T, len(slice))
	for i, elem := range slice {
		result[len(slice)-1-i] = elem
	}
	return result
}
