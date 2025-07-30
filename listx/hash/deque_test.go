package hash_test

import (
	"testing"

	"github.com/gosuda/stdx/listx"
	"github.com/gosuda/stdx/listx/hash"
)

// createHashDeque is a factory function for creating HashDeque instances
func createHashDeque[T any]() listx.Deque[T] {
	return hash.NewDeque[T]()
}

func TestHashDeque_AddFirst(t *testing.T) {
	testDequeAddFirst(t, createHashDeque[int])
}

func TestHashDeque_AddLast(t *testing.T) {
	testDequeAddLast(t, createHashDeque[int])
}

func TestHashDeque_RemoveFirst(t *testing.T) {
	testDequeRemoveFirst(t, createHashDeque[int])
}

func TestHashDeque_RemoveLast(t *testing.T) {
	testDequeRemoveLast(t, createHashDeque[int])
}

func TestHashDeque_PeekFirst(t *testing.T) {
	testDequePeekFirst(t, createHashDeque[int])
}

func TestHashDeque_PeekLast(t *testing.T) {
	testDequePeekLast(t, createHashDeque[int])
}

func TestHashDeque_Size(t *testing.T) {
	testDequeSize(t, createHashDeque[int])
}

func TestHashDeque_IsEmpty(t *testing.T) {
	testDequeIsEmpty(t, createHashDeque[int])
}

func TestHashDeque_Clear(t *testing.T) {
	testDequeClear(t, createHashDeque[int])
}

func TestHashDeque_ToSlice(t *testing.T) {
	testDequeToSlice(t, createHashDeque[int])
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

	firstOpt := d.PeekFirst()
	if firstOpt.IsNone() || firstOpt.Unwrap() != 3 {
		t.Errorf("Expected first element to be Some(3), got %v", firstOpt)
	}

	lastOpt := d.PeekLast()
	if lastOpt.IsNone() || lastOpt.Unwrap() != 1 {
		t.Errorf("Expected last element to be Some(1), got %v", lastOpt)
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

	firstOpt := d.PeekFirst()
	if firstOpt.IsNone() || firstOpt.Unwrap() != 1 {
		t.Errorf("Expected first element to be Some(1), got %v", firstOpt)
	}

	lastOpt := d.PeekLast()
	if lastOpt.IsNone() || lastOpt.Unwrap() != 3 {
		t.Errorf("Expected last element to be Some(3), got %v", lastOpt)
	}
}

func testDequeRemoveFirst(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	result := d.RemoveFirst()
	if result.IsErr() || result.Unwrap() != 1 {
		t.Errorf("Expected RemoveFirst to return Ok(1), got %v", result)
	}

	if d.Size() != 2 {
		t.Errorf("Expected size 2 after RemoveFirst, got %d", d.Size())
	}

	firstOpt := d.PeekFirst()
	if firstOpt.IsNone() || firstOpt.Unwrap() != 2 {
		t.Errorf("Expected first element to be Some(2), got %v", firstOpt)
	}

	// Test empty deque
	d.Clear()
	result = d.RemoveFirst()
	if result.IsOk() {
		t.Error("RemoveFirst on empty deque should return Err")
	}
}

func testDequeRemoveLast(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	result := d.RemoveLast()
	if result.IsErr() || result.Unwrap() != 3 {
		t.Errorf("Expected RemoveLast to return Ok(3), got %v", result)
	}

	if d.Size() != 2 {
		t.Errorf("Expected size 2 after RemoveLast, got %d", d.Size())
	}

	lastOpt := d.PeekLast()
	if lastOpt.IsNone() || lastOpt.Unwrap() != 2 {
		t.Errorf("Expected last element to be Some(2), got %v", lastOpt)
	}

	// Test empty deque
	d.Clear()
	result = d.RemoveLast()
	if result.IsOk() {
		t.Error("RemoveLast on empty deque should return Err")
	}
}

func testDequePeekFirst(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	valOpt := d.PeekFirst()
	if valOpt.IsNone() || valOpt.Unwrap() != 1 {
		t.Errorf("Expected PeekFirst to return Some(1), got %v", valOpt)
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
		t.Errorf("Expected PeekLast to return Some(3), got %v", valOpt)
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

func testDequeSize(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()

	if d.Size() != 0 {
		t.Errorf("Expected size 0 for empty deque, got %d", d.Size())
	}

	d.AddFirst(1)
	d.AddLast(2)

	if d.Size() != 2 {
		t.Errorf("Expected size 2, got %d", d.Size())
	}

	d.RemoveFirst()

	if d.Size() != 1 {
		t.Errorf("Expected size 1 after RemoveFirst, got %d", d.Size())
	}
}

func testDequeIsEmpty(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()

	if !d.IsEmpty() {
		t.Error("New deque should be empty")
	}

	d.AddFirst(1)

	if d.IsEmpty() {
		t.Error("Deque with elements should not be empty")
	}

	d.RemoveFirst()

	if !d.IsEmpty() {
		t.Error("Deque should be empty after removing all elements")
	}
}

func testDequeClear(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()
	d.AddFirst(1)
	d.AddLast(2)
	d.AddFirst(3)

	d.Clear()

	if !d.IsEmpty() {
		t.Error("Deque should be empty after Clear()")
	}

	if d.Size() != 0 {
		t.Errorf("Size should be 0 after Clear(), got %d", d.Size())
	}
}

func testDequeToSlice(t *testing.T, factory func() listx.Deque[int]) {
	d := factory()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	slice := d.ToSlice()

	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}

	// Deque ToSlice should return elements from first to last
	expected := []int{1, 2, 3}
	for i, exp := range expected {
		if slice[i] != exp {
			t.Errorf("Expected element %d at index %d, got %d", exp, i, slice[i])
		}
	}
}
