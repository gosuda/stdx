package linked

import (
	"errors"

	"github.com/gosuda/stdx/lists"
)

var _ lists.Stack[int] = (*LinkedStack[int])(nil)

// LinkedStack is a linked list implementation of the Stack interface
type LinkedStack[T any] struct {
	list *LinkedList[T]
}

// NewStack creates a new LinkedStack
func NewStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{
		list: New[T](),
	}
}

// Push adds an element to the top of the stack.
func (s *LinkedStack[T]) Push(element T) {
	s.list.Insert(0, element) // Insert at beginning (top of stack)
}

// Pop removes and returns the top element of the stack.
func (s *LinkedStack[T]) Pop() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack is empty")
	}
	return s.list.Remove(0) // Remove from beginning (top of stack)
}

// Peek returns the top element of the stack without removing it.
func (s *LinkedStack[T]) Peek() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack is empty")
	}
	return s.list.Get(0) // Get from beginning (top of stack)
}

// Size returns the size of the stack.
func (s *LinkedStack[T]) Size() int {
	return s.list.Size()
}

// IsEmpty checks if the stack is empty.
func (s *LinkedStack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

// Clear removes all elements from the stack.
func (s *LinkedStack[T]) Clear() {
	s.list.Clear()
}

// ToSlice returns all elements of the stack as a slice (from top to bottom).
func (s *LinkedStack[T]) ToSlice() []T {
	return s.list.ToSlice() // Already in top-to-bottom order
}
