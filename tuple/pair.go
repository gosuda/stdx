package tuple

// Pair represents a tuple with two values
type Pair[T, U any] struct {
	first  T
	second U
}

// NewPair creates a new Pair with the given values
func NewPair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{first: first, second: second}
}

// First returns the first value of the pair
func (p Pair[T, U]) First() T {
	return p.first
}

// Second returns the second value of the pair
func (p Pair[T, U]) Second() U {
	return p.second
}

// Map applies a function to both values and returns a new Pair
func (p Pair[T, U]) Map(f func(T, U) (T, U)) Pair[T, U] {
	first, second := f(p.first, p.second)
	return NewPair(first, second)
}

// MapFirst applies a function to the first value and returns a new Pair
func MapFirst[T, U, V any](p Pair[T, U], f func(T) V) Pair[V, U] {
	return NewPair(f(p.first), p.second)
}

// MapSecond applies a function to the second value and returns a new Pair
func MapSecond[T, U, V any](p Pair[T, U], f func(U) V) Pair[T, V] {
	return NewPair(p.first, f(p.second))
}

// Swap returns a new Pair with swapped values
func (p Pair[T, U]) Swap() Pair[U, T] {
	return NewPair(p.second, p.first)
}

// Apply applies a function that takes both values and returns a result
func Apply[T, U, R any](p Pair[T, U], f func(T, U) R) R {
	return f(p.first, p.second)
}

// Curry converts a function that takes two arguments into a curried function
func Curry[T, U, R any](f func(T, U) R) func(T) func(U) R {
	return func(t T) func(U) R {
		return func(u U) R {
			return f(t, u)
		}
	}
}

// Uncurry converts a curried function into a function that takes a Pair
func Uncurry[T, U, R any](f func(T) func(U) R) func(Pair[T, U]) R {
	return func(p Pair[T, U]) R {
		return f(p.first)(p.second)
	}
}

// Equal compares two pairs for equality
func (p Pair[T, U]) Equal(other Pair[T, U], eqT func(T, T) bool, eqU func(U, U) bool) bool {
	return eqT(p.first, other.first) && eqU(p.second, other.second)
}

// String returns a string representation of the pair
func (p Pair[T, U]) String() string {
	return "(" + toString(p.first) + ", " + toString(p.second) + ")"
}

// helper function for string conversion
func toString[T any](v T) string {
	switch val := any(v).(type) {
	case string:
		return val
	default:
		return ""
	}
}
