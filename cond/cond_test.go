package cond

import (
	"strconv"
	"testing"
)

func TestBasicCond(t *testing.T) {
	t.Run("Simple condition matching", func(t *testing.T) {
		x := 10

		result := New[string]().
			WhenValue(func() bool { return x < 5 }, "small").
			WhenValue(func() bool { return x < 15 }, "medium").
			WhenValue(func() bool { return x >= 15 }, "large").
			MustEval()

		if result != "medium" {
			t.Errorf("Expected 'medium', got %v", result)
		}
	})

	t.Run("First matching condition wins", func(t *testing.T) {
		x := 10

		result := New[string]().
			WhenValue(func() bool { return x > 5 }, "first").
			WhenValue(func() bool { return x > 0 }, "second").
			WhenValue(func() bool { return x == 10 }, "third").
			MustEval()

		if result != "first" {
			t.Errorf("Expected 'first', got %v", result)
		}
	})

	t.Run("No condition matches", func(t *testing.T) {
		x := 10

		result, ok := New[string]().
			WhenValue(func() bool { return x > 20 }, "large").
			WhenValue(func() bool { return x < 5 }, "small").
			Eval()

		if ok {
			t.Error("Expected no match, but got a result")
		}

		if result != "" {
			t.Errorf("Expected empty string, got %v", result)
		}
	})

	t.Run("Else clause", func(t *testing.T) {
		x := 10

		result := New[string]().
			WhenValue(func() bool { return x > 20 }, "large").
			WhenValue(func() bool { return x < 5 }, "small").
			ElseValue("default").
			MustEval()

		if result != "default" {
			t.Errorf("Expected 'default', got %v", result)
		}
	})
}

func TestLazyEvaluation(t *testing.T) {
	t.Run("Lazy evaluation with side effects", func(t *testing.T) {
		executed := false
		x := 5

		result := New[string]().
			WhenValue(func() bool { return x == 5 }, "found").
			When(func() bool { return x > 0 }, func() string {
				executed = true
				return "side effect"
			}).
			MustEval()

		if result != "found" {
			t.Errorf("Expected 'found', got %v", result)
		}

		if executed {
			t.Error("Second clause should not have been executed due to lazy evaluation")
		}
	})

	t.Run("Expensive computation only when needed", func(t *testing.T) {
		computeCount := 0
		expensiveComputation := func() string {
			computeCount++
			return "expensive result"
		}

		x := 1
		result := New[string]().
			WhenValue(func() bool { return x == 1 }, "quick").
			When(func() bool { return x > 0 }, expensiveComputation).
			MustEval()

		if result != "quick" {
			t.Errorf("Expected 'quick', got %v", result)
		}

		if computeCount != 0 {
			t.Error("Expensive computation should not have been called")
		}
	})
}

func TestCondFunc(t *testing.T) {
	t.Run("Standalone CondFunc", func(t *testing.T) {
		x := 7

		result, ok := CondFunc(
			WhenValue(func() bool { return x < 5 }, "small"),
			WhenValue(func() bool { return x < 10 }, "medium"),
			WhenValue(func() bool { return x >= 10 }, "large"),
		)

		if !ok {
			t.Error("Expected a match")
		}

		if result != "medium" {
			t.Errorf("Expected 'medium', got %v", result)
		}
	})

	t.Run("MustCondFunc with panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic, but didn't get one")
			}
		}()

		x := 10
		MustCondFunc(
			WhenValue(func() bool { return x > 20 }, "large"),
			WhenValue(func() bool { return x < 5 }, "small"),
		)
	})
}

func TestSwitchCond(t *testing.T) {
	t.Run("String switch", func(t *testing.T) {
		color := "red"

		result := Switch[string, string](color, func(a, b string) bool { return a == b }).
			CaseValue("red", "stop").
			CaseValue("yellow", "caution").
			CaseValue("green", "go").
			DefaultValue("unknown").
			MustEval()

		if result != "stop" {
			t.Errorf("Expected 'stop', got %v", result)
		}
	})

	t.Run("Integer switch with computation", func(t *testing.T) {
		number := 2

		result := Switch[int, string](number, func(a, b int) bool { return a == b }).
			Case(1, func() string { return "one" }).
			Case(2, func() string { return "two" }).
			Case(3, func() string { return "three" }).
			Default(func() string { return "other: " + strconv.Itoa(number) }).
			MustEval()

		if result != "two" {
			t.Errorf("Expected 'two', got %v", result)
		}
	})

	t.Run("Switch with default", func(t *testing.T) {
		value := 99

		result := Switch[int, string](value, func(a, b int) bool { return a == b }).
			CaseValue(1, "one").
			CaseValue(2, "two").
			DefaultValue("unknown").
			MustEval()

		if result != "unknown" {
			t.Errorf("Expected 'unknown', got %v", result)
		}
	})

	t.Run("Switch without default - no match", func(t *testing.T) {
		value := 99

		result, ok := Switch[int, string](value, func(a, b int) bool { return a == b }).
			CaseValue(1, "one").
			CaseValue(2, "two").
			Eval()

		if ok {
			t.Error("Expected no match")
		}

		if result != "" {
			t.Errorf("Expected empty string, got %v", result)
		}
	})
}

func TestGuardAndUnless(t *testing.T) {
	t.Run("Guard clause with side effect", func(t *testing.T) {
		executed := false
		x := 5

		New[struct{}]().
			Add(Guard(func() bool { return x == 5 }, func() {
				executed = true
			})).
			MustEval()

		if !executed {
			t.Error("Guard clause should have executed")
		}
	})

	t.Run("Unless clause", func(t *testing.T) {
		x := 3

		result := New[string]().
			UnlessValue(func() bool { return x > 10 }, "not large").
			UnlessValue(func() bool { return x < 0 }, "not negative").
			ElseValue("fallback").
			MustEval()

		if result != "not large" {
			t.Errorf("Expected 'not large', got %v", result)
		}
	})

	t.Run("Unless with false condition", func(t *testing.T) {
		x := 15

		result := New[string]().
			UnlessValue(func() bool { return x > 10 }, "not large").   // x > 10 is true, so this won't match
			UnlessValue(func() bool { return x < 0 }, "not negative"). // x < 0 is false, so this will match
			ElseValue("fallback").
			MustEval()

		if result != "not negative" {
			t.Errorf("Expected 'not negative', got %v", result)
		}
	})
}

func TestComplexConditions(t *testing.T) {
	t.Run("Multiple conditions", func(t *testing.T) {
		age := 25
		hasLicense := true
		hasInsurance := true

		result := New[string]().
			WhenValue(func() bool { return age < 18 }, "too young").
			WhenValue(func() bool { return !hasLicense }, "no license").
			WhenValue(func() bool { return !hasInsurance }, "no insurance").
			WhenValue(func() bool { return age >= 18 && hasLicense && hasInsurance }, "can drive").
			ElseValue("unknown status").
			MustEval()

		if result != "can drive" {
			t.Errorf("Expected 'can drive', got %v", result)
		}
	})

	t.Run("Nested conditions", func(t *testing.T) {
		x := 10
		y := 20

		result := New[string]().
			When(func() bool { return x > 5 }, func() string {
				return New[string]().
					WhenValue(func() bool { return y > 15 }, "x>5 and y>15").
					WhenValue(func() bool { return y <= 15 }, "x>5 and y<=15").
					MustEval()
			}).
			ElseValue("x<=5").
			MustEval()

		if result != "x>5 and y>15" {
			t.Errorf("Expected 'x>5 and y>15', got %v", result)
		}
	})
}

func TestErrorHandling(t *testing.T) {
	t.Run("MustEval panics on no match", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic")
			}
		}()

		New[string]().
			WhenValue(func() bool { return false }, "never").
			MustEval()
	})

	t.Run("EvalOr returns default", func(t *testing.T) {
		result := New[string]().
			WhenValue(func() bool { return false }, "never").
			EvalOr("default value")

		if result != "default value" {
			t.Errorf("Expected 'default value', got %v", result)
		}
	})
}

func TestChaining(t *testing.T) {
	t.Run("Method chaining", func(t *testing.T) {
		score := 85

		grade := New[string]().
			WhenValue(func() bool { return score >= 90 }, "A").
			WhenValue(func() bool { return score >= 80 }, "B").
			WhenValue(func() bool { return score >= 70 }, "C").
			WhenValue(func() bool { return score >= 60 }, "D").
			ElseValue("F").
			MustEval()

		if grade != "B" {
			t.Errorf("Expected 'B', got %v", grade)
		}
	})
}

func TestTypeInference(t *testing.T) {
	t.Run("Different return types", func(t *testing.T) {
		// Integer result
		intResult := New[int]().
			WhenValue(func() bool { return true }, 42).
			MustEval()

		if intResult != 42 {
			t.Errorf("Expected 42, got %v", intResult)
		}

		// Boolean result
		boolResult := New[bool]().
			WhenValue(func() bool { return true }, true).
			MustEval()

		if !boolResult {
			t.Errorf("Expected true, got %v", boolResult)
		}

		// Slice result
		sliceResult := New[[]string]().
			WhenValue(func() bool { return true }, []string{"a", "b", "c"}).
			MustEval()

		if len(sliceResult) != 3 {
			t.Errorf("Expected slice of length 3, got %v", sliceResult)
		}
	})
}
