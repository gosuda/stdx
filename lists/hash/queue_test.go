package hash_test

import (
	"testing"

	"github.com/gosuda/stdx/lists"
	"github.com/gosuda/stdx/lists/hash"
)

// createHashQueue is a factory function for creating HashQueue instances
func createHashQueue[T any]() lists.Queue[T] {
	return hash.NewQueue[T]()
}

func TestHashQueue_Enqueue(t *testing.T) {
	testQueueEnqueue(t, createHashQueue[int])
}

func TestHashQueue_Dequeue(t *testing.T) {
	testQueueDequeue(t, createHashQueue[int])
}

func TestHashQueue_Peek(t *testing.T) {
	testQueuePeek(t, createHashQueue[int])
}

func TestHashQueue_Size(t *testing.T) {
	testQueueSize(t, createHashQueue[int])
}

func TestHashQueue_IsEmpty(t *testing.T) {
	testQueueIsEmpty(t, createHashQueue[int])
}

func TestHashQueue_Clear(t *testing.T) {
	testQueueClear(t, createHashQueue[int])
}

func TestHashQueue_ToSlice(t *testing.T) {
	testQueueToSlice(t, createHashQueue[int])
}

// Common test functions for Queue implementations (copied from linked package)

func testQueueEnqueue(t *testing.T, factory func() lists.Queue[int]) {
	q := factory()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}

	val, err := q.Peek()
	if err != nil || val != 1 {
		t.Errorf("Expected front element to be 1, got %d", val)
	}
}

func testQueueDequeue(t *testing.T, factory func() lists.Queue[int]) {
	q := factory()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	val, err := q.Dequeue()
	if err != nil || val != 1 {
		t.Errorf("Expected Dequeue to return 1, got %d", val)
	}

	if q.Size() != 2 {
		t.Errorf("Expected size 2 after dequeue, got %d", q.Size())
	}

	val, err = q.Dequeue()
	if err != nil || val != 2 {
		t.Errorf("Expected Dequeue to return 2, got %d", val)
	}

	val, err = q.Dequeue()
	if err != nil || val != 3 {
		t.Errorf("Expected Dequeue to return 3, got %d", val)
	}

	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all elements")
	}

	// Test empty queue
	_, err = q.Dequeue()
	if err == nil {
		t.Error("Dequeue on empty queue should return error")
	}
}

func testQueuePeek(t *testing.T, factory func() lists.Queue[int]) {
	q := factory()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	val, err := q.Peek()
	if err != nil || val != 1 {
		t.Errorf("Expected Peek to return 1, got %d", val)
	}

	// Size should not change
	if q.Size() != 3 {
		t.Errorf("Expected size to remain 3, got %d", q.Size())
	}

	// Test empty queue
	q.Clear()
	_, err = q.Peek()
	if err == nil {
		t.Error("Peek on empty queue should return error")
	}
}

func testQueueSize(t *testing.T, factory func() lists.Queue[int]) {
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

func testQueueIsEmpty(t *testing.T, factory func() lists.Queue[int]) {
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

func testQueueClear(t *testing.T, factory func() lists.Queue[int]) {
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

func testQueueToSlice(t *testing.T, factory func() lists.Queue[int]) {
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
