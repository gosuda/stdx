package concurrentmap

import (
	"errors"
	"reflect"
	"sync"

	"github.com/gosuda/stdx/mapx"
	"github.com/gosuda/stdx/option"
	"github.com/gosuda/stdx/result"
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
func (c *ConcurrentMap[K, V]) Get(key K) option.Option[V] {
	if val, exists := c.elements.Load(key); exists {
		return option.Some(val.(V))
	}
	return option.None[V]()
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
func (c *ConcurrentMap[K, V]) Put(key K, value V) option.Option[V] {
	if val, loaded := c.elements.LoadOrStore(key, value); loaded {
		// Key existed, update with new value and return previous value
		c.elements.Store(key, value)
		return option.Some(val.(V))
	}
	// Key didn't exist, new entry created
	return option.None[V]()
}

// Remove implements mapx.Map.
func (c *ConcurrentMap[K, V]) Remove(key K) result.Result[V, error] {
	if val, loaded := c.elements.LoadAndDelete(key); loaded {
		return result.Ok[V, error](val.(V))
	}
	return result.Err[V, error](errors.New("key not found"))
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

// FindKey implements mapx.Map.
func (c *ConcurrentMap[K, V]) FindKey(value V) option.Option[K] {
	var found option.Option[K] = option.None[K]()
	c.elements.Range(func(key, val any) bool {
		if reflect.DeepEqual(val.(V), value) {
			found = option.Some(key.(K))
			return false // stop iteration
		}
		return true
	})
	return found
}

// FindEntry implements mapx.Map.
func (c *ConcurrentMap[K, V]) FindEntry(predicate func(K, V) bool) option.Option[mapx.Entry[K, V]] {
	var found option.Option[mapx.Entry[K, V]] = option.None[mapx.Entry[K, V]]()
	c.elements.Range(func(key, val any) bool {
		k := key.(K)
		v := val.(V)
		if predicate(k, v) {
			found = option.Some(mapx.Entry[K, V]{Key: k, Value: v})
			return false // stop iteration
		}
		return true
	})
	return found
}

// Filter implements mapx.Map.
func (c *ConcurrentMap[K, V]) Filter(predicate func(K, V) bool) mapx.Map[K, V] {
	result := New[K, V]()
	c.elements.Range(func(key, val any) bool {
		k := key.(K)
		v := val.(V)
		if predicate(k, v) {
			result.Put(k, v)
		}
		return true
	})
	return result
}
