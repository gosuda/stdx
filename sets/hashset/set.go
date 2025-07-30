package hashset

import "github.com/gosuda/stdx/sets"

var _ sets.Set[int] = (*HashSet[int])(nil)

type HashSet[T comparable] struct {
	elements map[T]struct{}
}

func New[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		elements: make(map[T]struct{}),
	}
}

// Add implements sets.Set.
func (h *HashSet[T]) Add(element T) bool {
	if _, exists := h.elements[element]; exists {
		return false
	}
	h.elements[element] = struct{}{}
	return true
}

// Clear implements sets.Set.
func (h *HashSet[T]) Clear() {
	h.elements = make(map[T]struct{})
}

// Contains implements sets.Set.
func (h *HashSet[T]) Contains(element T) bool {
	_, exists := h.elements[element]
	return exists
}

// Difference implements sets.Set.
func (h *HashSet[T]) Difference(other sets.Set[T]) sets.Set[T] {
	result := New[T]()
	for element := range h.elements {
		if !other.Contains(element) {
			result.Add(element)
		}
	}
	return result
}

// Intersection implements sets.Set.
func (h *HashSet[T]) Intersection(other sets.Set[T]) sets.Set[T] {
	result := New[T]()
	for element := range h.elements {
		if other.Contains(element) {
			result.Add(element)
		}
	}
	return result
}

// IsEmpty implements sets.Set.
func (h *HashSet[T]) IsEmpty() bool {
	return len(h.elements) == 0
}

// IsSubsetOf implements sets.Set.
func (h *HashSet[T]) IsSubsetOf(other sets.Set[T]) bool {
	for element := range h.elements {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

// IsSupersetOf implements sets.Set.
func (h *HashSet[T]) IsSupersetOf(other sets.Set[T]) bool {
	return other.IsSubsetOf(h)
}

// Remove implements sets.Set.
func (h *HashSet[T]) Remove(element T) bool {
	if _, exists := h.elements[element]; exists {
		delete(h.elements, element)
		return true
	}
	return false
}

// Size implements sets.Set.
func (h *HashSet[T]) Size() int {
	return len(h.elements)
}

// ToSlice implements sets.Set.
func (h *HashSet[T]) ToSlice() []T {
	result := make([]T, 0, len(h.elements))
	for element := range h.elements {
		result = append(result, element)
	}
	return result
}

// ForEach implements sets.Set.
func (h *HashSet[T]) ForEach(fn func(element T)) {
	for element := range h.elements {
		fn(element)
	}
}

// Union implements sets.Set.
func (h *HashSet[T]) Union(other sets.Set[T]) sets.Set[T] {
	result := New[T]()
	// Add all elements from current set
	for element := range h.elements {
		result.Add(element)
	}
	// Add all elements from other set
	other.ForEach(func(element T) {
		result.Add(element)
	})
	return result
}
