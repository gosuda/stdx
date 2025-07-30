package concurrentmap

import (
	"reflect"
	"sync"

	"github.com/gosuda/stdx/maps"
)

var _ maps.Map[int, string] = (*ConcurrentMap[int, string])(nil)

type ConcurrentMap[K comparable, V any] struct {
	elements sync.Map // Using sync.Map for concurrent access
}

func New[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{}
}

// Clear implements maps.Map.
func (c *ConcurrentMap[K, V]) Clear() {
	c.elements.Clear() // Clear the sync.Map
}

// ContainsKey implements maps.Map.
func (c *ConcurrentMap[K, V]) ContainsKey(key K) bool {
	_, exists := c.elements.Load(key)
	return exists
}

// ContainsValue implements maps.Map.
func (c *ConcurrentMap[K, V]) ContainsValue(value V) bool {
	found := false
	c.elements.Range(func(key, val any) bool {
		if reflect.DeepEqual(val.(V), value) {
			found = true
			return false // stop iteration
		}
		return true
	})
	return found
}

// Entries implements maps.Map.
func (c *ConcurrentMap[K, V]) Entries() []maps.Entry[K, V] {
	var result []maps.Entry[K, V]
	c.elements.Range(func(key, value any) bool {
		result = append(result, maps.Entry[K, V]{
			Key:   key.(K),
			Value: value.(V),
		})
		return true
	})
	return result
}

// ForEach implements maps.Map.
func (c *ConcurrentMap[K, V]) ForEach(fn func(key K, value V)) {
	c.elements.Range(func(key, value any) bool {
		fn(key.(K), value.(V))
		return true
	})
}

// Get implements maps.Map.
func (c *ConcurrentMap[K, V]) Get(key K) (value V, exists bool) {
	val, exists := c.elements.Load(key)
	if exists {
		value = val.(V)
	}
	return
}

// IsEmpty implements maps.Map.
func (c *ConcurrentMap[K, V]) IsEmpty() bool {
	isEmpty := true
	c.elements.Range(func(key, value any) bool {
		isEmpty = false
		return false // stop iteration
	})
	return isEmpty
}

// Keys implements maps.Map.
func (c *ConcurrentMap[K, V]) Keys() []K {
	var result []K
	c.elements.Range(func(key, value any) bool {
		result = append(result, key.(K))
		return true
	})
	return result
}

// Put implements maps.Map.
func (c *ConcurrentMap[K, V]) Put(key K, value V) (previousValue V, exists bool) {
	val, loaded := c.elements.LoadOrStore(key, value)
	if loaded {
		previousValue = val.(V)
		exists = true
		c.elements.Store(key, value) // Update with new value
	}
	return
}

// Remove implements maps.Map.
func (c *ConcurrentMap[K, V]) Remove(key K) (value V, exists bool) {
	val, loaded := c.elements.LoadAndDelete(key)
	if loaded {
		value = val.(V)
		exists = true
	}
	return
}

// Size implements maps.Map.
func (c *ConcurrentMap[K, V]) Size() int {
	count := 0
	c.elements.Range(func(key, value any) bool {
		count++
		return true
	})
	return count
}

// Values implements maps.Map.
func (c *ConcurrentMap[K, V]) Values() []V {
	var result []V
	c.elements.Range(func(key, value any) bool {
		result = append(result, value.(V))
		return true
	})
	return result
}
