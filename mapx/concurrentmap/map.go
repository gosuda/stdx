package concurrentmap

import (
	"reflect"
	"sync"

	"github.com/gosuda/stdx/mapx"
)

var _ mapx.Map[int, string] = (*ConcurrentMap[int, string])(nil)

type ConcurrentMap[K comparable, V any] struct {
	elements sync.Map // Using sync.Map for concurrent access
}

func New[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{}
}

// Clear implements mapx.Map.
func (c *ConcurrentMap[K, V]) Clear() {
	c.elements.Clear() // Clear the sync.Map
}

// ContainsKey implements mapx.Map.
func (c *ConcurrentMap[K, V]) ContainsKey(key K) bool {
	_, exists := c.elements.Load(key)
	return exists
}

// ContainsValue implements mapx.Map.
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

// Entries implements mapx.Map.
func (c *ConcurrentMap[K, V]) Entries() []mapx.Entry[K, V] {
	var result []mapx.Entry[K, V]
	c.elements.Range(func(key, value any) bool {
		result = append(result, mapx.Entry[K, V]{
			Key:   key.(K),
			Value: value.(V),
		})
		return true
	})
	return result
}

// ForEach implements mapx.Map.
func (c *ConcurrentMap[K, V]) ForEach(fn func(key K, value V)) {
	c.elements.Range(func(key, value any) bool {
		fn(key.(K), value.(V))
		return true
	})
}

// Get implements mapx.Map.
func (c *ConcurrentMap[K, V]) Get(key K) (value V, exists bool) {
	val, exists := c.elements.Load(key)
	if exists {
		value = val.(V)
	}
	return
}

// IsEmpty implements mapx.Map.
func (c *ConcurrentMap[K, V]) IsEmpty() bool {
	isEmpty := true
	c.elements.Range(func(key, value any) bool {
		isEmpty = false
		return false // stop iteration
	})
	return isEmpty
}

// Keys implements mapx.Map.
func (c *ConcurrentMap[K, V]) Keys() []K {
	var result []K
	c.elements.Range(func(key, value any) bool {
		result = append(result, key.(K))
		return true
	})
	return result
}

// Put implements mapx.Map.
func (c *ConcurrentMap[K, V]) Put(key K, value V) (previousValue V, exists bool) {
	val, loaded := c.elements.LoadOrStore(key, value)
	if loaded {
		previousValue = val.(V)
		exists = true
		c.elements.Store(key, value) // Update with new value
	}
	return
}

// Remove implements mapx.Map.
func (c *ConcurrentMap[K, V]) Remove(key K) (value V, exists bool) {
	val, loaded := c.elements.LoadAndDelete(key)
	if loaded {
		value = val.(V)
		exists = true
	}
	return
}

// Size implements mapx.Map.
func (c *ConcurrentMap[K, V]) Size() int {
	count := 0
	c.elements.Range(func(key, value any) bool {
		count++
		return true
	})
	return count
}

// Values implements mapx.Map.
func (c *ConcurrentMap[K, V]) Values() []V {
	var result []V
	c.elements.Range(func(key, value any) bool {
		result = append(result, value.(V))
		return true
	})
	return result
}
