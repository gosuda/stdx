package synx

import (
	"strings"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	t.Run("Basic Pool operations", func(t *testing.T) {
		pool := NewPool(func() string {
			return "default"
		})

		// Test Get with factory
		item := pool.Get()
		if item != "default" {
			t.Errorf("Expected 'default', got %v", item)
		}

		// Test Put and Get
		pool.Put("custom")
		item = pool.Get()
		if item != "custom" {
			t.Errorf("Expected 'custom', got %v", item)
		}
	})

	t.Run("Integer Pool", func(t *testing.T) {
		pool := NewPool(func() int {
			return 42
		})

		item := pool.Get()
		if item != 42 {
			t.Errorf("Expected 42, got %v", item)
		}

		pool.Put(100)
		item = pool.Get()
		if item != 100 {
			t.Errorf("Expected 100, got %v", item)
		}
	})

	t.Run("Struct Pool", func(t *testing.T) {
		type TestStruct struct {
			Name string
			ID   int
		}

		pool := NewPool(func() *TestStruct {
			return &TestStruct{Name: "default", ID: 0}
		})

		item := pool.Get()
		if item.Name != "default" || item.ID != 0 {
			t.Errorf("Expected default struct, got %+v", item)
		}

		customStruct := &TestStruct{Name: "custom", ID: 123}
		pool.Put(customStruct)

		item = pool.Get()
		if item.Name != "custom" || item.ID != 123 {
			t.Errorf("Expected custom struct, got %+v", item)
		}
	})

	t.Run("TryGet", func(t *testing.T) {
		pool := NewPool(func() string {
			return "factory"
		})

		// Pool is empty initially, TryGet should return zero value and false
		item, ok := pool.TryGet()
		if ok || item != "factory" {
			t.Errorf("Expected factory default and true, got %v, %v", item, ok)
		}

		// Put something and try again
		pool.Put("exists")
		item, ok = pool.TryGet()
		if !ok || item != "exists" {
			t.Errorf("Expected 'exists' and true, got %v, %v", item, ok)
		}
	})

	t.Run("Reset", func(t *testing.T) {
		pool := NewPool(func() string {
			return "original"
		})

		pool.Put("test")
		item := pool.Get()
		if item != "test" {
			t.Errorf("Expected 'test', got %v", item)
		}

		// Reset with new factory
		pool.Reset(func() string {
			return "new"
		})

		item = pool.Get()
		if item != "new" {
			t.Errorf("Expected 'new', got %v", item)
		}
	})
}

func TestStringPool(t *testing.T) {
	t.Run("Basic string pool", func(t *testing.T) {
		pool := NewStringPool()

		str := pool.Get()
		if str != "" {
			t.Errorf("Expected empty string, got %v", str)
		}

		pool.Put("hello")
		str = pool.Get()
		if str != "hello" {
			t.Errorf("Expected 'hello', got %v", str)
		}
	})
}

func TestByteSlicePool(t *testing.T) {
	t.Run("Basic byte slice pool", func(t *testing.T) {
		pool := NewByteSlicePool(10)

		buf := pool.Get()
		if len(buf) != 0 || cap(buf) != 10 {
			t.Errorf("Expected empty slice with cap 10, got len=%d cap=%d", len(buf), cap(buf))
		}

		// Use the buffer
		buf = append(buf, []byte("hello")...)
		pool.PutReset(buf)

		// Get it back - should be reset
		buf = pool.Get()
		if len(buf) != 0 {
			t.Errorf("Expected reset slice, got len=%d", len(buf))
		}
	})

	t.Run("GetWithCap", func(t *testing.T) {
		pool := NewByteSlicePool(5)

		// Request larger capacity
		buf := pool.GetWithCap(20)
		if cap(buf) < 20 {
			t.Errorf("Expected capacity >= 20, got %d", cap(buf))
		}

		// Request smaller capacity (should use pool)
		pool.Put(make([]byte, 0, 30))
		buf = pool.GetWithCap(10)
		if cap(buf) != 30 {
			t.Errorf("Expected to reuse buffer with cap 30, got %d", cap(buf))
		}
	})
}

func TestPoolConcurrency(t *testing.T) {
	t.Run("Concurrent access", func(t *testing.T) {
		pool := NewPool(func() *strings.Builder {
			return &strings.Builder{}
		})

		const numGoroutines = 100
		const numOperations = 100

		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer wg.Done()

				for j := 0; j < numOperations; j++ {
					builder := pool.Get()
					builder.Reset()
					builder.WriteString("test")

					if builder.String() != "test" {
						t.Errorf("Builder corrupted: got %v", builder.String())
					}

					pool.Put(builder)
				}
			}(i)
		}

		wg.Wait()
	})
}
