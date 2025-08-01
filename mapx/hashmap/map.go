package hashmap

import (
	"errors"
	"reflect"

	"github.com/gosuda/stdx/mapx"
	"github.com/gosuda/stdx/option"
	"github.com/gosuda/stdx/result"
)

var _ mapx.Map[int, string] = (*HashMap[int, string])(nil)

type HashMap[K comparable, V any] struct {
	elements map[K]V
}

func New[K comparable, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{
		elements: make(map[K]V),
	}
}

// Clear implements mapx.Map.
func (h *HashMap[K, V]) Clear() {
	h.elements = make(map[K]V)
}

// ContainsKey implements mapx.Map.
func (h *HashMap[K, V]) ContainsKey(key K) bool {
	_, exists := h.elements[key]
	return exists
}

// ContainsValue implements mapx.Map.
func (h *HashMap[K, V]) ContainsValue(value V) bool {
	for _, v := range h.elements {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Entries implements mapx.Map.
func (h *HashMap[K, V]) Entries() []mapx.Entry[K, V] {
	result := make([]mapx.Entry[K, V], 0, len(h.elements))
	for k, v := range h.elements {
		result = append(result, mapx.Entry[K, V]{Key: k, Value: v})
	}
	return result
}

// ForEach implements mapx.Map.
func (h *HashMap[K, V]) ForEach(fn func(key K, value V)) {
	for k, v := range h.elements {
		fn(k, v)
	}
}

// Get implements mapx.Map.
func (h *HashMap[K, V]) Get(key K) option.Option[V] {
	if value, exists := h.elements[key]; exists {
		return option.Some(value)
	}
	return option.None[V]()
}

// IsEmpty implements mapx.Map.
func (h *HashMap[K, V]) IsEmpty() bool {
	return len(h.elements) == 0
}

// Keys implements mapx.Map.
func (h *HashMap[K, V]) Keys() []K {
	result := make([]K, 0, len(h.elements))
	for k := range h.elements {
		result = append(result, k)
	}
	return result
}

// Put implements mapx.Map.
func (h *HashMap[K, V]) Put(key K, value V) option.Option[V] {
	if previousValue, exists := h.elements[key]; exists {
		h.elements[key] = value
		return option.Some(previousValue)
	}
	h.elements[key] = value
	return option.None[V]()
}

// Remove implements mapx.Map.
func (h *HashMap[K, V]) Remove(key K) result.Result[V, error] {
	if value, exists := h.elements[key]; exists {
		delete(h.elements, key)
		return result.Ok[V, error](value)
	}
	return result.Err[V, error](errors.New("key not found"))
}

// Size implements mapx.Map.
func (h *HashMap[K, V]) Size() int {
	return len(h.elements)
}

// Values implements mapx.Map.
func (h *HashMap[K, V]) Values() []V {
	result := make([]V, 0, len(h.elements))
	for _, v := range h.elements {
		result = append(result, v)
	}
	return result
}

// TryGet implements mapx.Map.
func (h *HashMap[K, V]) TryGet(key K) option.Option[V] {
	if value, exists := h.elements[key]; exists {
		return option.Some(value)
	}
	return option.None[V]()
}

// TryRemove implements mapx.Map.
func (h *HashMap[K, V]) TryRemove(key K) result.Result[V, error] {
	if value, exists := h.elements[key]; exists {
		delete(h.elements, key)
		return result.Ok[V, error](value)
	}
	return result.Err[V, error](errors.New("key not found in map"))
}

// FindKey implements mapx.Map.
func (h *HashMap[K, V]) FindKey(value V) option.Option[K] {
	for k, v := range h.elements {
		if reflect.DeepEqual(v, value) {
			return option.Some(k)
		}
	}
	return option.None[K]()
}

// FindEntry implements mapx.Map.
func (h *HashMap[K, V]) FindEntry(predicate func(K, V) bool) option.Option[mapx.Entry[K, V]] {
	for k, v := range h.elements {
		if predicate(k, v) {
			return option.Some(mapx.Entry[K, V]{Key: k, Value: v})
		}
	}
	return option.None[mapx.Entry[K, V]]()
}

// Filter implements mapx.Map.
func (h *HashMap[K, V]) Filter(predicate func(K, V) bool) mapx.Map[K, V] {
	result := New[K, V]()
	for k, v := range h.elements {
		if predicate(k, v) {
			result.Put(k, v)
		}
	}
	return result
}
