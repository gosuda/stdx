package hashset_test

import (
	"testing"

	"github.com/gosuda/stdx/setx"
	"github.com/gosuda/stdx/setx/hashset"
)

// createHashSet is a factory function for creating HashSet instances
func createHashSet[T comparable]() setx.Set[T] {
	return hashset.New[T]()
}

func TestHashSet_Add(t *testing.T) {
	testSetAdd(t, createHashSet[int])
}

func TestHashSet_Remove(t *testing.T) {
	testSetRemove(t, createHashSet[int])
}

func TestHashSet_Contains(t *testing.T) {
	testSetContains(t, createHashSet[int])
}

func TestHashSet_Size(t *testing.T) {
	testSetSize(t, createHashSet[int])
}

func TestHashSet_IsEmpty(t *testing.T) {
	testSetIsEmpty(t, createHashSet[int])
}

func TestHashSet_Clear(t *testing.T) {
	testSetClear(t, createHashSet[int])
}

func TestHashSet_ToSlice(t *testing.T) {
	testSetToSlice(t, createHashSet[int])
}

func TestHashSet_ForEach(t *testing.T) {
	testSetForEach(t, createHashSet[int])
}

func TestHashSet_Union(t *testing.T) {
	testSetUnion(t, createHashSet[int])
}

func TestHashSet_Intersection(t *testing.T) {
	testSetIntersection(t, createHashSet[int])
}

func TestHashSet_Difference(t *testing.T) {
	testSetDifference(t, createHashSet[int])
}

func TestHashSet_IsSubsetOf(t *testing.T) {
	testSetIsSubsetOf(t, createHashSet[int])
}

func TestHashSet_IsSupersetOf(t *testing.T) {
	testSetIsSupersetOf(t, createHashSet[int])
}

func TestHashSet_Find(t *testing.T) {
	testSetFind(t, createHashSet[int])
}

func TestHashSet_GetAny(t *testing.T) {
	testSetGetAny(t, createHashSet[int])
}

func TestHashSet_TryRemove(t *testing.T) {
	testSetTryRemove(t, createHashSet[int])
}

func TestHashSet_Filter(t *testing.T) {
	testSetFilter(t, createHashSet[int])
}

// Common test functions that can be reused for any Set implementation

func testSetAdd(t *testing.T, factory func() setx.Set[int]) {
	set := factory()

	// Test adding new element
	if !set.Add(1) {
		t.Error("Add(1) should return true for new element")
	}
	if !set.Contains(1) {
		t.Error("Set should contain 1 after adding")
	}

	// Test adding duplicate element
	if set.Add(1) {
		t.Error("Add(1) should return false for duplicate element")
	}

	// Test size after additions
	set.Add(2)
	set.Add(3)
	if set.Size() != 3 {
		t.Errorf("Expected size 3, got %d", set.Size())
	}
}

func testSetRemove(t *testing.T, factory func() setx.Set[int]) {
	set := factory()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	// Test removing existing element
	if !set.Remove(2) {
		t.Error("Remove(2) should return true for existing element")
	}
	if set.Contains(2) {
		t.Error("Set should not contain 2 after removal")
	}

	// Test removing non-existing element
	if set.Remove(4) {
		t.Error("Remove(4) should return false for non-existing element")
	}

	// Test size after removal
	if set.Size() != 2 {
		t.Errorf("Expected size 2, got %d", set.Size())
	}
}

func testSetContains(t *testing.T, factory func() setx.Set[int]) {
	set := factory()
	set.Add(1)
	set.Add(2)

	if !set.Contains(1) {
		t.Error("Set should contain 1")
	}
	if !set.Contains(2) {
		t.Error("Set should contain 2")
	}
	if set.Contains(3) {
		t.Error("Set should not contain 3")
	}
}

func testSetSize(t *testing.T, factory func() setx.Set[int]) {
	set := factory()

	if set.Size() != 0 {
		t.Errorf("Empty set size should be 0, got %d", set.Size())
	}

	set.Add(1)
	if set.Size() != 1 {
		t.Errorf("Size should be 1, got %d", set.Size())
	}

	set.Add(2)
	set.Add(3)
	if set.Size() != 3 {
		t.Errorf("Size should be 3, got %d", set.Size())
	}

	set.Remove(2)
	if set.Size() != 2 {
		t.Errorf("Size should be 2 after removal, got %d", set.Size())
	}
}

func testSetIsEmpty(t *testing.T, factory func() setx.Set[int]) {
	set := factory()

	if !set.IsEmpty() {
		t.Error("New set should be empty")
	}

	set.Add(1)
	if set.IsEmpty() {
		t.Error("Set with elements should not be empty")
	}

	set.Remove(1)
	if !set.IsEmpty() {
		t.Error("Set should be empty after removing all elements")
	}
}

func testSetClear(t *testing.T, factory func() setx.Set[int]) {
	set := factory()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	set.Clear()

	if !set.IsEmpty() {
		t.Error("Set should be empty after Clear()")
	}
	if set.Size() != 0 {
		t.Errorf("Size should be 0 after Clear(), got %d", set.Size())
	}
	if set.Contains(1) || set.Contains(2) || set.Contains(3) {
		t.Error("Set should not contain any elements after Clear()")
	}
}

func testSetToSlice(t *testing.T, factory func() setx.Set[int]) {
	set := factory()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	slice := set.ToSlice()

	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}

	// Check all elements are present (order doesn't matter)
	found := make(map[int]bool)
	for _, v := range slice {
		found[v] = true
	}

	if !found[1] || !found[2] || !found[3] {
		t.Error("ToSlice() should contain all set elements")
	}
}

func testSetForEach(t *testing.T, factory func() setx.Set[int]) {
	set := factory()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	visited := make(map[int]bool)
	set.ForEach(func(element int) {
		visited[element] = true
	})

	if len(visited) != 3 {
		t.Errorf("Expected to visit 3 elements, visited %d", len(visited))
	}

	if !visited[1] || !visited[2] || !visited[3] {
		t.Error("ForEach should visit all elements")
	}
}

func testSetUnion(t *testing.T, factory func() setx.Set[int]) {
	set1 := factory()
	set2 := factory()

	set1.Add(1)
	set1.Add(2)
	set2.Add(2)
	set2.Add(3)

	union := set1.Union(set2)

	if union.Size() != 3 {
		t.Errorf("Union size should be 3, got %d", union.Size())
	}

	if !union.Contains(1) || !union.Contains(2) || !union.Contains(3) {
		t.Error("Union should contain elements from both sets")
	}
}

func testSetIntersection(t *testing.T, factory func() setx.Set[int]) {
	set1 := factory()
	set2 := factory()

	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	set2.Add(2)
	set2.Add(3)
	set2.Add(4)

	intersection := set1.Intersection(set2)

	if intersection.Size() != 2 {
		t.Errorf("Intersection size should be 2, got %d", intersection.Size())
	}

	if !intersection.Contains(2) || !intersection.Contains(3) {
		t.Error("Intersection should contain common elements")
	}

	if intersection.Contains(1) || intersection.Contains(4) {
		t.Error("Intersection should not contain non-common elements")
	}
}

func testSetDifference(t *testing.T, factory func() setx.Set[int]) {
	set1 := factory()
	set2 := factory()

	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	set2.Add(2)
	set2.Add(4)

	difference := set1.Difference(set2)

	if difference.Size() != 2 {
		t.Errorf("Difference size should be 2, got %d", difference.Size())
	}

	if !difference.Contains(1) || !difference.Contains(3) {
		t.Error("Difference should contain elements only in first set")
	}

	if difference.Contains(2) || difference.Contains(4) {
		t.Error("Difference should not contain common or second set only elements")
	}
}

func testSetIsSubsetOf(t *testing.T, factory func() setx.Set[int]) {
	set1 := factory()
	set2 := factory()

	set1.Add(1)
	set1.Add(2)
	set2.Add(1)
	set2.Add(2)
	set2.Add(3)

	if !set1.IsSubsetOf(set2) {
		t.Error("set1 should be subset of set2")
	}

	if set2.IsSubsetOf(set1) {
		t.Error("set2 should not be subset of set1")
	}

	// Test with empty set
	emptySet := factory()
	if !emptySet.IsSubsetOf(set1) {
		t.Error("Empty set should be subset of any set")
	}
}

func testSetIsSupersetOf(t *testing.T, factory func() setx.Set[int]) {
	set1 := factory()
	set2 := factory()

	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	set2.Add(1)
	set2.Add(2)

	if !set1.IsSupersetOf(set2) {
		t.Error("set1 should be superset of set2")
	}

	if set2.IsSupersetOf(set1) {
		t.Error("set2 should not be superset of set1")
	}

	// Test with empty set
	emptySet := factory()
	if !set1.IsSupersetOf(emptySet) {
		t.Error("Any set should be superset of empty set")
	}
}

// Test functions for new Option/Result-based methods

func testSetFind(t *testing.T, factory func() setx.Set[int]) {
	set := factory()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)
	set.Add(5)

	// Find even number
	result := set.Find(func(x int) bool { return x%2 == 0 })
	if result.IsNone() {
		t.Error("Should find an even number")
	}

	value := result.Unwrap()
	if value%2 != 0 {
		t.Errorf("Found value should be even, got %d", value)
	}

	// Find number greater than 10 (should not exist)
	result = set.Find(func(x int) bool { return x > 10 })
	if result.IsSome() {
		t.Error("Should not find number greater than 10")
	}
}

func testSetGetAny(t *testing.T, factory func() setx.Set[int]) {
	set := factory()

	// Empty set
	result := set.GetAny()
	if result.IsSome() {
		t.Error("Empty set should return None")
	}

	// Non-empty set
	set.Add(42)
	result = set.GetAny()
	if result.IsNone() {
		t.Error("Non-empty set should return Some")
	}
	if result.Unwrap() != 42 {
		t.Errorf("Expected 42, got %d", result.Unwrap())
	}
}

func testSetTryRemove(t *testing.T, factory func() setx.Set[int]) {
	set := factory()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	// Remove existing element
	result := set.TryRemove(2)
	if result.IsErr() {
		t.Errorf("Should successfully remove existing element: %v", result.UnwrapErr())
	}
	if result.Unwrap() != 2 {
		t.Errorf("Expected removed element to be 2, got %d", result.Unwrap())
	}

	if set.Contains(2) {
		t.Error("Element 2 should be removed from set")
	}

	// Try to remove non-existing element
	result = set.TryRemove(10)
	if result.IsOk() {
		t.Error("Should fail to remove non-existing element")
	}
}

func testSetFilter(t *testing.T, factory func() setx.Set[int]) {
	set := factory()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)
	set.Add(5)

	// Filter even numbers
	evenSet := set.Filter(func(x int) bool { return x%2 == 0 })

	if evenSet.Size() != 2 {
		t.Errorf("Expected 2 even numbers, got %d", evenSet.Size())
	}

	if !evenSet.Contains(2) || !evenSet.Contains(4) {
		t.Error("Even set should contain 2 and 4")
	}

	if evenSet.Contains(1) || evenSet.Contains(3) || evenSet.Contains(5) {
		t.Error("Even set should not contain odd numbers")
	}

	// Filter with no matches
	largeSet := set.Filter(func(x int) bool { return x > 10 })
	if !largeSet.IsEmpty() {
		t.Error("Filter with no matches should return empty set")
	}
}
