package option_test

import (
	"encoding/json"
	"testing"

	"github.com/gosuda/stdx/option"
)

func TestOption_Some(t *testing.T) {
	opt := option.Some(42)

	if !opt.IsSome() {
		t.Error("Expected Some to be Some")
	}

	if opt.IsNone() {
		t.Error("Expected Some not to be None")
	}

	if opt.Unwrap() != 42 {
		t.Errorf("Expected unwrapped value to be 42, got %d", opt.Unwrap())
	}
}

func TestOption_None(t *testing.T) {
	opt := option.None[int]()

	if opt.IsSome() {
		t.Error("Expected None not to be Some")
	}

	if !opt.IsNone() {
		t.Error("Expected None to be None")
	}
}

func TestOption_Unwrap_Panic(t *testing.T) {
	opt := option.None[int]()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected Unwrap on None to panic")
		}
	}()

	opt.Unwrap()
}

func TestOption_UnwrapOr(t *testing.T) {
	some := option.Some(42)
	none := option.None[int]()

	if some.UnwrapOr(0) != 42 {
		t.Error("Expected Some to unwrap to its value")
	}

	if none.UnwrapOr(123) != 123 {
		t.Error("Expected None to unwrap to default value")
	}
}

func TestOption_UnwrapOrElse(t *testing.T) {
	some := option.Some(42)
	none := option.None[int]()

	if some.UnwrapOrElse(func() int { return 0 }) != 42 {
		t.Error("Expected Some to unwrap to its value")
	}

	if none.UnwrapOrElse(func() int { return 123 }) != 123 {
		t.Error("Expected None to compute default value")
	}
}

func TestOption_Map(t *testing.T) {
	some := option.Some(42)
	none := option.None[int]()

	mapped := option.Map(some, func(x int) string { return "value" })
	if !mapped.IsSome() || mapped.Unwrap() != "value" {
		t.Error("Expected Some to map to Some")
	}

	mappedNone := option.Map(none, func(x int) string { return "value" })
	if !mappedNone.IsNone() {
		t.Error("Expected None to map to None")
	}
}

func TestOption_FlatMap(t *testing.T) {
	some := option.Some(42)
	none := option.None[int]()

	flatMapped := option.FlatMap(some, func(x int) option.Option[string] {
		if x > 0 {
			return option.Some("positive")
		}
		return option.None[string]()
	})

	if !flatMapped.IsSome() || flatMapped.Unwrap() != "positive" {
		t.Error("Expected Some to flat map to Some")
	}

	flatMappedNone := option.FlatMap(none, func(x int) option.Option[string] {
		return option.Some("never called")
	})

	if !flatMappedNone.IsNone() {
		t.Error("Expected None to flat map to None")
	}
}

func TestOption_Filter(t *testing.T) {
	some := option.Some(42)

	filtered := some.Filter(func(x int) bool { return x > 0 })
	if !filtered.IsSome() {
		t.Error("Expected filtered Some to remain Some when predicate is true")
	}

	filteredOut := some.Filter(func(x int) bool { return x < 0 })
	if !filteredOut.IsNone() {
		t.Error("Expected filtered Some to become None when predicate is false")
	}

	none := option.None[int]()
	filteredNone := none.Filter(func(x int) bool { return true })
	if !filteredNone.IsNone() {
		t.Error("Expected filtered None to remain None")
	}
}

func TestOption_Or(t *testing.T) {
	some1 := option.Some(1)
	some2 := option.Some(2)
	none := option.None[int]()

	if some1.Or(some2).Unwrap() != 1 {
		t.Error("Expected Some.Or(Some) to return first Some")
	}

	if some1.Or(none).Unwrap() != 1 {
		t.Error("Expected Some.Or(None) to return Some")
	}

	if none.Or(some2).Unwrap() != 2 {
		t.Error("Expected None.Or(Some) to return Some")
	}

	if !none.Or(option.None[int]()).IsNone() {
		t.Error("Expected None.Or(None) to return None")
	}
}

func TestOption_And(t *testing.T) {
	some1 := option.Some(1)
	some2 := option.Some(2)
	none := option.None[int]()

	if some1.And(some2).Unwrap() != 2 {
		t.Error("Expected Some.And(Some) to return second Some")
	}

	if !some1.And(none).IsNone() {
		t.Error("Expected Some.And(None) to return None")
	}

	if !none.And(some2).IsNone() {
		t.Error("Expected None.And(Some) to return None")
	}
}

func TestOption_Match(t *testing.T) {
	some := option.Some(42)
	none := option.None[int]()

	var result string

	some.Match(
		func(x int) { result = "some" },
		func() { result = "none" },
	)
	if result != "some" {
		t.Error("Expected Some to match some branch")
	}

	none.Match(
		func(x int) { result = "some" },
		func() { result = "none" },
	)
	if result != "none" {
		t.Error("Expected None to match none branch")
	}
}

func TestOption_MatchReturn(t *testing.T) {
	some := option.Some(42)
	none := option.None[int]()

	someResult := option.MatchReturn(some,
		func(x int) string { return "some" },
		func() string { return "none" },
	)
	if someResult != "some" {
		t.Error("Expected Some to return some result")
	}

	noneResult := option.MatchReturn(none,
		func(x int) string { return "some" },
		func() string { return "none" },
	)
	if noneResult != "none" {
		t.Error("Expected None to return none result")
	}
}

func TestOption_String(t *testing.T) {
	some := option.Some(42)
	none := option.None[int]()

	if some.String() != "Some(42)" {
		t.Errorf("Expected Some string representation, got %s", some.String())
	}

	if none.String() != "None" {
		t.Errorf("Expected None string representation, got %s", none.String())
	}
}

func TestOption_JSON(t *testing.T) {
	some := option.Some(42)
	none := option.None[int]()

	// Test marshaling
	someJSON, err := json.Marshal(some)
	if err != nil {
		t.Fatalf("Failed to marshal Some: %v", err)
	}
	if string(someJSON) != "42" {
		t.Errorf("Expected Some JSON to be '42', got %s", string(someJSON))
	}

	noneJSON, err := json.Marshal(none)
	if err != nil {
		t.Fatalf("Failed to marshal None: %v", err)
	}
	if string(noneJSON) != "null" {
		t.Errorf("Expected None JSON to be 'null', got %s", string(noneJSON))
	}

	// Test unmarshaling
	var unmarshaledSome option.Option[int]
	if err := json.Unmarshal([]byte("42"), &unmarshaledSome); err != nil {
		t.Fatalf("Failed to unmarshal Some: %v", err)
	}
	if !unmarshaledSome.IsSome() || unmarshaledSome.Unwrap() != 42 {
		t.Error("Failed to unmarshal Some correctly")
	}

	var unmarshaledNone option.Option[int]
	if err := json.Unmarshal([]byte("null"), &unmarshaledNone); err != nil {
		t.Fatalf("Failed to unmarshal None: %v", err)
	}
	if !unmarshaledNone.IsNone() {
		t.Error("Failed to unmarshal None correctly")
	}
}

func TestOption_FromPtr(t *testing.T) {
	value := 42
	ptr := &value
	var nilPtr *int

	someFromPtr := option.FromPtr(ptr)
	if !someFromPtr.IsSome() || someFromPtr.Unwrap() != 42 {
		t.Error("Expected FromPtr with valid pointer to create Some")
	}

	noneFromPtr := option.FromPtr(nilPtr)
	if !noneFromPtr.IsNone() {
		t.Error("Expected FromPtr with nil pointer to create None")
	}
}

func TestOption_ToPtr(t *testing.T) {
	some := option.Some(42)
	none := option.None[int]()

	somePtr := some.ToPtr()
	if somePtr == nil || *somePtr != 42 {
		t.Error("Expected Some.ToPtr() to return valid pointer")
	}

	nonePtr := none.ToPtr()
	if nonePtr != nil {
		t.Error("Expected None.ToPtr() to return nil")
	}
}
