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

// Common test functions that can be reused for any Map implementation

func testMapPut(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()

	// Test putting new key-value pair
	prev, exists := m.Put("key1", 100)
	if exists {
		t.Error("Put should return false for new key")
	}
	if prev != 0 {
		t.Errorf("Previous value should be zero value, got %d", prev)
	}

	// Test updating existing key
	prev, exists = m.Put("key1", 200)
	if !exists {
		t.Error("Put should return true for existing key")
	}
	if prev != 100 {
		t.Errorf("Previous value should be 100, got %d", prev)
	}

	// Verify the value was updated
	val, found := m.Get("key1")
	if !found || val != 200 {
		t.Errorf("Expected value 200, got %d", val)
	}
}

func testMapGet(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()
	m.Put("key1", 100)
	m.Put("key2", 200)

	// Test getting existing key
	val, exists := m.Get("key1")
	if !exists {
		t.Error("Get should return true for existing key")
	}
	if val != 100 {
		t.Errorf("Expected value 100, got %d", val)
	}

	// Test getting non-existing key
	val, exists = m.Get("key3")
	if exists {
		t.Error("Get should return false for non-existing key")
	}
	if val != 0 {
		t.Errorf("Value should be zero value for non-existing key, got %d", val)
	}
}

func testMapRemove(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()
	m.Put("key1", 100)
	m.Put("key2", 200)

	// Test removing existing key
	val, removed := m.Remove("key1")
	if !removed {
		t.Error("Remove should return true for existing key")
	}
	if val != 100 {
		t.Errorf("Removed value should be 100, got %d", val)
	}

	// Verify key was removed
	_, exists := m.Get("key1")
	if exists {
		t.Error("Key should not exist after removal")
	}

	// Test removing non-existing key
	val, removed = m.Remove("key3")
	if removed {
		t.Error("Remove should return false for non-existing key")
	}
	if val != 0 {
		t.Errorf("Value should be zero value for non-existing key, got %d", val)
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
