package linked_test

import (
	"testing"

	"github.com/gosuda/stdx/listx"
	"github.com/gosuda/stdx/listx/linked"
)

// createLinkedDeque is a factory function for creating LinkedDeque instances
func createLinkedDeque[T any]() listx.Deque[T] {
	return linked.NewDeque[T]()
}

func TestLinkedDeque_AddFirst(t *testing.T) {
	testDequeAddFirst(t, createLinkedDeque[int])
}

func TestLinkedDeque_AddLast(t *testing.T) {
	testDequeAddLast(t, createLinkedDeque[int])
}

func TestLinkedDeque_RemoveFirst(t *testing.T) {
	testDequeRemoveFirst(t, createLinkedDeque[int])
}

func TestLinkedDeque_RemoveLast(t *testing.T) {
	testDequeRemoveLast(t, createLinkedDeque[int])
}

func TestLinkedDeque_PeekFirst(t *testing.T) {
	testDequePeekFirst(t, createLinkedDeque[int])
}

func TestLinkedDeque_PeekLast(t *testing.T) {
	testDequePeekLast(t, createLinkedDeque[int])
}

func TestLinkedDeque_ListMethods(t *testing.T) {
	testDequeListMethods(t, createLinkedDeque[int])
}

// Common test functions for Deque implementations

func testDequeAddFirst(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()

	d.AddFirst(1)
	d.AddFirst(2)
	d.AddFirst(3)

	if d.Size() != 3 {
		t.Errorf("Expected size 3, got %d", d.Size())
	}

	valOpt := d.Get(0)
	if valOpt.IsNone() || valOpt.Unwrap() != 3 {
		t.Errorf("Expected first element to be 3, got %v", valOpt)
	}

	valOpt = d.Get(2)
	if valOpt.IsNone() || valOpt.Unwrap() != 1 {
		t.Errorf("Expected last element to be 1, got %v", valOpt)
	}
}

func testDequeAddLast(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()

	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	if d.Size() != 3 {
		t.Errorf("Expected size 3, got %d", d.Size())
	}

	valOpt := d.Get(0)
	if valOpt.IsNone() || valOpt.Unwrap() != 1 {
		t.Errorf("Expected first element to be 1, got %v", valOpt)
	}

	valOpt = d.Get(2)
	if valOpt.IsNone() || valOpt.Unwrap() != 3 {
		t.Errorf("Expected last element to be 3, got %v", valOpt)
	}
}

func testDequeRemoveFirst(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	result := d.RemoveFirst()
	if result.IsErr() || result.Unwrap() != 1 {
		t.Errorf("Expected RemoveFirst to return 1, got %v", result)
	}

	if d.Size() != 2 {
		t.Errorf("Expected size 2 after removal, got %d", d.Size())
	}

	valOpt := d.Get(0)
	if valOpt.IsNone() || valOpt.Unwrap() != 2 {
		t.Errorf("Expected first element to be 2, got %v", valOpt)
	}

	// Test empty deque
	d.Clear()
	result = d.RemoveFirst()
	if result.IsOk() {
		t.Error("RemoveFirst on empty deque should return error")
	}
}

func testDequeRemoveLast(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	result := d.RemoveLast()
	if result.IsErr() || result.Unwrap() != 3 {
		t.Errorf("Expected RemoveLast to return 3, got %v", result)
	}

	if d.Size() != 2 {
		t.Errorf("Expected size 2 after removal, got %d", d.Size())
	}

	valOpt := d.Get(1)
	if valOpt.IsNone() || valOpt.Unwrap() != 2 {
		t.Errorf("Expected last element to be 2, got %v", valOpt)
	}

	// Test empty deque
	d.Clear()
	result = d.RemoveLast()
	if result.IsOk() {
		t.Error("RemoveLast on empty deque should return error")
	}
}

func testDequePeekFirst(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	valOpt := d.PeekFirst()
	if valOpt.IsNone() || valOpt.Unwrap() != 1 {
		t.Errorf("Expected PeekFirst to return 1, got %v", valOpt)
	}

	// Size should not change
	if d.Size() != 3 {
		t.Errorf("Expected size to remain 3, got %d", d.Size())
	}

	// Test empty deque
	d.Clear()
	valOpt = d.PeekFirst()
	if valOpt.IsSome() {
		t.Error("PeekFirst on empty deque should return None")
	}
}

func testDequePeekLast(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	valOpt := d.PeekLast()
	if valOpt.IsNone() || valOpt.Unwrap() != 3 {
		t.Errorf("Expected PeekLast to return 3, got %v", valOpt)
	}

	// Size should not change
	if d.Size() != 3 {
		t.Errorf("Expected size to remain 3, got %d", d.Size())
	}

	// Test empty deque
	d.Clear()
	valOpt = d.PeekLast()
	if valOpt.IsSome() {
		t.Error("PeekLast on empty deque should return None")
	}
}

func testDequeListMethods(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()

	// Test that Deque also supports List methods
	d.Add(1)
	d.Add(2)
	d.Add(3)

	if d.Size() != 3 {
		t.Errorf("Expected size 3, got %d", d.Size())
	}

	if !d.Contains(2) {
		t.Error("Deque should contain 2")
	}

	slice := d.ToSlice()
	expected := []int{1, 2, 3}
	for i, exp := range expected {
		if slice[i] != exp {
			t.Errorf("Expected element %d at index %d, got %d", exp, i, slice[i])
		}
	}
}
