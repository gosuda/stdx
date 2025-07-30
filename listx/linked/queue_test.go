package linked_test

import (
	"testing"

	"github.com/gosuda/stdx/listx"
	"github.com/gosuda/stdx/listx/linked"
)

// createLinkedQueue is a factory function for creating LinkedQueue instances
func createLinkedQueue[T any]() listx.Queue[T] {
	return linked.NewQueue[T]()
}

func TestLinkedQueue_Enqueue(t *testing.T) {
	testQueueEnqueue(t, createLinkedQueue[int])
}

func TestLinkedQueue_Dequeue(t *testing.T) {
	testQueueDequeue(t, createLinkedQueue[int])
}

func TestLinkedQueue_Peek(t *testing.T) {
	testQueuePeek(t, createLinkedQueue[int])
}

func TestLinkedQueue_Size(t *testing.T) {
	testQueueSize(t, createLinkedQueue[int])
}

func TestLinkedQueue_IsEmpty(t *testing.T) {
	testQueueIsEmpty(t, createLinkedQueue[int])
}

func TestLinkedQueue_Clear(t *testing.T) {
	testQueueClear(t, createLinkedQueue[int])
}

func TestLinkedQueue_ToSlice(t *testing.T) {
	testQueueToSlice(t, createLinkedQueue[int])
}

// Common test functions for Queue implementations

func testQueueEnqueue(t *testing.T, factory func() listx.Queue[int]) {
	q := factory()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}

	valOpt := q.Peek()
	if valOpt.IsNone() || valOpt.Unwrap() != 1 {
		t.Errorf("Expected front element to be 1, got %v", valOpt)
	}
}

func testQueueDequeue(t *testing.T, factory func() listx.Queue[int]) {
	q := factory()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	result := q.Dequeue()
	if result.IsErr() || result.Unwrap() != 1 {
		t.Errorf("Expected Dequeue to return 1, got %v", result)
	}

	if q.Size() != 2 {
		t.Errorf("Expected size 2 after dequeue, got %d", q.Size())
	}

	result = q.Dequeue()
	if result.IsErr() || result.Unwrap() != 2 {
		t.Errorf("Expected Dequeue to return 2, got %v", result)
	}

	result = q.Dequeue()
	if result.IsErr() || result.Unwrap() != 3 {
		t.Errorf("Expected Dequeue to return 3, got %v", result)
	}

	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all elements")
	}

	// Test empty queue
	result = q.Dequeue()
	if result.IsOk() {
		t.Error("Dequeue on empty queue should return error")
	}
}

func testQueuePeek(t *testing.T, factory func() listx.Queue[int]) {
	q := factory()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	valOpt := q.Peek()
	if valOpt.IsNone() || valOpt.Unwrap() != 1 {
		t.Errorf("Expected Peek to return 1, got %v", valOpt)
	}

	// Size should not change
	if q.Size() != 3 {
		t.Errorf("Expected size to remain 3, got %d", q.Size())
	}

	// Test empty queue
	q.Clear()
	valOpt = q.Peek()
	if valOpt.IsSome() {
		t.Error("Peek on empty queue should return None")
	}
}

func testQueueSize(t *testing.T, factory func() listx.Queue[int]) {
	q := factory()

	if q.Size() != 0 {
		t.Errorf("Expected size 0 for empty queue, got %d", q.Size())
	}

	q.Enqueue(1)
	q.Enqueue(2)

	if q.Size() != 2 {
		t.Errorf("Expected size 2, got %d", q.Size())
	}

	q.Dequeue()

	if q.Size() != 1 {
		t.Errorf("Expected size 1 after dequeue, got %d", q.Size())
	}
}

func testQueueIsEmpty(t *testing.T, factory func() listx.Queue[int]) {
	q := factory()

	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}

	q.Enqueue(1)

	if q.IsEmpty() {
		t.Error("Queue with elements should not be empty")
	}

	q.Dequeue()

	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all elements")
	}
}

func testQueueClear(t *testing.T, factory func() listx.Queue[int]) {
	q := factory()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	q.Clear()

	if !q.IsEmpty() {
		t.Error("Queue should be empty after Clear()")
	}

	if q.Size() != 0 {
		t.Errorf("Size should be 0 after Clear(), got %d", q.Size())
	}
}

func testQueueToSlice(t *testing.T, factory func() listx.Queue[int]) {
	q := factory()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	slice := q.ToSlice()

	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}

	// Queue ToSlice should return elements from front to back
	expected := []int{1, 2, 3}
	for i, exp := range expected {
		if slice[i] != exp {
			t.Errorf("Expected element %d at index %d, got %d", exp, i, slice[i])
		}
	}
}
