package slices_test

import (
	"testing"

	"github.com/gosuda/stdx/listx"
	"github.com/gosuda/stdx/listx/slices"
)

// createSlicesStack is a factory function for creating LinkedStack instances
func createSlicesStack[T any]() listx.Stack[T] {
	return slices.NewStack[T]()
}

func TestSlicesStack_Push(t *testing.T) {
	testStackPush(t, createSlicesStack[int])
}

func TestSlicesStack_Pop(t *testing.T) {
	testStackPop(t, createSlicesStack[int])
}

func TestSlicesStack_Peek(t *testing.T) {
	testStackPeek(t, createSlicesStack[int])
}

func TestSlicesStack_Size(t *testing.T) {
	testStackSize(t, createSlicesStack[int])
}

func TestSlicesStack_IsEmpty(t *testing.T) {
	testStackIsEmpty(t, createSlicesStack[int])
}

func TestSlicesStack_Clear(t *testing.T) {
	testStackClear(t, createSlicesStack[int])
}

func TestSlicesStack_ToSlice(t *testing.T) {
	testStackToSlice(t, createSlicesStack[int])
}

// Common test functions for Stack implementations

func testStackPush(t *testing.T, factory func() listx.Stack[int]) {
	s := factory()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}

	valOpt := s.Peek()
	if valOpt.IsNone() || valOpt.Unwrap() != 3 {
		t.Errorf("Expected top element to be 3, got %v", valOpt)
	}
}

func testStackPop(t *testing.T, factory func() listx.Stack[int]) {
	s := factory()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	result := s.Pop()
	if result.IsErr() || result.Unwrap() != 3 {
		t.Errorf("Expected Pop to return 3, got %v", result)
	}

	if s.Size() != 2 {
		t.Errorf("Expected size 2 after pop, got %d", s.Size())
	}

	result = s.Pop()
	if result.IsErr() || result.Unwrap() != 2 {
		t.Errorf("Expected Pop to return 2, got %v", result)
	}

	result = s.Pop()
	if result.IsErr() || result.Unwrap() != 1 {
		t.Errorf("Expected Pop to return 1, got %v", result)
	}

	if !s.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}

	// Test empty stack
	result = s.Pop()
	if result.IsOk() {
		t.Error("Pop on empty stack should return error")
	}
}

func testStackPeek(t *testing.T, factory func() listx.Stack[int]) {
	s := factory()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	valOpt := s.Peek()
	if valOpt.IsNone() || valOpt.Unwrap() != 3 {
		t.Errorf("Expected Peek to return 3, got %v", valOpt)
	}

	// Size should not change
	if s.Size() != 3 {
		t.Errorf("Expected size to remain 3, got %d", s.Size())
	}

	// Test empty stack
	s.Clear()
	valOpt = s.Peek()
	if valOpt.IsSome() {
		t.Error("Peek on empty stack should return None")
	}
}

func testStackSize(t *testing.T, factory func() listx.Stack[int]) {
	s := factory()

	if s.Size() != 0 {
		t.Errorf("Expected size 0 for empty stack, got %d", s.Size())
	}

	s.Push(1)
	s.Push(2)

	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}

	s.Pop()

	if s.Size() != 1 {
		t.Errorf("Expected size 1 after pop, got %d", s.Size())
	}
}

func testStackIsEmpty(t *testing.T, factory func() listx.Stack[int]) {
	s := factory()

	if !s.IsEmpty() {
		t.Error("New stack should be empty")
	}

	s.Push(1)

	if s.IsEmpty() {
		t.Error("Stack with elements should not be empty")
	}

	s.Pop()

	if !s.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}
}

func testStackClear(t *testing.T, factory func() listx.Stack[int]) {
	s := factory()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	s.Clear()

	if !s.IsEmpty() {
		t.Error("Stack should be empty after Clear()")
	}

	if s.Size() != 0 {
		t.Errorf("Size should be 0 after Clear(), got %d", s.Size())
	}
}

func testStackToSlice(t *testing.T, factory func() listx.Stack[int]) {
	s := factory()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	slice := s.ToSlice()

	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}

	// Stack ToSlice should return elements from top to bottom
	expected := []int{3, 2, 1}
	for i, exp := range expected {
		if slice[i] != exp {
			t.Errorf("Expected element %d at index %d, got %d", exp, i, slice[i])
		}
	}
}
