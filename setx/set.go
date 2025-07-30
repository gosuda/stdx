package setx

import (
	"github.com/gosuda/stdx/option"
	"github.com/gosuda/stdx/result"
)

// Set interface defines basic operations for set data structures.
type Set[T comparable] interface {
	// Add adds an element to the set. Returns false if it already exists, true if newly added.
	Add(element T) bool

	// Remove removes an element from the set. Returns true if removed, false if not found.
	Remove(element T) bool

	// Contains checks if an element is contained in the set.
	Contains(element T) bool

	// Size returns the size of the set.
	Size() int

	// IsEmpty checks if the set is empty.
	IsEmpty() bool

	// Clear removes all elements from the set.
	Clear()

	// ToSlice returns all elements of the set as a slice.
	ToSlice() []T

	// ForEach executes a function for every element in the set.
	ForEach(fn func(element T))

	// Union returns the union of the current set and another set.
	Union(other Set[T]) Set[T]

	// Intersection returns the intersection of the current set and another set.
	Intersection(other Set[T]) Set[T]

	// Difference returns the difference of the current set minus another set.
	Difference(other Set[T]) Set[T]

	// IsSubsetOf checks if the current set is a subset of another set.
	IsSubsetOf(other Set[T]) bool

	// IsSupersetOf checks if the current set is a superset of another set.
	IsSupersetOf(other Set[T]) bool

	// Find returns the first element that matches the predicate, or None if not found.
	Find(predicate func(T) bool) option.Option[T]

	// GetAny returns any element from the set, or None if the set is empty.
	GetAny() option.Option[T]

	// TryRemove attempts to remove an element and returns it, or an error if not found.
	TryRemove(element T) result.Result[T, error]

	// Filter returns a new set containing only elements that match the predicate.
	Filter(predicate func(T) bool) Set[T]
}
