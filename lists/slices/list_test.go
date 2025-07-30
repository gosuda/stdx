package slices_test

import (
	"testing"

	"github.com/gosuda/stdx/lists"
	"github.com/gosuda/stdx/lists/slices"
)

// createSliceList is a factory function for creating SliceList instances
func createSliceList[T any]() lists.List[T] {
	return slices.New[T]()
}

func TestSliceList_Add(t *testing.T) {
	testListAdd(t, createSliceList[int])
}

func TestSliceList_Insert(t *testing.T) {
	testListInsert(t, createSliceList[int])
}

func TestSliceList_Get(t *testing.T) {
	testListGet(t, createSliceList[int])
}

func TestSliceList_Set(t *testing.T) {
	testListSet(t, createSliceList[int])
}

func TestSliceList_Remove(t *testing.T) {
	testListRemove(t, createSliceList[int])
}

func TestSliceList_RemoveElement(t *testing.T) {
	testListRemoveElement(t, createSliceList[int])
}

func TestSliceList_IndexOf(t *testing.T) {
	testListIndexOf(t, createSliceList[int])
}

func TestSliceList_LastIndexOf(t *testing.T) {
	testListLastIndexOf(t, createSliceList[int])
}

func TestSliceList_Contains(t *testing.T) {
	testListContains(t, createSliceList[int])
}

func TestSliceList_Size(t *testing.T) {
	testListSize(t, createSliceList[int])
}

func TestSliceList_IsEmpty(t *testing.T) {
	testListIsEmpty(t, createSliceList[int])
}

func TestSliceList_Clear(t *testing.T) {
	testListClear(t, createSliceList[int])
}

func TestSliceList_ToSlice(t *testing.T) {
	testListToSlice(t, createSliceList[int])
}

func TestSliceList_ForEach(t *testing.T) {
	testListForEach(t, createSliceList[int])
}

// Common test functions (copied from linked package)

func testListAdd(t *testing.T, factory func() lists.List[int]) {
	l := factory()

	l.Add(1)
	l.Add(2)
	l.Add(3)

	if l.Size() != 3 {
		t.Errorf("Expected size 3, got %d", l.Size())
	}

	val, err := l.Get(0)
	if err != nil || val != 1 {
		t.Errorf("Expected first element to be 1, got %d", val)
	}

	val, err = l.Get(2)
	if err != nil || val != 3 {
		t.Errorf("Expected third element to be 3, got %d", val)
	}
}

func testListInsert(t *testing.T, factory func() lists.List[int]) {
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

func testListGet(t *testing.T, factory func() lists.List[int]) {
	l := factory()
	l.Add(10)
	l.Add(20)
	l.Add(30)

	val, err := l.Get(1)
	if err != nil || val != 20 {
		t.Errorf("Expected Get(1) to return 20, got %d", val)
	}

	// Test out of bounds
	_, err = l.Get(-1)
	if err == nil {
		t.Error("Get with negative index should fail")
	}

	_, err = l.Get(3)
	if err == nil {
		t.Error("Get with too large index should fail")
	}
}

func testListSet(t *testing.T, factory func() lists.List[int]) {
	l := factory()
	l.Add(10)
	l.Add(20)
	l.Add(30)

	err := l.Set(1, 99)
	if err != nil {
		t.Errorf("Set failed: %v", err)
	}

	val, _ := l.Get(1)
	if val != 99 {
		t.Errorf("Expected Set to change value to 99, got %d", val)
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

func testListRemove(t *testing.T, factory func() lists.List[int]) {
	l := factory()
	l.Add(10)
	l.Add(20)
	l.Add(30)

	// Remove from middle
	val, err := l.Remove(1)
	if err != nil || val != 20 {
		t.Errorf("Expected Remove(1) to return 20, got %d", val)
	}

	if l.Size() != 2 {
		t.Errorf("Expected size 2 after removal, got %d", l.Size())
	}

	// Verify remaining elements
	val, _ = l.Get(0)
	if val != 10 {
		t.Errorf("Expected first element to be 10, got %d", val)
	}

	val, _ = l.Get(1)
	if val != 30 {
		t.Errorf("Expected second element to be 30, got %d", val)
	}

	// Test out of bounds
	_, err = l.Remove(-1)
	if err == nil {
		t.Error("Remove with negative index should fail")
	}

	_, err = l.Remove(2)
	if err == nil {
		t.Error("Remove with too large index should fail")
	}
}

func testListRemoveElement(t *testing.T, factory func() lists.List[int]) {
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
	val, _ := l.Get(0)
	if val != 20 {
		t.Errorf("Expected first element to be 20, got %d", val)
	}

	// Remove non-existing element
	removed = l.RemoveElement(99)
	if removed {
		t.Error("RemoveElement should return false for non-existing element")
	}
}

func testListIndexOf(t *testing.T, factory func() lists.List[int]) {
	l := factory()
	l.Add(10)
	l.Add(20)
	l.Add(10)

	index := l.IndexOf(10)
	if index != 0 {
		t.Errorf("Expected IndexOf(10) to return 0, got %d", index)
	}

	index = l.IndexOf(20)
	if index != 1 {
		t.Errorf("Expected IndexOf(20) to return 1, got %d", index)
	}

	index = l.IndexOf(99)
	if index != -1 {
		t.Errorf("Expected IndexOf(99) to return -1, got %d", index)
	}
}

func testListLastIndexOf(t *testing.T, factory func() lists.List[int]) {
	l := factory()
	l.Add(10)
	l.Add(20)
	l.Add(10)

	index := l.LastIndexOf(10)
	if index != 2 {
		t.Errorf("Expected LastIndexOf(10) to return 2, got %d", index)
	}

	index = l.LastIndexOf(20)
	if index != 1 {
		t.Errorf("Expected LastIndexOf(20) to return 1, got %d", index)
	}

	index = l.LastIndexOf(99)
	if index != -1 {
		t.Errorf("Expected LastIndexOf(99) to return -1, got %d", index)
	}
}

func testListContains(t *testing.T, factory func() lists.List[int]) {
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

func testListSize(t *testing.T, factory func() lists.List[int]) {
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

func testListIsEmpty(t *testing.T, factory func() lists.List[int]) {
	l := factory()

	if !l.IsEmpty() {
		t.Error("New list should be empty")
	}

	l.Add(1)

	if l.IsEmpty() {
		t.Error("List with elements should not be empty")
	}
}

func testListClear(t *testing.T, factory func() lists.List[int]) {
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

func testListToSlice(t *testing.T, factory func() lists.List[int]) {
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

func testListForEach(t *testing.T, factory func() lists.List[int]) {
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
