package hash

import (
	"errors"
	"reflect"

	"github.com/gosuda/stdx/listx"
	"github.com/gosuda/stdx/option"
	"github.com/gosuda/stdx/result"
)

var _ listx.List[int] = (*HashList[int])(nil)

// HashList is a hash-based implementation of the List interface using map[int]T
type HashList[T any] struct {
	elements map[int]T
	size     int
}

// New creates a new HashList
func New[T any]() *HashList[T] {
	return &HashList[T]{
		elements: make(map[int]T),
		size:     0,
	}
}

// Add appends an element to the end of the list.
func (h *HashList[T]) Add(element T) {
	h.elements[h.size] = element
	h.size++
}

// Insert inserts an element at the specified index.
func (h *HashList[T]) Insert(index int, element T) error {
	if index < 0 || index > h.size {
		return errors.New("index out of bounds")
	}

	if index == h.size {
		h.Add(element)
		return nil
	}

	// Shift all elements from index onwards to the right
	for i := h.size; i > index; i-- {
		h.elements[i] = h.elements[i-1]
	}

	h.elements[index] = element
	h.size++
	return nil
}

// Get returns the element at the specified index.
func (h *HashList[T]) Get(index int) option.Option[T] {
	if index < 0 || index >= h.size {
		return option.None[T]()
	}

	element, exists := h.elements[index]
	if !exists {
		return option.None[T]()
	}

	return option.Some(element)
}

// Set sets the element at the specified index to a new value.
func (h *HashList[T]) Set(index int, element T) error {
	if index < 0 || index >= h.size {
		return errors.New("index out of bounds")
	}

	h.elements[index] = element
	return nil
}

// Remove removes the element at the specified index.
func (h *HashList[T]) Remove(index int) result.Result[T, error] {
	if index < 0 || index >= h.size {
		return result.Err[T, error](errors.New("index out of bounds"))
	}

	removedElement := h.elements[index]

	// Shift all elements from index+1 onwards to the left
	for i := index; i < h.size-1; i++ {
		h.elements[i] = h.elements[i+1]
	}

	// Remove the last element
	delete(h.elements, h.size-1)
	h.size--

	return result.Ok[T, error](removedElement)
}

// RemoveElement removes the first matching element.
func (h *HashList[T]) RemoveElement(element T) bool {
	indexOpt := h.IndexOf(element)
	if indexOpt.IsNone() {
		return false
	}

	result := h.Remove(indexOpt.Unwrap())
	return result.IsOk()
}

// IndexOf returns the first index of the element, or None if not found.
func (h *HashList[T]) IndexOf(element T) option.Option[int] {
	for i := 0; i < h.size; i++ {
		if elem, exists := h.elements[i]; exists && reflect.DeepEqual(elem, element) {
			return option.Some(i)
		}
	}
	return option.None[int]()
}

// LastIndexOf returns the last index of the element, or None if not found.
func (h *HashList[T]) LastIndexOf(element T) option.Option[int] {
	lastIndex := -1
	for i := 0; i < h.size; i++ {
		if elem, exists := h.elements[i]; exists && reflect.DeepEqual(elem, element) {
			lastIndex = i
		}
	}
	if lastIndex == -1 {
		return option.None[int]()
	}
	return option.Some(lastIndex)
}

// Contains checks if the element is contained in the list.
func (h *HashList[T]) Contains(element T) bool {
	return h.IndexOf(element).IsSome()
}

// Size returns the size of the list.
func (h *HashList[T]) Size() int {
	return h.size
}

// IsEmpty checks if the list is empty.
func (h *HashList[T]) IsEmpty() bool {
	return h.size == 0
}

// Clear removes all elements from the list.
func (h *HashList[T]) Clear() {
	h.elements = make(map[int]T)
	h.size = 0
}

// ToSlice returns all elements of the list as a slice.
func (h *HashList[T]) ToSlice() []T {
	result := make([]T, h.size)
	for i := 0; i < h.size; i++ {
		if elem, exists := h.elements[i]; exists {
			result[i] = elem
		}
	}
	return result
}

// ForEach executes a function for every element in the list.
func (h *HashList[T]) ForEach(fn func(element T)) {
	for i := 0; i < h.size; i++ {
		if elem, exists := h.elements[i]; exists {
			fn(elem)
		}
	}
}
