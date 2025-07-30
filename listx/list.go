package listx

import (
	"github.com/gosuda/stdx/option"
	"github.com/gosuda/stdx/result"
)

// List interface defines basic operations for ordered collections.
type List[T any] interface {
	// Add appends an element to the end of the list.
	Add(element T)

	// Insert inserts an element at the specified index.
	Insert(index int, element T) error

	// Get returns the element at the specified index.
	Get(index int) option.Option[T]

	// Set sets the element at the specified index to a new value.
	Set(index int, element T) error

	// Remove removes the element at the specified index.
	Remove(index int) result.Result[T, error]

	// RemoveElement removes the first matching element.
	RemoveElement(element T) bool

	// IndexOf returns the first index of the element, or None if not found.
	IndexOf(element T) option.Option[int]

	// LastIndexOf returns the last index of the element, or None if not found.
	LastIndexOf(element T) option.Option[int]

	// Contains checks if the element is contained in the list.
	Contains(element T) bool

	// Size returns the size of the list.
	Size() int

	// IsEmpty checks if the list is empty.
	IsEmpty() bool

	// Clear removes all elements from the list.
	Clear()

	// ToSlice returns all elements of the list as a slice.
	ToSlice() []T

	// ForEach executes a function for every element in the list.
	ForEach(fn func(element T))
}

// Deque interface defines operations for double-ended queue.
type Deque[T any] interface {
	List[T]

	// AddFirst adds an element to the front of the deque.
	AddFirst(element T)

	// AddLast adds an element to the back of the deque.
	AddLast(element T)

	// RemoveFirst removes and returns the first element of the deque.
	RemoveFirst() result.Result[T, error]

	// RemoveLast removes and returns the last element of the deque.
	RemoveLast() result.Result[T, error]

	// PeekFirst returns the first element of the deque without removing it.
	PeekFirst() option.Option[T]

	// PeekLast returns the last element of the deque without removing it.
	PeekLast() option.Option[T]
}

// Stack interface defines operations for stack (LIFO) data structure.
type Stack[T any] interface {
	// Push adds an element to the top of the stack.
	Push(element T)

	// Pop removes and returns the top element of the stack.
	Pop() result.Result[T, error]

	// Peek returns the top element of the stack without removing it.
	Peek() option.Option[T]

	// Size returns the size of the stack.
	Size() int

	// IsEmpty checks if the stack is empty.
	IsEmpty() bool

	// Clear removes all elements from the stack.
	Clear()

	// ToSlice returns all elements of the stack as a slice (from top to bottom).
	ToSlice() []T
}

// Queue interface defines operations for queue (FIFO) data structure.
type Queue[T any] interface {
	// Enqueue adds an element to the back of the queue.
	Enqueue(element T)

	// Dequeue removes and returns the front element of the queue.
	Dequeue() result.Result[T, error]

	// Peek returns the front element of the queue without removing it.
	Peek() option.Option[T]

	// Size returns the size of the queue.
	Size() int

	// IsEmpty checks if the queue is empty.
	IsEmpty() bool

	// Clear removes all elements from the queue.
	Clear()

	// ToSlice returns all elements of the queue as a slice (from front to back).
	ToSlice() []T
}
