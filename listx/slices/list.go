package slices

import (
	"errors"
	"reflect"

	"github.com/gosuda/stdx/listx"
	"github.com/gosuda/stdx/option"
	"github.com/gosuda/stdx/result"
)

var _ listx.List[int] = (*SliceList[int])(nil)

// SliceList is a slice-based implementation of the List interface
type SliceList[T any] struct {
	elements []T
}

// New creates a new SliceList
func New[T any]() *SliceList[T] {
	return &SliceList[T]{
		elements: make([]T, 0),
	}
}

// Add appends an element to the end of the list.
func (s *SliceList[T]) Add(element T) {
	s.elements = append(s.elements, element)
}

// Insert inserts an element at the specified index.
func (s *SliceList[T]) Insert(index int, element T) error {
	if index < 0 || index > len(s.elements) {
		return errors.New("index out of bounds")
	}

	if index == len(s.elements) {
		s.elements = append(s.elements, element)
		return nil
	}

	// Expand slice by one element
	s.elements = append(s.elements, *new(T))

	// Shift elements to the right
	copy(s.elements[index+1:], s.elements[index:])

	// Insert the new element
	s.elements[index] = element

	return nil
}

// Get returns the element at the specified index.
func (s *SliceList[T]) Get(index int) option.Option[T] {
	if index < 0 || index >= len(s.elements) {
		return option.None[T]()
	}
	return option.Some(s.elements[index])
}

// Set sets the element at the specified index to a new value.
func (s *SliceList[T]) Set(index int, element T) error {
	if index < 0 || index >= len(s.elements) {
		return errors.New("index out of bounds")
	}
	s.elements[index] = element
	return nil
}

// Remove removes the element at the specified index.
func (s *SliceList[T]) Remove(index int) result.Result[T, error] {
	if index < 0 || index >= len(s.elements) {
		return result.Err[T, error](errors.New("index out of bounds"))
	}

	removedElement := s.elements[index]

	// Shift elements to the left
	copy(s.elements[index:], s.elements[index+1:])

	// Shrink slice
	s.elements = s.elements[:len(s.elements)-1]

	return result.Ok[T, error](removedElement)
}

// RemoveElement removes the first matching element.
func (s *SliceList[T]) RemoveElement(element T) bool {
	indexOpt := s.IndexOf(element)
	if indexOpt.IsNone() {
		return false
	}

	result := s.Remove(indexOpt.Unwrap())
	return result.IsOk()
}

// IndexOf returns the first index of the element.
func (s *SliceList[T]) IndexOf(element T) option.Option[int] {
	for i, elem := range s.elements {
		if reflect.DeepEqual(elem, element) {
			return option.Some(i)
		}
	}
	return option.None[int]()
}

// LastIndexOf returns the last index of the element, or None if not found.
func (s *SliceList[T]) LastIndexOf(element T) option.Option[int] {
	lastIndex := -1
	for i, elem := range s.elements {
		if reflect.DeepEqual(elem, element) {
			lastIndex = i
		}
	}
	if lastIndex == -1 {
		return option.None[int]()
	}
	return option.Some(lastIndex)
}

// Contains checks if the element is contained in the list.
func (s *SliceList[T]) Contains(element T) bool {
	return s.IndexOf(element).IsSome()
}

// Size returns the size of the list.
func (s *SliceList[T]) Size() int {
	return len(s.elements)
}

// IsEmpty checks if the list is empty.
func (s *SliceList[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Clear removes all elements from the list.
func (s *SliceList[T]) Clear() {
	s.elements = s.elements[:0] // Keep underlying array capacity
}

// ToSlice returns all elements of the list as a slice.
func (s *SliceList[T]) ToSlice() []T {
	// Return a copy to prevent external modification
	result := make([]T, len(s.elements))
	copy(result, s.elements)
	return result
}

// ForEach executes a function for every element in the list.
func (s *SliceList[T]) ForEach(fn func(element T)) {
	for _, element := range s.elements {
		fn(element)
	}
}
