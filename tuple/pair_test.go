package tuple

import (
	"strconv"
	"testing"
)

func TestPairBasicOperations(t *testing.T) {
	t.Run("NewPair and getters", func(t *testing.T) {
		p := NewPair(42, "hello")

		if p.First() != 42 {
			t.Errorf("Expected First() to be 42, got %v", p.First())
		}

		if p.Second() != "hello" {
			t.Errorf("Expected Second() to be 'hello', got %v", p.Second())
		}
	})

	t.Run("Different types", func(t *testing.T) {
		// Integer pair
		intPair := NewPair(10, 20)
		if intPair.First() != 10 || intPair.Second() != 20 {
			t.Errorf("Integer pair failed: got (%v, %v)", intPair.First(), intPair.Second())
		}

		// String pair
		stringPair := NewPair("foo", "bar")
		if stringPair.First() != "foo" || stringPair.Second() != "bar" {
			t.Errorf("String pair failed: got (%v, %v)", stringPair.First(), stringPair.Second())
		}

		// Mixed types
		mixedPair := NewPair(3.14, true)
		if mixedPair.First() != 3.14 || mixedPair.Second() != true {
			t.Errorf("Mixed pair failed: got (%v, %v)", mixedPair.First(), mixedPair.Second())
		}

		// Complex types
		slicePair := NewPair([]int{1, 2, 3}, map[string]int{"a": 1})
		if len(slicePair.First()) != 3 || slicePair.Second()["a"] != 1 {
			t.Error("Complex types pair failed")
		}
	})
}

func TestPairMap(t *testing.T) {
	t.Run("Map with same types", func(t *testing.T) {
		p := NewPair(5, 10)
		doubled := p.Map(func(a, b int) (int, int) {
			return a * 2, b * 2
		})

		if doubled.First() != 10 || doubled.Second() != 20 {
			t.Errorf("Map failed: expected (10, 20), got (%v, %v)", doubled.First(), doubled.Second())
		}
	})

	t.Run("Map with string manipulation", func(t *testing.T) {
		p := NewPair("hello", "world")
		uppercased := p.Map(func(a, b string) (string, string) {
			return a + "!", b + "!"
		})

		if uppercased.First() != "hello!" || uppercased.Second() != "world!" {
			t.Errorf("String map failed: got (%v, %v)", uppercased.First(), uppercased.Second())
		}
	})

	t.Run("Map with complex transformation", func(t *testing.T) {
		p := NewPair(1, 2)
		transformed := p.Map(func(a, b int) (int, int) {
			return a*a + b, a + b*b
		})

		expected1, expected2 := 1*1+2, 1+2*2 // 3, 5
		if transformed.First() != expected1 || transformed.Second() != expected2 {
			t.Errorf("Complex map failed: expected (%d, %d), got (%v, %v)",
				expected1, expected2, transformed.First(), transformed.Second())
		}
	})
}

func TestPairMapFirst(t *testing.T) {
	t.Run("MapFirst type conversion", func(t *testing.T) {
		p := NewPair(42, "hello")
		result := MapFirst(p, func(i int) string {
			return "number: " + strconv.Itoa(i)
		})

		if result.First() != "number: 42" {
			t.Errorf("MapFirst failed: expected 'number: 42', got %v", result.First())
		}

		if result.Second() != "hello" {
			t.Errorf("MapFirst should preserve second value: expected 'hello', got %v", result.Second())
		}
	})

	t.Run("MapFirst with different transformations", func(t *testing.T) {
		p := NewPair(10, 3.14)

		// Convert int to bool
		boolResult := MapFirst(p, func(i int) bool {
			return i > 5
		})

		if !boolResult.First() {
			t.Error("MapFirst bool conversion failed")
		}

		// Convert int to slice
		sliceResult := MapFirst(p, func(i int) []int {
			return make([]int, i)
		})

		if len(sliceResult.First()) != 10 {
			t.Errorf("MapFirst slice conversion failed: expected length 10, got %d", len(sliceResult.First()))
		}
	})
}

func TestPairMapSecond(t *testing.T) {
	t.Run("MapSecond type conversion", func(t *testing.T) {
		p := NewPair("hello", 42)
		result := MapSecond(p, func(i int) string {
			return strconv.Itoa(i) + "!"
		})

		if result.First() != "hello" {
			t.Errorf("MapSecond should preserve first value: expected 'hello', got %v", result.First())
		}

		if result.Second() != "42!" {
			t.Errorf("MapSecond failed: expected '42!', got %v", result.Second())
		}
	})

	t.Run("MapSecond with complex transformation", func(t *testing.T) {
		p := NewPair(true, []string{"a", "b", "c"})

		// Convert slice to its length
		lengthResult := MapSecond(p, func(s []string) int {
			return len(s)
		})

		if lengthResult.Second() != 3 {
			t.Errorf("MapSecond length conversion failed: expected 3, got %v", lengthResult.Second())
		}

		// Convert slice to joined string
		joinResult := MapSecond(p, func(s []string) string {
			result := ""
			for i, str := range s {
				if i > 0 {
					result += "-"
				}
				result += str
			}
			return result
		})

		if joinResult.Second() != "a-b-c" {
			t.Errorf("MapSecond join failed: expected 'a-b-c', got %v", joinResult.Second())
		}
	})
}

func TestPairSwap(t *testing.T) {
	t.Run("Basic swap", func(t *testing.T) {
		p := NewPair("first", 123)
		swapped := p.Swap()

		if swapped.First() != 123 {
			t.Errorf("Swap failed: expected 123 as first, got %v", swapped.First())
		}

		if swapped.Second() != "first" {
			t.Errorf("Swap failed: expected 'first' as second, got %v", swapped.Second())
		}
	})

	t.Run("Double swap should return original", func(t *testing.T) {
		original := NewPair(3.14, true)
		doubleSwapped := original.Swap().Swap()

		if doubleSwapped.First() != original.First() || doubleSwapped.Second() != original.Second() {
			t.Error("Double swap should return to original state")
		}
	})

	t.Run("Swap with same types", func(t *testing.T) {
		p := NewPair(10, 20)
		swapped := p.Swap()

		if swapped.First() != 20 || swapped.Second() != 10 {
			t.Errorf("Same type swap failed: expected (20, 10), got (%v, %v)",
				swapped.First(), swapped.Second())
		}
	})
}

func TestPairApply(t *testing.T) {
	t.Run("Apply for calculation", func(t *testing.T) {
		p := NewPair(15, 25)
		sum := Apply(p, func(a, b int) int {
			return a + b
		})

		if sum != 40 {
			t.Errorf("Apply sum failed: expected 40, got %v", sum)
		}

		product := Apply(p, func(a, b int) int {
			return a * b
		})

		if product != 375 {
			t.Errorf("Apply product failed: expected 375, got %v", product)
		}
	})

	t.Run("Apply for string operations", func(t *testing.T) {
		p := NewPair("Hello", "World")

		concatenated := Apply(p, func(a, b string) string {
			return a + " " + b
		})

		if concatenated != "Hello World" {
			t.Errorf("Apply concatenation failed: expected 'Hello World', got %v", concatenated)
		}

		comparison := Apply(p, func(a, b string) bool {
			return len(a) == len(b)
		})

		if !comparison {
			t.Error("Apply comparison failed: both strings have same length")
		}
	})

	t.Run("Apply with mixed types", func(t *testing.T) {
		p := NewPair("items", 5)
		description := Apply(p, func(name string, count int) string {
			return strconv.Itoa(count) + " " + name
		})

		if description != "5 items" {
			t.Errorf("Apply mixed types failed: expected '5 items', got %v", description)
		}
	})
}

func TestPairCurrying(t *testing.T) {
	t.Run("Basic currying", func(t *testing.T) {
		add := func(a, b int) int {
			return a + b
		}

		curriedAdd := Curry(add)
		add5 := curriedAdd(5)
		result := add5(10)

		if result != 15 {
			t.Errorf("Currying failed: expected 15, got %v", result)
		}
	})

	t.Run("Currying with different types", func(t *testing.T) {
		format := func(template string, value int) string {
			return template + ": " + strconv.Itoa(value)
		}

		curriedFormat := Curry(format)
		ageFormatter := curriedFormat("Age")
		result := ageFormatter(25)

		if result != "Age: 25" {
			t.Errorf("Type-mixed currying failed: expected 'Age: 25', got %v", result)
		}
	})

	t.Run("Partial application", func(t *testing.T) {
		multiply := func(a, b float64) float64 {
			return a * b
		}

		curriedMultiply := Curry(multiply)
		double := curriedMultiply(2.0)

		testValues := []float64{1.5, 3.0, 4.5}
		expected := []float64{3.0, 6.0, 9.0}

		for i, val := range testValues {
			result := double(val)
			if result != expected[i] {
				t.Errorf("Partial application failed for %v: expected %v, got %v",
					val, expected[i], result)
			}
		}
	})
}

func TestPairUncurry(t *testing.T) {
	t.Run("Basic uncurrying", func(t *testing.T) {
		curriedAdd := func(a int) func(int) int {
			return func(b int) int {
				return a + b
			}
		}

		uncurriedAdd := Uncurry(curriedAdd)
		p := NewPair(10, 20)
		result := uncurriedAdd(p)

		if result != 30 {
			t.Errorf("Uncurrying failed: expected 30, got %v", result)
		}
	})

	t.Run("Curry then uncurry", func(t *testing.T) {
		original := func(a, b string) string {
			return a + b
		}

		// Curry then uncurry should behave like original
		roundTrip := Uncurry(Curry(original))
		p := NewPair("Hello", "World")

		originalResult := original("Hello", "World")
		roundTripResult := roundTrip(p)

		if originalResult != roundTripResult {
			t.Errorf("Curry->Uncurry round trip failed: expected %v, got %v",
				originalResult, roundTripResult)
		}
	})
}

func TestPairEqual(t *testing.T) {
	t.Run("Equal pairs", func(t *testing.T) {
		p1 := NewPair(42, "test")
		p2 := NewPair(42, "test")

		intEq := func(a, b int) bool { return a == b }
		strEq := func(a, b string) bool { return a == b }

		if !p1.Equal(p2, intEq, strEq) {
			t.Error("Equal pairs should be equal")
		}
	})

	t.Run("Unequal pairs", func(t *testing.T) {
		p1 := NewPair(42, "test")
		p2 := NewPair(43, "test")
		p3 := NewPair(42, "different")

		intEq := func(a, b int) bool { return a == b }
		strEq := func(a, b string) bool { return a == b }

		if p1.Equal(p2, intEq, strEq) {
			t.Error("Pairs with different first values should not be equal")
		}

		if p1.Equal(p3, intEq, strEq) {
			t.Error("Pairs with different second values should not be equal")
		}
	})

	t.Run("Custom equality", func(t *testing.T) {
		p1 := NewPair([]int{1, 2, 3}, []string{"a", "b"})
		p2 := NewPair([]int{1, 2, 3}, []string{"a", "b"})

		sliceIntEq := func(a, b []int) bool {
			if len(a) != len(b) {
				return false
			}
			for i := range a {
				if a[i] != b[i] {
					return false
				}
			}
			return true
		}

		sliceStrEq := func(a, b []string) bool {
			if len(a) != len(b) {
				return false
			}
			for i := range a {
				if a[i] != b[i] {
					return false
				}
			}
			return true
		}

		if !p1.Equal(p2, sliceIntEq, sliceStrEq) {
			t.Error("Custom equality for slices failed")
		}
	})
}

func TestPairChaining(t *testing.T) {
	t.Run("Chain multiple operations", func(t *testing.T) {
		result := NewPair(2, 3).
			Map(func(a, b int) (int, int) { return a * 2, b * 3 }). // (4, 9)
			Swap().                                                 // (9, 4)
			Map(func(a, b int) (int, int) { return a + 1, b - 1 })  // (10, 3)

		if result.First() != 10 || result.Second() != 3 {
			t.Errorf("Chaining failed: expected (10, 3), got (%v, %v)",
				result.First(), result.Second())
		}
	})

	t.Run("Complex chaining with type changes", func(t *testing.T) {
		original := NewPair(5, "hello")

		// Convert first to string, then concatenate
		step1 := MapFirst(original, func(i int) string {
			return strconv.Itoa(i)
		})

		final := step1.Map(func(a, b string) (string, string) {
			return a + b, b + a
		})

		if final.First() != "5hello" || final.Second() != "hello5" {
			t.Errorf("Complex chaining failed: expected ('5hello', 'hello5'), got (%v, %v)",
				final.First(), final.Second())
		}
	})
}

func TestPairImmutability(t *testing.T) {
	t.Run("Original pair unchanged after operations", func(t *testing.T) {
		original := NewPair(10, "test")
		originalFirst := original.First()
		originalSecond := original.Second()

		// Perform various operations
		_ = original.Map(func(a int, b string) (int, string) { return a * 2, b + "!" })
		_ = MapFirst(original, func(a int) string { return "changed" })
		_ = original.Swap()
		_ = Apply(original, func(a int, b string) string { return b })

		// Original should be unchanged
		if original.First() != originalFirst || original.Second() != originalSecond {
			t.Error("Original pair was modified - immutability violated")
		}
	})
}
