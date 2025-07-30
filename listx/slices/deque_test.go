package slices_test

import (
	"testing"

	"github.com/gosuda/stdx/listx"
	"github.com/gosuda/stdx/listx/slices"
)

// createSliceDeque is a factory function for creating SliceDeque instances
func createSliceDeque[T any]() listx.Deque[T] {
	return slices.NewDeque[T]()
}

func TestSliceDeque_AddFirst(t *testing.T) {
	testDequeAddFirst(t, createSliceDeque[int])
}

func TestSliceDeque_AddLast(t *testing.T) {
	testDequeAddLast(t, createSliceDeque[int])
}

func TestSliceDeque_RemoveFirst(t *testing.T) {
	testDequeRemoveFirst(t, createSliceDeque[int])
}

func TestSliceDeque_RemoveLast(t *testing.T) {
	testDequeRemoveLast(t, createSliceDeque[int])
}

func TestSliceDeque_PeekFirst(t *testing.T) {
	testDequePeekFirst(t, createSliceDeque[int])
}

func TestSliceDeque_PeekLast(t *testing.T) {
	testDequePeekLast(t, createSliceDeque[int])
}

func TestSliceDeque_ListMethods(t *testing.T) {
	testDequeListMethods(t, createSliceDeque[int])
}

// Common test functions for Deque implementations (copied from linked package)

func testDequeAddFirst(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()

	d.AddFirst(1)
	d.AddFirst(2)
	d.AddFirst(3)

	if d.Size() != 3 {
		t.Errorf("Expected size 3, got %d", d.Size())
	}

	val, err := d.Get(0)
	if err != nil || val != 3 {
		t.Errorf("Expected first element to be 3, got %d", val)
	}

	val, err = d.Get(2)
	if err != nil || val != 1 {
		t.Errorf("Expected last element to be 1, got %d", val)
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

	val, err := d.Get(0)
	if err != nil || val != 1 {
		t.Errorf("Expected first element to be 1, got %d", val)
	}

	val, err = d.Get(2)
	if err != nil || val != 3 {
		t.Errorf("Expected last element to be 3, got %d", val)
	}
}

func testDequeRemoveFirst(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	val, err := d.RemoveFirst()
	if err != nil || val != 1 {
		t.Errorf("Expected RemoveFirst to return 1, got %d", val)
	}

	if d.Size() != 2 {
		t.Errorf("Expected size 2 after removal, got %d", d.Size())
	}

	val, err = d.Get(0)
	if err != nil || val != 2 {
		t.Errorf("Expected first element to be 2, got %d", val)
	}

	// Test empty deque
	d.Clear()
	_, err = d.RemoveFirst()
	if err == nil {
		t.Error("RemoveFirst on empty deque should return error")
	}
}

func testDequeRemoveLast(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	val, err := d.RemoveLast()
	if err != nil || val != 3 {
		t.Errorf("Expected RemoveLast to return 3, got %d", val)
	}

	if d.Size() != 2 {
		t.Errorf("Expected size 2 after removal, got %d", d.Size())
	}

	val, err = d.Get(1)
	if err != nil || val != 2 {
		t.Errorf("Expected last element to be 2, got %d", val)
	}

	// Test empty deque
	d.Clear()
	_, err = d.RemoveLast()
	if err == nil {
		t.Error("RemoveLast on empty deque should return error")
	}
}

func testDequePeekFirst(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	val, err := d.PeekFirst()
	if err != nil || val != 1 {
		t.Errorf("Expected PeekFirst to return 1, got %d", val)
	}

	// Size should not change
	if d.Size() != 3 {
		t.Errorf("Expected size to remain 3, got %d", d.Size())
	}

	// Test empty deque
	d.Clear()
	_, err = d.PeekFirst()
	if err == nil {
		t.Error("PeekFirst on empty deque should return error")
	}
}

func testDequePeekLast(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	val, err := d.PeekLast()
	if err != nil || val != 3 {
		t.Errorf("Expected PeekLast to return 3, got %d", val)
	}

	// Size should not change
	if d.Size() != 3 {
		t.Errorf("Expected size to remain 3, got %d", d.Size())
	}

	// Test empty deque
	d.Clear()
	_, err = d.PeekLast()
	if err == nil {
		t.Error("PeekLast on empty deque should return error")
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
