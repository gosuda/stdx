package concurrentmap_test

import (
	"sync"
	"testing"

	"github.com/gosuda/stdx/mapx"
	"github.com/gosuda/stdx/mapx/concurrentmap"
)

// createConcurrentMap is a factory function for creating ConcurrentMap instances
func createConcurrentMap[K comparable, V any]() mapx.Map[K, V] {
	return concurrentmap.New[K, V]()
}

func TestConcurrentMap_Put(t *testing.T) {
	testMapPut(t, createConcurrentMap[string, int])
}

func TestConcurrentMap_Get(t *testing.T) {
	testMapGet(t, createConcurrentMap[string, int])
}

func TestConcurrentMap_Remove(t *testing.T) {
	testMapRemove(t, createConcurrentMap[string, int])
}

func TestConcurrentMap_ContainsKey(t *testing.T) {
	testMapContainsKey(t, createConcurrentMap[string, int])
}

func TestConcurrentMap_ContainsValue(t *testing.T) {
	testMapContainsValue(t, createConcurrentMap[string, int])
}

func TestConcurrentMap_Size(t *testing.T) {
	testMapSize(t, createConcurrentMap[string, int])
}

func TestConcurrentMap_IsEmpty(t *testing.T) {
	testMapIsEmpty(t, createConcurrentMap[string, int])
}

func TestConcurrentMap_Clear(t *testing.T) {
	testMapClear(t, createConcurrentMap[string, int])
}

func TestConcurrentMap_Keys(t *testing.T) {
	testMapKeys(t, createConcurrentMap[string, int])
}

func TestConcurrentMap_Values(t *testing.T) {
	testMapValues(t, createConcurrentMap[string, int])
}

func TestConcurrentMap_Entries(t *testing.T) {
	testMapEntries(t, createConcurrentMap[string, int])
}

func TestConcurrentMap_ForEach(t *testing.T) {
	testMapForEach(t, createConcurrentMap[string, int])
}

func TestConcurrentMap_FindKey(t *testing.T) {
	testMapFindKey(t, createConcurrentMap[string, int])
}

func TestConcurrentMap_FindEntry(t *testing.T) {
	testMapFindEntry(t, createConcurrentMap[string, int])
}

func TestConcurrentMap_Filter(t *testing.T) {
	testMapFilter(t, createConcurrentMap[string, int])
}

// Concurrent-specific tests
func TestConcurrentMap_ConcurrentPut(t *testing.T) {
	m := concurrentmap.New[int, int]()
	const numGoroutines = 100
	const numOperations = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Launch multiple goroutines to put key-value pairs concurrently
	for i := 0; i < numGoroutines; i++ {
		go func(start int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				m.Put(start*numOperations+j, start*numOperations+j)
			}
		}(i)
	}

	wg.Wait()

	expectedSize := numGoroutines * numOperations
	if m.Size() != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, m.Size())
	}
}

func TestConcurrentMap_ConcurrentPutGet(t *testing.T) {
	m := concurrentmap.New[int, int]()
	const numElements = 1000
	const numReaders = 10
	const numWriters = 5

	// Pre-populate the map
	for i := 0; i < numElements; i++ {
		m.Put(i, i*2)
	}

	var wg sync.WaitGroup
	wg.Add(numReaders + numWriters)

	// Launch reader goroutines - they just verify that keys exist and can be read
	for i := 0; i < numReaders; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numElements; j++ {
				valOpt := m.Get(j)
				if valOpt.IsNone() {
					t.Errorf("Key %d should exist in map", j)
				}
			}
		}()
	}

	// Launch writer goroutines
	for i := 0; i < numWriters; i++ {
		go func(writerID int) {
			defer wg.Done()
			for j := 0; j < numElements; j++ {
				// Update values
				m.Put(j, j*3+writerID)
			}
		}(i)
	}

	wg.Wait()

	// Test should complete without deadlocks or data races
	if m.Size() != numElements {
		t.Errorf("Expected size %d, got %d", numElements, m.Size())
	}
}

func TestConcurrentMap_ConcurrentPutRemove(t *testing.T) {
	m := concurrentmap.New[int, int]()
	const numGoroutines = 50
	const numOperations = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines * 2) // numGoroutines for put, numGoroutines for remove

	// Launch goroutines to put elements
	for i := 0; i < numGoroutines; i++ {
		go func(start int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				m.Put(start*numOperations+j, start*numOperations+j)
			}
		}(i)
	}

	// Launch goroutines to remove elements
	for i := 0; i < numGoroutines; i++ {
		go func(start int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				m.Remove(start*numOperations + j)
			}
		}(i)
	}

	wg.Wait()

	// The map should be consistent (no race conditions should cause corruption)
	// The exact final size is non-deterministic due to timing
	finalSize := m.Size()
	if finalSize < 0 || finalSize > numGoroutines*numOperations {
		t.Errorf("Final size %d is out of expected range [0, %d]", finalSize, numGoroutines*numOperations)
	}
}

func TestConcurrentMap_ConcurrentIteration(t *testing.T) {
	m := concurrentmap.New[int, int]()
	const numElements = 500

	// Pre-populate the map
	for i := 0; i < numElements; i++ {
		m.Put(i, i*2)
	}

	var wg sync.WaitGroup
	wg.Add(3) // 3 different iteration methods

	// Test concurrent Keys()
	go func() {
		defer wg.Done()
		keys := m.Keys()
		if len(keys) > numElements {
			t.Errorf("Keys length %d is out of range", len(keys))
		}
	}()

	// Test concurrent Values()
	go func() {
		defer wg.Done()
		values := m.Values()
		if len(values) > numElements {
			t.Errorf("Values length %d is out of range", len(values))
		}
	}()

	// Test concurrent ForEach()
	go func() {
		defer wg.Done()
		count := 0
		m.ForEach(func(key int, value int) {
			count++
		})
		if count < 0 || count > numElements {
			t.Errorf("ForEach count %d is out of range", count)
		}
	}()

	wg.Wait()

	// Test should complete without deadlocks
}

// Include the same common test functions as in hashmap_test.go

func testMapPut(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()

	// Test putting new key-value pair
	prevOpt := m.Put("key1", 100)
	if prevOpt.IsSome() {
		t.Error("Put should return None for new key")
	}

	// Test updating existing key
	prevOpt = m.Put("key1", 200)
	if prevOpt.IsNone() {
		t.Error("Put should return Some for existing key")
	}
	if prevOpt.Unwrap() != 100 {
		t.Errorf("Previous value should be 100, got %d", prevOpt.Unwrap())
	}

	// Verify the value was updated
	valOpt := m.Get("key1")
	if valOpt.IsNone() {
		t.Error("Get should return Some for existing key")
	}
	if valOpt.Unwrap() != 200 {
		t.Errorf("Expected value 200, got %d", valOpt.Unwrap())
	}
}

func testMapGet(t *testing.T, factory func() mapx.Map[string, int]) {
	m := factory()
	m.Put("key1", 100)
	m.Put("key2", 200)

	// Test getting existing key
	valOpt := m.Get("key1")
	if valOpt.IsNone() {
		t.Error("Get should return Some for existing key")
	}
	if valOpt.Unwrap() != 100 {
		t.Errorf("Expected value 100, got %d", valOpt.Unwrap())
	}

	// Test getting non-existing key
	valOpt = m.Get("key3")
	if valOpt.IsSome() {
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
		t.Errorf("Remove should succeed for existing key, got error: %v", result.UnwrapErr())
	}
	if result.Unwrap() != 100 {
		t.Errorf("Removed value should be 100, got %d", result.Unwrap())
	}

	// Verify key was removed
	valOpt := m.Get("key1")
	if valOpt.IsSome() {
		t.Error("Key should not exist after removal")
	}

	// Test removing non-existing key
	result = m.Remove("key3")
	if result.IsOk() {
		t.Error("Remove should return error for non-existing key")
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
