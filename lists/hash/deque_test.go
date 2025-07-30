package hash_test

import (
	"testing"

	"github.com/gosuda/stdx/lists"
	"github.com/gosuda/stdx/lists/hash"
)

// createHashDeque is a factory function for creating HashDeque instances
func createHashDeque[T any]() lists.Deque[T] {
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

func testDequeAddFirst(t *testing.T, factory func() lists.Deque[int]) {
	d := factory()

	d.AddFirst(1)
	d.AddFirst(2)
	d.AddFirst(3)

	if d.Size() != 3 {
		t.Errorf("Expected size 3, got %d", d.Size())
	}

	first, err := d.PeekFirst()
	if err != nil || first != 3 {
		t.Errorf("Expected first element to be 3, got %d", first)
	}

	last, err := d.PeekLast()
	if err != nil || last != 1 {
		t.Errorf("Expected last element to be 1, got %d", last)
	}
}

func testDequeAddLast(t *testing.T, factory func() lists.Deque[int]) {
	d := factory()

	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	if d.Size() != 3 {
		t.Errorf("Expected size 3, got %d", d.Size())
	}

	first, err := d.PeekFirst()
	if err != nil || first != 1 {
		t.Errorf("Expected first element to be 1, got %d", first)
	}

	last, err := d.PeekLast()
	if err != nil || last != 3 {
		t.Errorf("Expected last element to be 3, got %d", last)
	}
}

func testDequeRemoveFirst(t *testing.T, factory func() lists.Deque[int]) {
	d := factory()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	val, err := d.RemoveFirst()
	if err != nil || val != 1 {
		t.Errorf("Expected RemoveFirst to return 1, got %d", val)
	}

	if d.Size() != 2 {
		t.Errorf("Expected size 2 after RemoveFirst, got %d", d.Size())
	}

	first, err := d.PeekFirst()
	if err != nil || first != 2 {
		t.Errorf("Expected first element to be 2, got %d", first)
	}

	// Test empty deque
	d.Clear()
	_, err = d.RemoveFirst()
	if err == nil {
		t.Error("RemoveFirst on empty deque should return error")
	}
}

func testDequeRemoveLast(t *testing.T, factory func() lists.Deque[int]) {
	d := factory()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)

	val, err := d.RemoveLast()
	if err != nil || val != 3 {
		t.Errorf("Expected RemoveLast to return 3, got %d", val)
	}

	if d.Size() != 2 {
		t.Errorf("Expected size 2 after RemoveLast, got %d", d.Size())
	}

	last, err := d.PeekLast()
	if err != nil || last != 2 {
		t.Errorf("Expected last element to be 2, got %d", last)
	}

	// Test empty deque
	d.Clear()
	_, err = d.RemoveLast()
	if err == nil {
		t.Error("RemoveLast on empty deque should return error")
	}
}

func testDequePeekFirst(t *testing.T, factory func() lists.Deque[int]) {
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

func testDequePeekLast(t *testing.T, factory func() lists.Deque[int]) {
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

func testDequeSize(t *testing.T, factory func() lists.Deque[int]) {
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

func testDequeIsEmpty(t *testing.T, factory func() lists.Deque[int]) {
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

func testDequeClear(t *testing.T, factory func() lists.Deque[int]) {
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

func testDequeToSlice(t *testing.T, factory func() lists.Deque[int]) {
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
