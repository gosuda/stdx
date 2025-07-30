package hashmap_test

import (
	"testing"

	"github.com/gosuda/stdx/mapx"
	"github.com/gosuda/stdx/mapx/hashmap"
)

// createHashMap is a factory function for creating HashMap instances
func createHashMap[K comparable, V any]() mapx.Map[K, V] {
	return hashmap.New[K, V]()
}

func TestHashMap_Put(t *testing.T) {
	testMapPut(t, createHashMap[string, int])
}

func TestHashMap_Get(t *testing.T) {
	testMapGet(t, createHashMap[string, int])
}

func TestHashMap_Remove(t *testing.T) {
	testMapRemove(t, createHashMap[string, int])
}

func TestHashMap_ContainsKey(t *testing.T) {
	testMapContainsKey(t, createHashMap[string, int])
}

func TestHashMap_ContainsValue(t *testing.T) {
	testMapContainsValue(t, createHashMap[string, int])
}

func TestHashMap_Size(t *testing.T) {
	testMapSize(t, createHashMap[string, int])
}

func TestHashMap_IsEmpty(t *testing.T) {
	testMapIsEmpty(t, createHashMap[string, int])
}

func TestHashMap_Clear(t *testing.T) {
	testMapClear(t, createHashMap[string, int])
}

func TestHashMap_Keys(t *testing.T) {
	testMapKeys(t, createHashMap[string, int])
}

func TestHashMap_Values(t *testing.T) {
	testMapValues(t, createHashMap[string, int])
}

func TestHashMap_Entries(t *testing.T) {
	testMapEntries(t, createHashMap[string, int])
}

func TestHashMap_ForEach(t *testing.T) {
	testMapForEach(t, createHashMap[string, int])
}

func TestHashMap_FindKey(t *testing.T) {
	testMapFindKey(t, createHashMap[string, int])
}

func TestHashMap_FindEntry(t *testing.T) {
	testMapFindEntry(t, createHashMap[string, int])
}

func TestHashMap_Filter(t *testing.T) {
	testMapFilter(t, createHashMap[string, int])
}

// Common test functions that can be reused for any Map implementation

func testMapPut(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()

	// Test putting new key-value pair
	result := m.Put("key1", 100)
	if result.IsSome() {
		t.Error("Put should return None for new key")
	}

	// Test updating existing key
	result = m.Put("key1", 200)
	if result.IsNone() {
		t.Error("Put should return Some for existing key")
	}
	if result.Unwrap() != 100 {
		t.Errorf("Previous value should be 100, got %d", result.Unwrap())
	}

	// Verify the value was updated
	getResult := m.Get("key1")
	if getResult.IsNone() || getResult.Unwrap() != 200 {
		t.Errorf("Expected value 200, got %d", getResult.UnwrapOr(0))
	}
}

func testMapGet(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()
	m.Put("key1", 100)
	m.Put("key2", 200)

	// Test getting existing key
	result := m.Get("key1")
	if result.IsNone() {
		t.Error("Get should return Some for existing key")
	}
	if result.Unwrap() != 100 {
		t.Errorf("Expected value 100, got %d", result.Unwrap())
	}

	// Test getting non-existing key
	result = m.Get("key3")
	if result.IsSome() {
		t.Error("Get should return None for non-existing key")
	}
}

func testMapRemove(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()
	m.Put("key1", 100)
	m.Put("key2", 200)

	// Test removing existing key
	result := m.Remove("key1")
	if result.IsErr() {
		t.Errorf("Remove should succeed for existing key: %v", result.UnwrapErr())
	}
	if result.Unwrap() != 100 {
		t.Errorf("Removed value should be 100, got %d", result.Unwrap())
	}

	// Verify key was removed
	getResult := m.Get("key1")
	if getResult.IsSome() {
		t.Error("Key should not exist after removal")
	}

	// Test removing non-existing key
	result = m.Remove("key3")
	if result.IsOk() {
		t.Error("Remove should fail for non-existing key")
	}

	// Test size after removal
	if m.Size() != 1 {
		t.Errorf("Expected size 1 after removal, got %d", m.Size())
	}
}

func testMapContainsKey(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()
	m.Put("key1", 100)
	m.Put("key2", 200)

	if !m.ContainsKey("key1") {
		t.Error("Map should contain key1")
	}
	if !m.ContainsKey("key2") {
		t.Error("Map should contain key2")
	}
	if m.ContainsKey("key3") {
		t.Error("Map should not contain key3")
	}
}

func testMapContainsValue(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()
	m.Put("key1", 100)
	m.Put("key2", 200)

	if !m.ContainsValue(100) {
		t.Error("Map should contain value 100")
	}
	if !m.ContainsValue(200) {
		t.Error("Map should contain value 200")
	}
	if m.ContainsValue(300) {
		t.Error("Map should not contain value 300")
	}
}

func testMapSize(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()

	if m.Size() != 0 {
		t.Errorf("Empty map size should be 0, got %d", m.Size())
	}

	m.Put("key1", 100)
	if m.Size() != 1 {
		t.Errorf("Size should be 1, got %d", m.Size())
	}

	m.Put("key2", 200)
	m.Put("key3", 300)
	if m.Size() != 3 {
		t.Errorf("Size should be 3, got %d", m.Size())
	}

	m.Remove("key2")
	if m.Size() != 2 {
		t.Errorf("Size should be 2 after removal, got %d", m.Size())
	}
}

func testMapIsEmpty(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()

	if !m.IsEmpty() {
		t.Error("New map should be empty")
	}

	m.Put("key1", 100)
	if m.IsEmpty() {
		t.Error("Map with elements should not be empty")
	}

	m.Remove("key1")
	if !m.IsEmpty() {
		t.Error("Map should be empty after removing all elements")
	}
}

func testMapClear(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()
	m.Put("key1", 100)
	m.Put("key2", 200)
	m.Put("key3", 300)

	m.Clear()

	if !m.IsEmpty() {
		t.Error("Map should be empty after Clear()")
	}
	if m.Size() != 0 {
		t.Errorf("Size should be 0 after Clear(), got %d", m.Size())
	}
	if m.ContainsKey("key1") || m.ContainsKey("key2") || m.ContainsKey("key3") {
		t.Error("Map should not contain any keys after Clear()")
	}
}

func testMapKeys(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()
	m.Put("key1", 100)
	m.Put("key2", 200)
	m.Put("key3", 300)

	keys := m.Keys()

	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	// Check all keys are present (order doesn't matter)
	keySet := make(map[string]bool)
	for _, k := range keys {
		keySet[k] = true
	}

	if !keySet["key1"] || !keySet["key2"] || !keySet["key3"] {
		t.Error("Keys() should contain all map keys")
	}
}

func testMapValues(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()
	m.Put("key1", 100)
	m.Put("key2", 200)
	m.Put("key3", 300)

	values := m.Values()

	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}

	// Check all values are present (order doesn't matter)
	valueSet := make(map[int]bool)
	for _, v := range values {
		valueSet[v] = true
	}

	if !valueSet[100] || !valueSet[200] || !valueSet[300] {
		t.Error("Values() should contain all map values")
	}
}

func testMapEntries(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()
	m.Put("key1", 100)
	m.Put("key2", 200)
	m.Put("key3", 300)

	entries := m.Entries()

	if len(entries) != 3 {
		t.Errorf("Expected 3 entries, got %d", len(entries))
	}

	// Check all entries are present (order doesn't matter)
	entryMap := make(map[string]int)
	for _, entry := range entries {
		entryMap[entry.Key] = entry.Value
	}

	if entryMap["key1"] != 100 || entryMap["key2"] != 200 || entryMap["key3"] != 300 {
		t.Error("Entries() should contain all key-value pairs")
	}
}

func testMapForEach(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()
	m.Put("key1", 100)
	m.Put("key2", 200)
	m.Put("key3", 300)

	visited := make(map[string]int)
	m.ForEach(func(key string, value int) {
		visited[key] = value
	})

	if len(visited) != 3 {
		t.Errorf("Expected to visit 3 entries, visited %d", len(visited))
	}

	if visited["key1"] != 100 || visited["key2"] != 200 || visited["key3"] != 300 {
		t.Error("ForEach should visit all key-value pairs")
	}
}

// Test functions for new Option/Result-based methods

func testMapFindKey(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()
	m.Put("key1", 100)
	m.Put("key2", 200)
	m.Put("key3", 100) // duplicate value

	// Find existing value
	result := m.FindKey(100)
	if result.IsNone() {
		t.Error("Should find a key for existing value")
	}

	foundKey := result.Unwrap()
	if foundKey != "key1" && foundKey != "key3" {
		t.Errorf("Found key should be 'key1' or 'key3', got %s", foundKey)
	}

	// Find non-existing value
	result = m.FindKey(999)
	if result.IsSome() {
		t.Error("Should not find key for non-existing value")
	}
}

func testMapFindEntry(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()
	m.Put("key1", 100)
	m.Put("key2", 200)
	m.Put("key3", 300)

	// Find entry where value > 150
	result := m.FindEntry(func(k string, v int) bool { return v > 150 })
	if result.IsNone() {
		t.Error("Should find an entry with value > 150")
	}

	entry := result.Unwrap()
	if entry.Value <= 150 {
		t.Errorf("Found entry value should be > 150, got %d", entry.Value)
	}

	// Find entry that doesn't exist
	result = m.FindEntry(func(k string, v int) bool { return v > 500 })
	if result.IsSome() {
		t.Error("Should not find entry with value > 500")
	}
}

func testMapFilter(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()
	m.Put("key1", 100)
	m.Put("key2", 200)
	m.Put("key3", 300)
	m.Put("key4", 400)

	// Filter entries with even values
	evenMap := m.Filter(func(k string, v int) bool { return v%200 == 0 })

	if evenMap.Size() != 2 {
		t.Errorf("Expected 2 entries with even hundreds, got %d", evenMap.Size())
	}

	if !evenMap.ContainsKey("key2") || !evenMap.ContainsKey("key4") {
		t.Error("Filtered map should contain key2 and key4")
	}

	if evenMap.ContainsKey("key1") || evenMap.ContainsKey("key3") {
		t.Error("Filtered map should not contain key1 or key3")
	}

	// Filter with no matches
	emptyMap := m.Filter(func(k string, v int) bool { return v > 1000 })
	if !emptyMap.IsEmpty() {
		t.Error("Filter with no matches should return empty map")
	}
}
