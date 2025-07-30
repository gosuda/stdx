package tuple

// Triple represents a tuple with three values
type Triple[T, U, V any] struct {
	first  T
	second U
	third  V
}

// NewTriple creates a new Triple with the given values
func NewTriple[T, U, V any](first T, second U, third V) Triple[T, U, V] {
	return Triple[T, U, V]{first: first, second: second, third: third}
}

// First returns the first value of the triple
func (t Triple[T, U, V]) First() T {
	return t.first
}

// Second returns the second value of the triple
func (t Triple[T, U, V]) Second() U {
	return t.second
}

// Third returns the third value of the triple
func (t Triple[T, U, V]) Third() V {
	return t.third
}

// Map applies a function to all three values and returns a new Triple
func (t Triple[T, U, V]) Map(f func(T, U, V) (T, U, V)) Triple[T, U, V] {
	first, second, third := f(t.first, t.second, t.third)
	return NewTriple(first, second, third)
}

// MapFirstTriple applies a function to the first value and returns a new Triple
func MapFirstTriple[T, U, V, W any](t Triple[T, U, V], f func(T) W) Triple[W, U, V] {
	return NewTriple(f(t.first), t.second, t.third)
}

// MapSecondTriple applies a function to the second value and returns a new Triple
func MapSecondTriple[T, U, V, W any](t Triple[T, U, V], f func(U) W) Triple[T, W, V] {
	return NewTriple(t.first, f(t.second), t.third)
}

// MapThirdTriple applies a function to the third value and returns a new Triple
func MapThirdTriple[T, U, V, W any](t Triple[T, U, V], f func(V) W) Triple[T, U, W] {
	return NewTriple(t.first, t.second, f(t.third))
}

// Apply applies a function that takes all three values and returns a result
func ApplyTriple[T, U, V, R any](t Triple[T, U, V], f func(T, U, V) R) R {
	return f(t.first, t.second, t.third)
}

// ToTuple converts a Triple to a Pair by dropping the third element
func (t Triple[T, U, V]) ToPair() Pair[T, U] {
	return NewPair(t.first, t.second)
}

// FromPair creates a Triple from a Pair and a third value
func FromPair[T, U, V any](p Pair[T, U], third V) Triple[T, U, V] {
	return NewTriple(p.First(), p.Second(), third)
}

// RotateLeft rotates the values to the left (first becomes last)
func (t Triple[T, U, V]) RotateLeft() Triple[U, V, T] {
	return NewTriple(t.second, t.third, t.first)
}

// RotateRight rotates the values to the right (last becomes first)
func (t Triple[T, U, V]) RotateRight() Triple[V, T, U] {
	return NewTriple(t.third, t.first, t.second)
}

// Curry converts a function that takes three arguments into a curried function
func CurryTriple[T, U, V, R any](f func(T, U, V) R) func(T) func(U) func(V) R {
	return func(t T) func(U) func(V) R {
		return func(u U) func(V) R {
			return func(v V) R {
				return f(t, u, v)
			}
		}
	}
}

// Uncurry converts a curried function into a function that takes a Triple
func UncurryTriple[T, U, V, R any](f func(T) func(U) func(V) R) func(Triple[T, U, V]) R {
	return func(t Triple[T, U, V]) R {
		return f(t.first)(t.second)(t.third)
	}
}

// Equal compares two triples for equality
func (t Triple[T, U, V]) Equal(other Triple[T, U, V], eqT func(T, T) bool, eqU func(U, U) bool, eqV func(V, V) bool) bool {
	return eqT(t.first, other.first) && eqU(t.second, other.second) && eqV(t.third, other.third)
}

// String returns a string representation of the triple
func (t Triple[T, U, V]) String() string {
	return "(" + toString(t.first) + ", " + toString(t.second) + ", " + toString(t.third) + ")"
}
