package mapx

import (
	"github.com/gosuda/stdx/option"
	"github.com/gosuda/stdx/result"
)

// Map interface defines basic operations for key-value pair storage data structures.
type Map[K comparable, V any] interface {
	// Put stores a key-value pair in the map. Returns Some(previousValue) if key existed, None otherwise.
	Put(key K, value V) option.Option[V]

	// Get returns Some(value) if key exists, None otherwise.
	Get(key K) option.Option[V]

	// Remove removes the entry corresponding to the key. Returns Ok(removedValue) if successful, Err if key not found.
	Remove(key K) result.Result[V, error]

	// ContainsKey checks if the key exists in the map.
	ContainsKey(key K) bool

	// ContainsValue checks if the value exists in the map.
	ContainsValue(value V) bool

	// Size returns the size of the map.
	Size() int

	// IsEmpty checks if the map is empty.
	IsEmpty() bool

	// Clear removes all entries from the map.
	Clear()

	// Keys returns all keys in the map as a slice.
	Keys() []K

	// Values returns all values in the map as a slice.
	Values() []V

	// Entries returns all key-value pairs in the map.
	Entries() []Entry[K, V]

	// ForEach executes a function for every key-value pair in the map.
	ForEach(fn func(key K, value V))

	// FindKey returns the first key that maps to the given value, or None if not found.
	FindKey(value V) option.Option[K]

	// FindEntry returns the first entry that matches the predicate, or None if not found.
	FindEntry(predicate func(K, V) bool) option.Option[Entry[K, V]]

	// Filter returns a new map containing only entries that match the predicate.
	Filter(predicate func(K, V) bool) Map[K, V]
}

// Entry represents a key-value pair.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
