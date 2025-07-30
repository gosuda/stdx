// Package option provides a Rust-inspired Option[T] monad for Go.
// Option[T] represents an optional value: every Option is either Some and contains a value, or None, and does not.
package option

import (
	"encoding/json"
	"fmt"
)

// Option represents an optional value: every Option is either Some and contains a value, or None, and does not.
type Option[T any] struct {
	value *T
}

// Some creates an Option containing the given value.
func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

// None creates an empty Option.
func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

// IsSome returns true if the Option is Some.
func (o Option[T]) IsSome() bool {
	return o.value != nil
}

// IsNone returns true if the Option is None.
func (o Option[T]) IsNone() bool {
	return o.value == nil
}

// Unwrap returns the contained value.
// Panics if the Option is None.
func (o Option[T]) Unwrap() T {
	if o.IsNone() {
		panic("called `Option.Unwrap()` on a `None` value")
	}
	return *o.value
}

// UnwrapOr returns the contained value or the provided default.
func (o Option[T]) UnwrapOr(defaultValue T) T {
	if o.IsNone() {
		return defaultValue
	}
	return *o.value
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.IsNone() {
		return f()
	}
	return *o.value
}

// Map transforms the contained value (if any) by applying function f to it.
func Map[T, U any](o Option[T], f func(T) U) Option[U] {
	if o.IsNone() {
		return None[U]()
	}
	return Some(f(*o.value))
}

// FlatMap applies function f to the contained value (if any), or returns None if the Option is None.
// This is useful for chaining operations that return Options.
func FlatMap[T, U any](o Option[T], f func(T) Option[U]) Option[U] {
	if o.IsNone() {
		return None[U]()
	}
	return f(*o.value)
}

// Filter returns None if the Option is None, otherwise calls predicate with the wrapped value and returns:
// - Some(t) if predicate returns true (where t is the wrapped value), and
// - None if predicate returns false.
func (o Option[T]) Filter(predicate func(T) bool) Option[T] {
	if o.IsNone() || !predicate(*o.value) {
		return None[T]()
	}
	return o
}

// Or returns the Option if it contains a value, otherwise returns optb.
func (o Option[T]) Or(optb Option[T]) Option[T] {
	if o.IsSome() {
		return o
	}
	return optb
}

// OrElse returns the Option if it contains a value, otherwise calls f and returns the result.
func (o Option[T]) OrElse(f func() Option[T]) Option[T] {
	if o.IsSome() {
		return o
	}
	return f()
}

// And returns None if the Option is None, otherwise returns optb.
func (o Option[T]) And(optb Option[T]) Option[T] {
	if o.IsNone() {
		return None[T]()
	}
	return optb
}

// AndThen returns None if the Option is None, otherwise calls f with the wrapped value and returns the result.
func AndThen[T, U any](o Option[T], f func(T) Option[U]) Option[U] {
	return FlatMap(o, f)
}

// Match pattern matches on the Option value.
func (o Option[T]) Match(some func(T), none func()) {
	if o.IsSome() {
		some(*o.value)
	} else {
		none()
	}
}

// MatchReturn pattern matches on the Option value and returns a result.
func MatchReturn[T, R any](o Option[T], some func(T) R, none func() R) R {
	if o.IsSome() {
		return some(*o.value)
	}
	return none()
}

// String implements the fmt.Stringer interface.
func (o Option[T]) String() string {
	if o.IsNone() {
		return "None"
	}
	return fmt.Sprintf("Some(%v)", *o.value)
}

// MarshalJSON implements the json.Marshaler interface.
func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.IsNone() {
		return json.Marshal(nil)
	}
	return json.Marshal(*o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (o *Option[T]) UnmarshalJSON(data []byte) error {
	var v *T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if v == nil {
		*o = None[T]()
	} else {
		*o = Some(*v)
	}
	return nil
}

// FromPtr creates an Option from a pointer. If the pointer is nil, returns None.
func FromPtr[T any](ptr *T) Option[T] {
	if ptr == nil {
		return None[T]()
	}
	return Some(*ptr)
}

// ToPtr returns a pointer to the contained value, or nil if None.
func (o Option[T]) ToPtr() *T {
	if o.IsNone() {
		return nil
	}
	value := *o.value
	return &value
}
