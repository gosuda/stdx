package hashmap

import (
	"reflect"

	"github.com/gosuda/stdx/maps"
)

var _ maps.Map[int, string] = (*HashMap[int, string])(nil)

type HashMap[K comparable, V any] struct {
	elements map[K]V
}

func New[K comparable, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{
		elements: make(map[K]V),
	}
}

// Clear implements maps.Map.
func (h *HashMap[K, V]) Clear() {
	h.elements = make(map[K]V)
}

// ContainsKey implements maps.Map.
func (h *HashMap[K, V]) ContainsKey(key K) bool {
	_, exists := h.elements[key]
	return exists
}

// ContainsValue implements maps.Map.
func (h *HashMap[K, V]) ContainsValue(value V) bool {
	for _, v := range h.elements {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Entries implements maps.Map.
func (h *HashMap[K, V]) Entries() []maps.Entry[K, V] {
	result := make([]maps.Entry[K, V], 0, len(h.elements))
	for k, v := range h.elements {
		result = append(result, maps.Entry[K, V]{Key: k, Value: v})
	}
	return result
}

// ForEach implements maps.Map.
func (h *HashMap[K, V]) ForEach(fn func(key K, value V)) {
	for k, v := range h.elements {
		fn(k, v)
	}
}

// Get implements maps.Map.
func (h *HashMap[K, V]) Get(key K) (value V, exists bool) {
	value, exists = h.elements[key]
	return
}

// IsEmpty implements maps.Map.
func (h *HashMap[K, V]) IsEmpty() bool {
	return len(h.elements) == 0
}

// Keys implements maps.Map.
func (h *HashMap[K, V]) Keys() []K {
	result := make([]K, 0, len(h.elements))
	for k := range h.elements {
		result = append(result, k)
	}
	return result
}

// Put implements maps.Map.
func (h *HashMap[K, V]) Put(key K, value V) (previousValue V, exists bool) {
	previousValue, exists = h.elements[key]
	h.elements[key] = value
	return
}

// Remove implements maps.Map.
func (h *HashMap[K, V]) Remove(key K) (value V, exists bool) {
	value, exists = h.elements[key]
	if exists {
		delete(h.elements, key)
	}
	return
}

// Size implements maps.Map.
func (h *HashMap[K, V]) Size() int {
	return len(h.elements)
}

// Values implements maps.Map.
func (h *HashMap[K, V]) Values() []V {
	result := make([]V, 0, len(h.elements))
	for _, v := range h.elements {
		result = append(result, v)
	}
	return result
}
