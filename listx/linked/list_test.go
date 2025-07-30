package linked_test

import (
	"testing"

	"github.com/gosuda/stdx/listx"
	"github.com/gosuda/stdx/listx/linked"
)

// createLinkedList is a factory function for creating LinkedList instances
func createLinkedList[T any]() listx.List[T] {
	return linked.New[T]()
}

func TestLinkedList_Add(t *testing.T) {
	testListAdd(t, createLinkedList[int])
}

func TestLinkedList_Insert(t *testing.T) {
	testListInsert(t, createLinkedList[int])
}

func TestLinkedList_Get(t *testing.T) {
	testListGet(t, createLinkedList[int])
}

func TestLinkedList_Set(t *testing.T) {
	testListSet(t, createLinkedList[int])
}

func TestLinkedList_Remove(t *testing.T) {
	testListRemove(t, createLinkedList[int])
}

func TestLinkedList_RemoveElement(t *testing.T) {
	testListRemoveElement(t, createLinkedList[int])
}

func TestLinkedList_IndexOf(t *testing.T) {
	testListIndexOf(t, createLinkedList[int])
}

func TestLinkedList_LastIndexOf(t *testing.T) {
	testListLastIndexOf(t, createLinkedList[int])
}

func TestLinkedList_Contains(t *testing.T) {
	testListContains(t, createLinkedList[int])
}

func TestLinkedList_Size(t *testing.T) {
	testListSize(t, createLinkedList[int])
}

func TestLinkedList_IsEmpty(t *testing.T) {
	testListIsEmpty(t, createLinkedList[int])
}

func TestLinkedList_Clear(t *testing.T) {
	testListClear(t, createLinkedList[int])
}

func TestLinkedList_ToSlice(t *testing.T) {
	testListToSlice(t, createLinkedList[int])
}

func TestLinkedList_ForEach(t *testing.T) {
	testListForEach(t, createLinkedList[int])
}

// Common test functions that can be reused for any List implementation

func testListAdd(t *testing.T, factory func() listx.List[int]) {
	l := factory()

	l.Add(1)
	l.Add(2)
	l.Add(3)

	if l.Size() != 3 {
		t.Errorf("Expected size 3, got %d", l.Size())
	}

	valOpt := l.Get(0)
	if valOpt.IsNone() || valOpt.Unwrap() != 1 {
		t.Errorf("Expected first element to be 1, got %v", valOpt)
	}

	valOpt = l.Get(2)
	if valOpt.IsNone() || valOpt.Unwrap() != 3 {
		t.Errorf("Expected third element to be 3, got %v", valOpt)
	}
}

func testListInsert(t *testing.T, factory func() listx.List[int]) {
	l := factory()

	// Insert into empty list
	err := l.Insert(0, 1)
	if err != nil {
		t.Errorf("Insert into empty list failed: %v", err)
	}

	// Insert at beginning
	err = l.Insert(0, 0)
	if err != nil {
		t.Errorf("Insert at beginning failed: %v", err)
	}

	// Insert at end
	err = l.Insert(2, 2)
	if err != nil {
		t.Errorf("Insert at end failed: %v", err)
	}

	// Insert in middle
	err = l.Insert(2, 99)
	if err != nil {
		t.Errorf("Insert in middle failed: %v", err)
	}

	expected := []int{0, 1, 99, 2}
	slice := l.ToSlice()
	for i, exp := range expected {
		if slice[i] != exp {
			t.Errorf("Expected element %d at index %d, got %d", exp, i, slice[i])
		}
	}

	// Test out of bounds
	err = l.Insert(-1, 100)
	if err == nil {
		t.Error("Insert with negative index should fail")
	}

	err = l.Insert(10, 100)
	if err == nil {
		t.Error("Insert with too large index should fail")
	}
}

func testListGet(t *testing.T, factory func() listx.List[int]) {
	l := factory()
	l.Add(10)
	l.Add(20)
	l.Add(30)

	valOpt := l.Get(1)
	if valOpt.IsNone() || valOpt.Unwrap() != 20 {
		t.Errorf("Expected Get(1) to return 20, got %v", valOpt)
	}

	// Test out of bounds
	valOpt = l.Get(-1)
	if valOpt.IsSome() {
		t.Error("Get with negative index should return None")
	}

	valOpt = l.Get(3)
	if valOpt.IsSome() {
		t.Error("Get with too large index should return None")
	}
}

func testListSet(t *testing.T, factory func() listx.List[int]) {
	l := factory()
	l.Add(10)
	l.Add(20)
	l.Add(30)

	err := l.Set(1, 99)
	if err != nil {
		t.Errorf("Set failed: %v", err)
	}

	valOpt := l.Get(1)
	if valOpt.IsNone() || valOpt.Unwrap() != 99 {
		t.Errorf("Expected Set to change value to 99, got %v", valOpt)
	}

	// Test out of bounds
	err = l.Set(-1, 100)
	if err == nil {
		t.Error("Set with negative index should fail")
	}

	err = l.Set(3, 100)
	if err == nil {
		t.Error("Set with too large index should fail")
	}
}

func testListRemove(t *testing.T, factory func() listx.List[int]) {
	l := factory()
	l.Add(10)
	l.Add(20)
	l.Add(30)

	// Remove from middle
	result := l.Remove(1)
	if result.IsErr() || result.Unwrap() != 20 {
		t.Errorf("Expected Remove(1) to return 20, got %v", result)
	}

	if l.Size() != 2 {
		t.Errorf("Expected size 2 after removal, got %d", l.Size())
	}

	// Verify remaining elements
	valOpt := l.Get(0)
	if valOpt.IsNone() || valOpt.Unwrap() != 10 {
		t.Errorf("Expected first element to be 10, got %v", valOpt)
	}

	valOpt = l.Get(1)
	if valOpt.IsNone() || valOpt.Unwrap() != 30 {
		t.Errorf("Expected second element to be 30, got %v", valOpt)
	}

	// Test out of bounds
	result = l.Remove(-1)
	if result.IsOk() {
		t.Error("Remove with negative index should return error")
	}

	result = l.Remove(2)
	if result.IsOk() {
		t.Error("Remove with too large index should return error")
	}
}

func testListRemoveElement(t *testing.T, factory func() listx.List[int]) {
	l := factory()
	l.Add(10)
	l.Add(20)
	l.Add(10)

	// Remove existing element
	removed := l.RemoveElement(10)
	if !removed {
		t.Error("RemoveElement should return true for existing element")
	}

	if l.Size() != 2 {
		t.Errorf("Expected size 2 after removal, got %d", l.Size())
	}

	// Verify first occurrence was removed
	valOpt := l.Get(0)
	if valOpt.IsNone() || valOpt.Unwrap() != 20 {
		t.Errorf("Expected first element to be 20, got %v", valOpt)
	}

	// Remove non-existing element
	removed = l.RemoveElement(99)
	if removed {
		t.Error("RemoveElement should return false for non-existing element")
	}
}

func testListIndexOf(t *testing.T, factory func() listx.List[int]) {
	l := factory()
	l.Add(10)
	l.Add(20)
	l.Add(10)

	indexOpt := l.IndexOf(10)
	if indexOpt.IsNone() || indexOpt.Unwrap() != 0 {
		t.Errorf("Expected IndexOf(10) to return 0, got %v", indexOpt)
	}

	indexOpt = l.IndexOf(20)
	if indexOpt.IsNone() || indexOpt.Unwrap() != 1 {
		t.Errorf("Expected IndexOf(20) to return 1, got %v", indexOpt)
	}

	indexOpt = l.IndexOf(99)
	if indexOpt.IsSome() {
		t.Error("IndexOf non-existing element should return None")
	}
}

func testListLastIndexOf(t *testing.T, factory func() listx.List[int]) {
	l := factory()
	l.Add(10)
	l.Add(20)
	l.Add(10)

	indexOpt := l.LastIndexOf(10)
	if indexOpt.IsNone() || indexOpt.Unwrap() != 2 {
		t.Errorf("Expected LastIndexOf(10) to return 2, got %v", indexOpt)
	}

	indexOpt = l.LastIndexOf(20)
	if indexOpt.IsNone() || indexOpt.Unwrap() != 1 {
		t.Errorf("Expected LastIndexOf(20) to return 1, got %v", indexOpt)
	}

	indexOpt = l.LastIndexOf(99)
	if indexOpt.IsSome() {
		t.Error("LastIndexOf non-existing element should return None")
	}
}

func testListContains(t *testing.T, factory func() listx.List[int]) {
	l := factory()
	l.Add(10)
	l.Add(20)
	l.Add(30)

	if !l.Contains(20) {
		t.Error("List should contain 20")
	}

	if l.Contains(99) {
		t.Error("List should not contain 99")
	}
}

func testListSize(t *testing.T, factory func() listx.List[int]) {
	l := factory()

	if l.Size() != 0 {
		t.Errorf("Expected size 0 for empty list, got %d", l.Size())
	}

	l.Add(1)
	l.Add(2)

	if l.Size() != 2 {
		t.Errorf("Expected size 2, got %d", l.Size())
	}
}

func testListIsEmpty(t *testing.T, factory func() listx.List[int]) {
	l := factory()

	if !l.IsEmpty() {
		t.Error("New list should be empty")
	}

	l.Add(1)

	if l.IsEmpty() {
		t.Error("List with elements should not be empty")
	}
}

func testListClear(t *testing.T, factory func() listx.List[int]) {
	l := factory()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	l.Clear()

	if !l.IsEmpty() {
		t.Error("List should be empty after Clear()")
	}

	if l.Size() != 0 {
		t.Errorf("Size should be 0 after Clear(), got %d", l.Size())
	}
}

func testListToSlice(t *testing.T, factory func() listx.List[int]) {
	l := factory()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	slice := l.ToSlice()

	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}

	expected := []int{1, 2, 3}
	for i, exp := range expected {
		if slice[i] != exp {
			t.Errorf("Expected element %d at index %d, got %d", exp, i, slice[i])
		}
	}
}

func testListForEach(t *testing.T, factory func() listx.List[int]) {
	l := factory()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	sum := 0
	l.ForEach(func(element int) {
		sum += element
	})

	if sum != 6 {
		t.Errorf("Expected sum 6, got %d", sum)
	}
}
