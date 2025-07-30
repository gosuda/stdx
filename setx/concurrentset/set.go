package concurrentset

import (
	"errors"
	"sync"

	"github.com/gosuda/stdx/option"
	"github.com/gosuda/stdx/result"
	"github.com/gosuda/stdx/setx"
)

var _ setx.Set[int] = (*ConcurrentSet[int])(nil)

type ConcurrentSet[T comparable] struct {
	elements sync.Map
}

func New[T comparable]() *ConcurrentSet[T] {
	return &ConcurrentSet[T]{}
}

// Add implements setx.Set.
func (c *ConcurrentSet[T]) Add(element T) bool {
	_, loaded := c.elements.LoadOrStore(element, struct{}{})
	return !loaded // true if element was not present before
}

// Clear implements setx.Set.
func (c *ConcurrentSet[T]) Clear() {
	c.elements.Clear()
}

// Contains implements setx.Set.
func (c *ConcurrentSet[T]) Contains(element T) bool {
	_, exists := c.elements.Load(element)
	return exists
}

// Difference implements setx.Set.
func (c *ConcurrentSet[T]) Difference(other setx.Set[T]) setx.Set[T] {
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

// ForEach implements setx.Set.
func (c *ConcurrentSet[T]) ForEach(fn func(element T)) {
	c.elements.Range(func(key, value any) bool {
		fn(key.(T))
		return true
	})
}

// Intersection implements setx.Set.
func (c *ConcurrentSet[T]) Intersection(other setx.Set[T]) setx.Set[T] {
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

// IsEmpty implements setx.Set.
func (c *ConcurrentSet[T]) IsEmpty() bool {
	isEmpty := true
	c.elements.Range(func(key, value any) bool {
		isEmpty = false
		return false // stop iteration
	})
	return isEmpty
}

// IsSubsetOf implements setx.Set.
func (c *ConcurrentSet[T]) IsSubsetOf(other setx.Set[T]) bool {
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

// IsSupersetOf implements setx.Set.
func (c *ConcurrentSet[T]) IsSupersetOf(other setx.Set[T]) bool {
	return other.IsSubsetOf(c)
}

// Remove implements setx.Set.
func (c *ConcurrentSet[T]) Remove(element T) bool {
	_, existed := c.elements.LoadAndDelete(element)
	return existed
}

// Size implements setx.Set.
func (c *ConcurrentSet[T]) Size() int {
	count := 0
	c.elements.Range(func(key, value any) bool {
		count++
		return true
	})
	return count
}

// ToSlice implements setx.Set.
func (c *ConcurrentSet[T]) ToSlice() []T {
	var result []T
	c.elements.Range(func(key, value any) bool {
		result = append(result, key.(T))
		return true
	})
	return result
}

// Union implements setx.Set.
func (c *ConcurrentSet[T]) Union(other setx.Set[T]) setx.Set[T] {
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

// Find implements setx.Set.
func (c *ConcurrentSet[T]) Find(predicate func(T) bool) option.Option[T] {
	var found option.Option[T] = option.None[T]()
	c.elements.Range(func(key, value any) bool {
		element := key.(T)
		if predicate(element) {
			found = option.Some(element)
			return false // stop iteration
		}
		return true // continue iteration
	})
	return found
}

// GetAny implements setx.Set.
func (c *ConcurrentSet[T]) GetAny() option.Option[T] {
	var found option.Option[T] = option.None[T]()
	c.elements.Range(func(key, value any) bool {
		found = option.Some(key.(T))
		return false // stop after first element
	})
	return found
}

// TryRemove implements setx.Set.
func (c *ConcurrentSet[T]) TryRemove(element T) result.Result[T, error] {
	if _, exists := c.elements.LoadAndDelete(element); exists {
		return result.Ok[T, error](element)
	}
	return result.Err[T, error](errors.New("element not found in set"))
}

// Filter implements setx.Set.
func (c *ConcurrentSet[T]) Filter(predicate func(T) bool) setx.Set[T] {
	result := New[T]()
	c.elements.Range(func(key, value any) bool {
		element := key.(T)
		if predicate(element) {
			result.Add(element)
		}
		return true
	})
	return result
}
