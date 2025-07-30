package linked

import (
	"errors"
	"reflect"

	"github.com/gosuda/stdx/listx"
	"github.com/gosuda/stdx/option"
	"github.com/gosuda/stdx/result"
)

var _ listx.List[int] = (*LinkedList[int])(nil)

// Node represents a node in the linked list
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList is a linked list implementation of the List interface
type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// New creates a new LinkedList
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add appends an element to the end of the list.
func (l *LinkedList[T]) Add(element T) {
	newNode := &Node[T]{Value: element}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.Next = newNode
		l.tail = newNode
	}
	l.size++
}

// Insert inserts an element at the specified index.
func (l *LinkedList[T]) Insert(index int, element T) error {
	if index < 0 || index > l.size {
		return errors.New("index out of bounds")
	}

	if index == l.size {
		l.Add(element)
		return nil
	}

	newNode := &Node[T]{Value: element}

	if index == 0 {
		newNode.Next = l.head
		l.head = newNode
		if l.tail == nil {
			l.tail = newNode
		}
	} else {
		prev := l.getNodeAt(index - 1)
		newNode.Next = prev.Next
		prev.Next = newNode
	}

	l.size++
	return nil
}

// Get returns the element at the specified index.
func (l *LinkedList[T]) Get(index int) option.Option[T] {
	if index < 0 || index >= l.size {
		return option.None[T]()
	}

	node := l.getNodeAt(index)
	return option.Some(node.Value)
}

// Set sets the element at the specified index to a new value.
func (l *LinkedList[T]) Set(index int, element T) error {
	if index < 0 || index >= l.size {
		return errors.New("index out of bounds")
	}

	node := l.getNodeAt(index)
	node.Value = element
	return nil
}

// Remove removes the element at the specified index.
func (l *LinkedList[T]) Remove(index int) result.Result[T, error] {
	if index < 0 || index >= l.size {
		return result.Err[T, error](errors.New("index out of bounds"))
	}

	var removedValue T

	if index == 0 {
		removedValue = l.head.Value
		l.head = l.head.Next
		if l.head == nil {
			l.tail = nil
		}
	} else {
		prev := l.getNodeAt(index - 1)
		removedValue = prev.Next.Value
		prev.Next = prev.Next.Next
		if prev.Next == nil {
			l.tail = prev
		}
	}

	l.size--
	return result.Ok[T, error](removedValue)
}

// RemoveElement removes the first matching element.
func (l *LinkedList[T]) RemoveElement(element T) bool {
	if l.head == nil {
		return false
	}

	if reflect.DeepEqual(l.head.Value, element) {
		l.head = l.head.Next
		if l.head == nil {
			l.tail = nil
		}
		l.size--
		return true
	}

	current := l.head
	for current.Next != nil {
		if reflect.DeepEqual(current.Next.Value, element) {
			current.Next = current.Next.Next
			if current.Next == nil {
				l.tail = current
			}
			l.size--
			return true
		}
		current = current.Next
	}

	return false
}

// IndexOf returns the first index of the element.
func (l *LinkedList[T]) IndexOf(element T) option.Option[int] {
	current := l.head
	for i := 0; current != nil; i++ {
		if reflect.DeepEqual(current.Value, element) {
			return option.Some(i)
		}
		current = current.Next
	}
	return option.None[int]()
}

// LastIndexOf returns the last index of the element, or None if not found.
func (l *LinkedList[T]) LastIndexOf(element T) option.Option[int] {
	lastIndex := -1
	current := l.head
	for i := 0; current != nil; i++ {
		if reflect.DeepEqual(current.Value, element) {
			lastIndex = i
		}
		current = current.Next
	}
	if lastIndex == -1 {
		return option.None[int]()
	}
	return option.Some(lastIndex)
}

// Contains checks if the element is contained in the list.
func (l *LinkedList[T]) Contains(element T) bool {
	return l.IndexOf(element).IsSome()
}

// Size returns the size of the list.
func (l *LinkedList[T]) Size() int {
	return l.size
}

// IsEmpty checks if the list is empty.
func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Clear removes all elements from the list.
func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// ToSlice returns all elements of the list as a slice.
func (l *LinkedList[T]) ToSlice() []T {
	result := make([]T, l.size)
	current := l.head
	for i := 0; current != nil; i++ {
		result[i] = current.Value
		current = current.Next
	}
	return result
}

// ForEach executes a function for every element in the list.
func (l *LinkedList[T]) ForEach(fn func(element T)) {
	current := l.head
	for current != nil {
		fn(current.Value)
		current = current.Next
	}
}

// getNodeAt returns the node at the specified index (internal helper method)
func (l *LinkedList[T]) getNodeAt(index int) *Node[T] {
	current := l.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current
}
