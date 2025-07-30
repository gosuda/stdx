package tuple

import (
	"testing"
)

func TestPair(t *testing.T) {
	// Create a pair
	p := NewPair(42, "hello")

	// Test getters
	if p.First() != 42 {
		t.Errorf("Expected First() to be 42, got %v", p.First())
	}

	if p.Second() != "hello" {
		t.Errorf("Expected Second() to be 'hello', got %v", p.Second())
	}

	// Test Map
	p2 := p.Map(func(i int, s string) (int, string) {
		return i * 2, s + " world"
	})

	if p2.First() != 84 {
		t.Errorf("Expected mapped First() to be 84, got %v", p2.First())
	}

	if p2.Second() != "hello world" {
		t.Errorf("Expected mapped Second() to be 'hello world', got %v", p2.Second())
	}

	// Test MapFirst
	p3 := MapFirst(p, func(i int) string {
		return "number"
	})

	if p3.First() != "number" {
		t.Errorf("Expected MapFirst result to be 'number', got %v", p3.First())
	}

	// Test Swap
	swapped := p.Swap()
	if swapped.First() != "hello" || swapped.Second() != 42 {
		t.Errorf("Swap failed: got (%v, %v)", swapped.First(), swapped.Second())
	}

	// Test Apply
	result := Apply(p, func(i int, s string) string {
		return s + " " + string(rune(i))
	})

	if result == "" {
		t.Error("Apply should return a result")
	}
}

func TestTriple(t *testing.T) {
	// Create a triple
	tr := NewTriple(1, "test", true)

	// Test getters
	if tr.First() != 1 {
		t.Errorf("Expected First() to be 1, got %v", tr.First())
	}

	if tr.Second() != "test" {
		t.Errorf("Expected Second() to be 'test', got %v", tr.Second())
	}

	if tr.Third() != true {
		t.Errorf("Expected Third() to be true, got %v", tr.Third())
	}

	// Test Map
	tr2 := tr.Map(func(i int, s string, b bool) (int, string, bool) {
		return i + 1, s + "!", !b
	})

	if tr2.First() != 2 || tr2.Second() != "test!" || tr2.Third() != false {
		t.Errorf("Map failed: got (%v, %v, %v)", tr2.First(), tr2.Second(), tr2.Third())
	}

	// Test MapFirstTriple
	tr3 := MapFirstTriple(tr, func(i int) string {
		return "changed"
	})

	if tr3.First() != "changed" {
		t.Errorf("Expected MapFirstTriple result to be 'changed', got %v", tr3.First())
	}

	// Test rotation
	rotated := tr.RotateLeft()
	if rotated.First() != "test" || rotated.Second() != true || rotated.Third() != 1 {
		t.Errorf("RotateLeft failed: got (%v, %v, %v)", rotated.First(), rotated.Second(), rotated.Third())
	}

	// Test ToPair
	pair := tr.ToPair()
	if pair.First() != 1 || pair.Second() != "test" {
		t.Errorf("ToPair failed: got (%v, %v)", pair.First(), pair.Second())
	}

	// Test FromPair
	basePair := NewPair(10, "base")
	tripleFromPair := FromPair(basePair, 3.14)
	if tripleFromPair.First() != 10 || tripleFromPair.Second() != "base" || tripleFromPair.Third() != 3.14 {
		t.Errorf("FromPair failed: got (%v, %v, %v)", tripleFromPair.First(), tripleFromPair.Second(), tripleFromPair.Third())
	}
}

func TestCurrying(t *testing.T) {
	// Test currying for Pair
	add := func(a, b int) int {
		return a + b
	}

	curriedAdd := Curry(add)
	addFive := curriedAdd(5)
	result := addFive(3)

	if result != 8 {
		t.Errorf("Expected curried function to return 8, got %v", result)
	}

	// Test uncurrying
	uncurriedAdd := Uncurry(curriedAdd)
	pair := NewPair(10, 20)
	result2 := uncurriedAdd(pair)

	if result2 != 30 {
		t.Errorf("Expected uncurried function to return 30, got %v", result2)
	}
}

func TestTripleCurrying(t *testing.T) {
	// Test currying for Triple
	multiply := func(a, b, c int) int {
		return a * b * c
	}

	curriedMultiply := CurryTriple(multiply)
	result := curriedMultiply(2)(3)(4)

	if result != 24 {
		t.Errorf("Expected curried triple function to return 24, got %v", result)
	}

	// Test uncurrying
	uncurriedMultiply := UncurryTriple(curriedMultiply)
	triple := NewTriple(2, 3, 4)
	result2 := uncurriedMultiply(triple)

	if result2 != 24 {
		t.Errorf("Expected uncurried triple function to return 24, got %v", result2)
	}
}
