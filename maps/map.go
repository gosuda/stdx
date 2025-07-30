package maps

// Map interface defines basic operations for key-value pair storage data structures.
type Map[K comparable, V any] interface {
	// Put stores a key-value pair in the map. If the key already exists, it updates the value and returns the previous value.
	Put(key K, value V) (previousValue V, exists bool)

	// Get returns the value corresponding to the key.
	Get(key K) (value V, exists bool)

	// Remove removes the entry corresponding to the key.
	Remove(key K) (value V, exists bool)

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
}

// Entry represents a key-value pair.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
