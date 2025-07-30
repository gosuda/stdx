// Package result provides a Rust-inspired Result[T, E] monad for Go.
// Result[T, E] is the type used for returning and propagating errors.
// It is an enum with the variants, Ok(T), representing success and containing a value, and Err(E), representing error and containing an error value.
package result

import (
	"encoding/json"
	"fmt"

	"github.com/gosuda/stdx/option"
)

// Result represents either success (Ok) or failure (Err).
type Result[T, E any] struct {
	value *T
	err   *E
}

// Ok creates a Result representing success containing the given value.
func Ok[T, E any](value T) Result[T, E] {
	return Result[T, E]{value: &value, err: nil}
}

// Err creates a Result representing failure containing the given error.
func Err[T, E any](err E) Result[T, E] {
	return Result[T, E]{value: nil, err: &err}
}

// IsOk returns true if the Result is Ok.
func (r Result[T, E]) IsOk() bool {
	return r.value != nil
}

// IsErr returns true if the Result is Err.
func (r Result[T, E]) IsErr() bool {
	return r.err != nil
}

// Unwrap returns the contained Ok value.
// Panics if the Result is Err.
func (r Result[T, E]) Unwrap() T {
	if r.IsErr() {
		panic(fmt.Sprintf("called `Result.Unwrap()` on an `Err` value: %v", *r.err))
	}
	return *r.value
}

// UnwrapErr returns the contained Err value.
// Panics if the Result is Ok.
func (r Result[T, E]) UnwrapErr() E {
	if r.IsOk() {
		panic(fmt.Sprintf("called `Result.UnwrapErr()` on an `Ok` value: %v", *r.value))
	}
	return *r.err
}

// UnwrapOr returns the contained Ok value or the provided default.
func (r Result[T, E]) UnwrapOr(defaultValue T) T {
	if r.IsErr() {
		return defaultValue
	}
	return *r.value
}

// UnwrapOrElse returns the contained Ok value or computes it from the error.
func (r Result[T, E]) UnwrapOrElse(f func(E) T) T {
	if r.IsErr() {
		return f(*r.err)
	}
	return *r.value
}

// Map transforms the contained Ok value (if any) by applying function f to it.
func Map[T, U, E any](r Result[T, E], f func(T) U) Result[U, E] {
	if r.IsErr() {
		return Err[U, E](*r.err)
	}
	return Ok[U, E](f(*r.value))
}

// MapErr transforms the contained Err value (if any) by applying function f to it.
func MapErr[T, E, F any](r Result[T, E], f func(E) F) Result[T, F] {
	if r.IsOk() {
		return Ok[T, F](*r.value)
	}
	return Err[T, F](f(*r.err))
}

// FlatMap applies function f to the contained Ok value (if any), or returns the Err if the Result is Err.
// This is useful for chaining operations that return Results.
func FlatMap[T, U, E any](r Result[T, E], f func(T) Result[U, E]) Result[U, E] {
	if r.IsErr() {
		return Err[U, E](*r.err)
	}
	return f(*r.value)
}

// And returns res if the Result is Ok, otherwise returns the Err value of self.
func (r Result[T, E]) And(res Result[T, E]) Result[T, E] {
	if r.IsErr() {
		return r
	}
	return res
}

// AndThen calls f if the Result is Ok, otherwise returns the Err value of self.
func AndThen[T, U, E any](r Result[T, E], f func(T) Result[U, E]) Result[U, E] {
	return FlatMap(r, f)
}

// Or returns the Result if it is Ok, otherwise returns res.
func (r Result[T, E]) Or(res Result[T, E]) Result[T, E] {
	if r.IsOk() {
		return r
	}
	return res
}

// OrElse calls f if the Result is Err, otherwise returns the Ok value of self.
func (r Result[T, E]) OrElse(f func(E) Result[T, E]) Result[T, E] {
	if r.IsErr() {
		return f(*r.err)
	}
	return r
}

// Match pattern matches on the Result value.
func (r Result[T, E]) Match(ok func(T), err func(E)) {
	if r.IsOk() {
		ok(*r.value)
	} else {
		err(*r.err)
	}
}

// MatchReturn pattern matches on the Result value and returns a result.
func MatchReturn[T, E, R any](r Result[T, E], ok func(T) R, err func(E) R) R {
	if r.IsOk() {
		return ok(*r.value)
	}
	return err(*r.err)
}

// Ok returns an Option containing the Ok value, or None if Err.
func (r Result[T, E]) Ok() option.Option[T] {
	if r.IsErr() {
		return option.None[T]()
	}
	return option.Some(*r.value)
}

// Err returns an Option containing the Err value, or None if Ok.
func (r Result[T, E]) Err() option.Option[E] {
	if r.IsOk() {
		return option.None[E]()
	}
	return option.Some(*r.err)
}

// String implements the fmt.Stringer interface.
func (r Result[T, E]) String() string {
	if r.IsOk() {
		return fmt.Sprintf("Ok(%v)", *r.value)
	}
	return fmt.Sprintf("Err(%v)", *r.err)
}

// MarshalJSON implements the json.Marshaler interface.
func (r Result[T, E]) MarshalJSON() ([]byte, error) {
	if r.IsOk() {
		return json.Marshal(map[string]interface{}{"ok": *r.value})
	}
	return json.Marshal(map[string]interface{}{"err": *r.err})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (r *Result[T, E]) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if okData, exists := raw["ok"]; exists {
		var value T
		if err := json.Unmarshal(okData, &value); err != nil {
			return err
		}
		*r = Ok[T, E](value)
		return nil
	}

	if errData, exists := raw["err"]; exists {
		var errValue E
		if err := json.Unmarshal(errData, &errValue); err != nil {
			return err
		}
		*r = Err[T, E](errValue)
		return nil
	}

	return fmt.Errorf("invalid Result JSON format")
}

// Try converts a (value, error) pair into a Result.
func Try[T any](value T, err error) Result[T, error] {
	if err != nil {
		return Err[T, error](err)
	}
	return Ok[T, error](value)
}

// TryWith converts a function that returns (T, error) into a Result.
func TryWith[T any](f func() (T, error)) Result[T, error] {
	value, err := f()
	return Try(value, err)
}
