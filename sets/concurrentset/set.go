package concurrentset

import (
	"sync"

	"github.com/gosuda/stdx/sets"
)

var _ sets.Set[int] = (*ConcurrentSet[int])(nil)

type ConcurrentSet[T comparable] struct {
	elements sync.Map
}

func New[T comparable]() *ConcurrentSet[T] {
	return &ConcurrentSet[T]{}
}

// Add implements sets.Set.
func (c *ConcurrentSet[T]) Add(element T) bool {
	_, loaded := c.elements.LoadOrStore(element, struct{}{})
	return !loaded // true if element was not present before
}

// Clear implements sets.Set.
func (c *ConcurrentSet[T]) Clear() {
	c.elements.Clear()
}

// Contains implements sets.Set.
func (c *ConcurrentSet[T]) Contains(element T) bool {
	_, exists := c.elements.Load(element)
	return exists
}

// Difference implements sets.Set.
func (c *ConcurrentSet[T]) Difference(other sets.Set[T]) sets.Set[T] {
	result := New[T]()
	c.elements.Range(func(key, value any) bool {
		element := key.(T)
		if !other.Contains(element) {
			result.Add(element)
		}
		return true
	})
	return result
}

// ForEach implements sets.Set.
func (c *ConcurrentSet[T]) ForEach(fn func(element T)) {
	c.elements.Range(func(key, value any) bool {
		fn(key.(T))
		return true
	})
}

// Intersection implements sets.Set.
func (c *ConcurrentSet[T]) Intersection(other sets.Set[T]) sets.Set[T] {
	result := New[T]()
	c.elements.Range(func(key, value any) bool {
		element := key.(T)
		if other.Contains(element) {
			result.Add(element)
		}
		return true
	})
	return result
}

// IsEmpty implements sets.Set.
func (c *ConcurrentSet[T]) IsEmpty() bool {
	isEmpty := true
	c.elements.Range(func(key, value any) bool {
		isEmpty = false
		return false // stop iteration
	})
	return isEmpty
}

// IsSubsetOf implements sets.Set.
func (c *ConcurrentSet[T]) IsSubsetOf(other sets.Set[T]) bool {
	isSubset := true
	c.elements.Range(func(key, value any) bool {
		element := key.(T)
		if !other.Contains(element) {
			isSubset = false
			return false // stop iteration
		}
		return true
	})
	return isSubset
}

// IsSupersetOf implements sets.Set.
func (c *ConcurrentSet[T]) IsSupersetOf(other sets.Set[T]) bool {
	return other.IsSubsetOf(c)
}

// Remove implements sets.Set.
func (c *ConcurrentSet[T]) Remove(element T) bool {
	_, existed := c.elements.LoadAndDelete(element)
	return existed
}

// Size implements sets.Set.
func (c *ConcurrentSet[T]) Size() int {
	count := 0
	c.elements.Range(func(key, value any) bool {
		count++
		return true
	})
	return count
}

// ToSlice implements sets.Set.
func (c *ConcurrentSet[T]) ToSlice() []T {
	var result []T
	c.elements.Range(func(key, value any) bool {
		result = append(result, key.(T))
		return true
	})
	return result
}

// Union implements sets.Set.
func (c *ConcurrentSet[T]) Union(other sets.Set[T]) sets.Set[T] {
	result := New[T]()
	// Add all elements from current set
	c.elements.Range(func(key, value any) bool {
		result.Add(key.(T))
		return true
	})
	// Add all elements from other set
	other.ForEach(func(element T) {
		result.Add(element)
	})
	return result
}
