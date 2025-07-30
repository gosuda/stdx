package hash

import (
	"errors"
	"reflect"

	"github.com/gosuda/stdx/lists"
)

var _ lists.List[int] = (*HashList[int])(nil)

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
func (h *HashList[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= h.size {
		return zero, errors.New("index out of bounds")
	}

	element, exists := h.elements[index]
	if !exists {
		return zero, errors.New("element not found")
	}

	return element, nil
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
func (h *HashList[T]) Remove(index int) (T, error) {
	var zero T
	if index < 0 || index >= h.size {
		return zero, errors.New("index out of bounds")
	}

	removedElement := h.elements[index]

	// Shift all elements from index+1 onwards to the left
	for i := index; i < h.size-1; i++ {
		h.elements[i] = h.elements[i+1]
	}

	// Remove the last element
	delete(h.elements, h.size-1)
	h.size--

	return removedElement, nil
}

// RemoveElement removes the first matching element.
func (h *HashList[T]) RemoveElement(element T) bool {
	index := h.IndexOf(element)
	if index == -1 {
		return false
	}

	_, err := h.Remove(index)
	return err == nil
}

// IndexOf returns the first index of the element.
func (h *HashList[T]) IndexOf(element T) int {
	for i := 0; i < h.size; i++ {
		if elem, exists := h.elements[i]; exists && reflect.DeepEqual(elem, element) {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the last index of the element.
func (h *HashList[T]) LastIndexOf(element T) int {
	lastIndex := -1
	for i := 0; i < h.size; i++ {
		if elem, exists := h.elements[i]; exists && reflect.DeepEqual(elem, element) {
			lastIndex = i
		}
	}
	return lastIndex
}

// Contains checks if the element is contained in the list.
func (h *HashList[T]) Contains(element T) bool {
	return h.IndexOf(element) != -1
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
