package slices

import (
	"errors"

	"github.com/gosuda/stdx/listx"
)

var _ listx.Stack[int] = (*SliceStack[int])(nil)

// SliceStack is a slice-based implementation of the Stack interface
type SliceStack[T any] struct {
	list *SliceList[T]
}

// NewStack creates a new SliceStack
func NewStack[T any]() *SliceStack[T] {
	return &SliceStack[T]{
		list: New[T](),
	}
}

// Push adds an element to the top of the stack.
func (s *SliceStack[T]) Push(element T) {
	s.list.Add(element) // Add to end (top of stack for efficiency)
}

// Pop removes and returns the top element of the stack.
func (s *SliceStack[T]) Pop() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack is empty")
	}
	return s.list.Remove(s.list.Size() - 1) // Remove from end (top of stack)
}

// Peek returns the top element of the stack without removing it.
func (s *SliceStack[T]) Peek() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack is empty")
	}
	return s.list.Get(s.list.Size() - 1) // Get from end (top of stack)
}

// Size returns the size of the stack.
func (s *SliceStack[T]) Size() int {
	return s.list.Size()
}

// IsEmpty checks if the stack is empty.
func (s *SliceStack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

// Clear removes all elements from the stack.
func (s *SliceStack[T]) Clear() {
	s.list.Clear()
}

// ToSlice returns all elements of the stack as a slice (from top to bottom).
func (s *SliceStack[T]) ToSlice() []T {
	slice := s.list.ToSlice()
	// Reverse the slice to show top-to-bottom order
	result := make([]T, len(slice))
	for i, elem := range slice {
		result[len(slice)-1-i] = elem
	}
	return result
}
