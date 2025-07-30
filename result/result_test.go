package result_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/gosuda/stdx/result"
)

func TestResult_Ok(t *testing.T) {
	r := result.Ok[int, string](42)

	if !r.IsOk() {
		t.Error("Expected Ok to be Ok")
	}

	if r.IsErr() {
		t.Error("Expected Ok not to be Err")
	}

	if r.Unwrap() != 42 {
		t.Errorf("Expected unwrapped value to be 42, got %d", r.Unwrap())
	}
}

func TestResult_Err(t *testing.T) {
	r := result.Err[int, string]("error")

	if r.IsOk() {
		t.Error("Expected Err not to be Ok")
	}

	if !r.IsErr() {
		t.Error("Expected Err to be Err")
	}

	if r.UnwrapErr() != "error" {
		t.Errorf("Expected unwrapped error to be 'error', got %s", r.UnwrapErr())
	}
}

func TestResult_Unwrap_Panic(t *testing.T) {
	r := result.Err[int, string]("error")

	defer func() {
		if recover() == nil {
			t.Error("Expected Unwrap on Err to panic")
		}
	}()

	r.Unwrap()
}

func TestResult_UnwrapErr_Panic(t *testing.T) {
	r := result.Ok[int, string](42)

	defer func() {
		if recover() == nil {
			t.Error("Expected UnwrapErr on Ok to panic")
		}
	}()

	r.UnwrapErr()
}

func TestResult_UnwrapOr(t *testing.T) {
	ok := result.Ok[int, string](42)
	err := result.Err[int, string]("error")

	if ok.UnwrapOr(0) != 42 {
		t.Error("Expected Ok to unwrap to its value")
	}

	if err.UnwrapOr(123) != 123 {
		t.Error("Expected Err to unwrap to default value")
	}
}

func TestResult_UnwrapOrElse(t *testing.T) {
	ok := result.Ok[int, string](42)
	err := result.Err[int, string]("error")

	if ok.UnwrapOrElse(func(e string) int { return 0 }) != 42 {
		t.Error("Expected Ok to unwrap to its value")
	}

	if err.UnwrapOrElse(func(e string) int { return 123 }) != 123 {
		t.Error("Expected Err to compute value from error")
	}
}

func TestResult_Map(t *testing.T) {
	ok := result.Ok[int, string](42)
	err := result.Err[int, string]("error")

	mapped := result.Map(ok, func(x int) string { return "success" })
	if !mapped.IsOk() || mapped.Unwrap() != "success" {
		t.Error("Expected Ok to map to Ok")
	}

	mappedErr := result.Map(err, func(x int) string { return "never called" })
	if !mappedErr.IsErr() || mappedErr.UnwrapErr() != "error" {
		t.Error("Expected Err to map to Err with same error")
	}
}

func TestResult_MapErr(t *testing.T) {
	ok := result.Ok[int, string](42)
	err := result.Err[int, string]("error")

	mappedOk := result.MapErr(ok, func(e string) int { return 0 })
	if !mappedOk.IsOk() || mappedOk.Unwrap() != 42 {
		t.Error("Expected Ok to remain Ok after MapErr")
	}

	mappedErr := result.MapErr(err, func(e string) int { return 123 })
	if !mappedErr.IsErr() || mappedErr.UnwrapErr() != 123 {
		t.Error("Expected Err to map error value")
	}
}

func TestResult_FlatMap(t *testing.T) {
	ok := result.Ok[int, string](42)
	err := result.Err[int, string]("error")

	flatMapped := result.FlatMap(ok, func(x int) result.Result[string, string] {
		if x > 0 {
			return result.Ok[string, string]("positive")
		}
		return result.Err[string, string]("negative")
	})

	if !flatMapped.IsOk() || flatMapped.Unwrap() != "positive" {
		t.Error("Expected Ok to flat map to Ok")
	}

	flatMappedErr := result.FlatMap(err, func(x int) result.Result[string, string] {
		return result.Ok[string, string]("never called")
	})

	if !flatMappedErr.IsErr() || flatMappedErr.UnwrapErr() != "error" {
		t.Error("Expected Err to flat map to Err with original error")
	}
}

func TestResult_And(t *testing.T) {
	ok1 := result.Ok[int, string](1)
	ok2 := result.Ok[int, string](2)
	err := result.Err[int, string]("error")

	if ok1.And(ok2).Unwrap() != 2 {
		t.Error("Expected Ok.And(Ok) to return second Ok")
	}

	if !ok1.And(err).IsErr() {
		t.Error("Expected Ok.And(Err) to return Err")
	}

	if !err.And(ok2).IsErr() {
		t.Error("Expected Err.And(Ok) to return first Err")
	}
}

func TestResult_Or(t *testing.T) {
	ok1 := result.Ok[int, string](1)
	ok2 := result.Ok[int, string](2)
	err := result.Err[int, string]("error")

	if ok1.Or(ok2).Unwrap() != 1 {
		t.Error("Expected Ok.Or(Ok) to return first Ok")
	}

	if ok1.Or(err).Unwrap() != 1 {
		t.Error("Expected Ok.Or(Err) to return Ok")
	}

	if err.Or(ok2).Unwrap() != 2 {
		t.Error("Expected Err.Or(Ok) to return Ok")
	}
}

func TestResult_Match(t *testing.T) {
	ok := result.Ok[int, string](42)
	err := result.Err[int, string]("error")

	var result string

	ok.Match(
		func(x int) { result = "ok" },
		func(e string) { result = "err" },
	)
	if result != "ok" {
		t.Error("Expected Ok to match ok branch")
	}

	err.Match(
		func(x int) { result = "ok" },
		func(e string) { result = "err" },
	)
	if result != "err" {
		t.Error("Expected Err to match err branch")
	}
}

func TestResult_MatchReturn(t *testing.T) {
	ok := result.Ok[int, string](42)
	err := result.Err[int, string]("error")

	okResult := result.MatchReturn(ok,
		func(x int) string { return "ok" },
		func(e string) string { return "err" },
	)
	if okResult != "ok" {
		t.Error("Expected Ok to return ok result")
	}

	errResult := result.MatchReturn(err,
		func(x int) string { return "ok" },
		func(e string) string { return "err" },
	)
	if errResult != "err" {
		t.Error("Expected Err to return err result")
	}
}

func TestResult_OkOption(t *testing.T) {
	ok := result.Ok[int, string](42)
	err := result.Err[int, string]("error")

	okOpt := ok.Ok()
	if !okOpt.IsSome() || okOpt.Unwrap() != 42 {
		t.Error("Expected Ok to return Some option")
	}

	errOpt := err.Ok()
	if !errOpt.IsNone() {
		t.Error("Expected Err to return None option")
	}
}

func TestResult_ErrOption(t *testing.T) {
	ok := result.Ok[int, string](42)
	err := result.Err[int, string]("error")

	okErrOpt := ok.Err()
	if !okErrOpt.IsNone() {
		t.Error("Expected Ok to return None option for error")
	}

	errErrOpt := err.Err()
	if !errErrOpt.IsSome() || errErrOpt.Unwrap() != "error" {
		t.Error("Expected Err to return Some option for error")
	}
}

func TestResult_String(t *testing.T) {
	ok := result.Ok[int, string](42)
	err := result.Err[int, string]("error")

	if ok.String() != "Ok(42)" {
		t.Errorf("Expected Ok string representation, got %s", ok.String())
	}

	if err.String() != "Err(error)" {
		t.Errorf("Expected Err string representation, got %s", err.String())
	}
}

func TestResult_JSON(t *testing.T) {
	ok := result.Ok[int, string](42)
	err := result.Err[int, string]("error")

	// Test marshaling
	okJSON, jsonErr := json.Marshal(ok)
	if jsonErr != nil {
		t.Fatalf("Failed to marshal Ok: %v", jsonErr)
	}
	if string(okJSON) != `{"ok":42}` {
		t.Errorf("Expected Ok JSON, got %s", string(okJSON))
	}

	errJSON, jsonErr := json.Marshal(err)
	if jsonErr != nil {
		t.Fatalf("Failed to marshal Err: %v", jsonErr)
	}
	if string(errJSON) != `{"err":"error"}` {
		t.Errorf("Expected Err JSON, got %s", string(errJSON))
	}

	// Test unmarshaling
	var unmarshaledOk result.Result[int, string]
	if jsonErr := json.Unmarshal([]byte(`{"ok":42}`), &unmarshaledOk); jsonErr != nil {
		t.Fatalf("Failed to unmarshal Ok: %v", jsonErr)
	}
	if !unmarshaledOk.IsOk() || unmarshaledOk.Unwrap() != 42 {
		t.Error("Failed to unmarshal Ok correctly")
	}

	var unmarshaledErr result.Result[int, string]
	if jsonErr := json.Unmarshal([]byte(`{"err":"error"}`), &unmarshaledErr); jsonErr != nil {
		t.Fatalf("Failed to unmarshal Err: %v", jsonErr)
	}
	if !unmarshaledErr.IsErr() || unmarshaledErr.UnwrapErr() != "error" {
		t.Error("Failed to unmarshal Err correctly")
	}
}

func TestResult_Try(t *testing.T) {
	// Test successful case
	okResult := result.Try(42, nil)
	if !okResult.IsOk() || okResult.Unwrap() != 42 {
		t.Error("Expected Try with nil error to return Ok")
	}

	// Test error case
	testErr := errors.New("test error")
	errResult := result.Try(0, testErr)
	if !errResult.IsErr() || errResult.UnwrapErr() != testErr {
		t.Error("Expected Try with error to return Err")
	}
}

func TestResult_TryWith(t *testing.T) {
	// Test successful case
	okResult := result.TryWith(func() (int, error) {
		return 42, nil
	})
	if !okResult.IsOk() || okResult.Unwrap() != 42 {
		t.Error("Expected TryWith with successful function to return Ok")
	}

	// Test error case
	testErr := errors.New("test error")
	errResult := result.TryWith(func() (int, error) {
		return 0, testErr
	})
	if !errResult.IsErr() || errResult.UnwrapErr() != testErr {
		t.Error("Expected TryWith with error function to return Err")
	}
}
