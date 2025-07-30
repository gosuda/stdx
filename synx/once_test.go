package synx

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestOnce(t *testing.T) {
	t.Run("Basic Once Do", func(t *testing.T) {
		var once Once[string]
		var callCount int32

		result1 := once.Do(func() string {
			atomic.AddInt32(&callCount, 1)
			return "hello"
		})

		result2 := once.Do(func() string {
			atomic.AddInt32(&callCount, 1)
			return "world"
		})

		if result1 != "hello" {
			t.Errorf("Expected 'hello', got %v", result1)
		}

		if result2 != "hello" {
			t.Errorf("Expected cached 'hello', got %v", result2)
		}

		if callCount != 1 {
			t.Errorf("Expected function to be called once, got %d times", callCount)
		}
	})

	t.Run("Once DoWithError", func(t *testing.T) {
		var once Once[int]
		var callCount int32

		// First call - success
		result1, err1 := once.DoWithError(func() (int, error) {
			atomic.AddInt32(&callCount, 1)
			return 42, nil
		})

		// Second call - should return cached result even with different function
		result2, err2 := once.DoWithError(func() (int, error) {
			atomic.AddInt32(&callCount, 1)
			return 100, errors.New("should not be called")
		})

		if result1 != 42 || err1 != nil {
			t.Errorf("Expected 42 and no error, got %v and %v", result1, err1)
		}

		if result2 != 42 || err2 != nil {
			t.Errorf("Expected cached 42 and no error, got %v and %v", result2, err2)
		}

		if callCount != 1 {
			t.Errorf("Expected function to be called once, got %d times", callCount)
		}
	})

	t.Run("Once DoWithError - error case", func(t *testing.T) {
		var once Once[string]
		expectedErr := errors.New("test error")

		result1, err1 := once.DoWithError(func() (string, error) {
			return "", expectedErr
		})

		result2, err2 := once.DoWithError(func() (string, error) {
			return "success", nil // This should not be called
		})

		if result1 != "" || err1 != expectedErr {
			t.Errorf("Expected empty string and test error, got %v and %v", result1, err1)
		}

		if result2 != "" || err2 != expectedErr {
			t.Errorf("Expected cached empty string and test error, got %v and %v", result2, err2)
		}
	})

	t.Run("Once Value and Error getters", func(t *testing.T) {
		var once Once[string]

		// Before calling Do
		if once.Value() != "" {
			t.Errorf("Expected zero value before Do, got %v", once.Value())
		}

		// After calling DoWithError
		once.DoWithError(func() (string, error) {
			return "test", errors.New("test error")
		})

		if once.Value() != "test" {
			t.Errorf("Expected 'test', got %v", once.Value())
		}

		if once.Error().Error() != "test error" {
			t.Errorf("Expected 'test error', got %v", once.Error())
		}
	})

	t.Run("Once Reset", func(t *testing.T) {
		var once Once[int]

		result1 := once.Do(func() int {
			return 42
		})

		if result1 != 42 {
			t.Errorf("Expected 42, got %v", result1)
		}

		once.Reset()

		result2 := once.Do(func() int {
			return 100
		})

		if result2 != 100 {
			t.Errorf("Expected 100 after reset, got %v", result2)
		}
	})
}

func TestOnceValue(t *testing.T) {
	t.Run("OnceValue function", func(t *testing.T) {
		var callCount int32

		getValue := OnceValue(func() string {
			atomic.AddInt32(&callCount, 1)
			return "computed"
		})

		result1 := getValue()
		result2 := getValue()

		if result1 != "computed" || result2 != "computed" {
			t.Errorf("Expected 'computed' for both calls, got %v and %v", result1, result2)
		}

		if callCount != 1 {
			t.Errorf("Expected function to be called once, got %d times", callCount)
		}
	})

	t.Run("OnceValues function", func(t *testing.T) {
		var callCount int32

		getValues := OnceValues(func() (string, int) {
			atomic.AddInt32(&callCount, 1)
			return "hello", 42
		})

		str1, num1 := getValues()
		str2, num2 := getValues()

		if str1 != "hello" || num1 != 42 {
			t.Errorf("Expected 'hello' and 42, got %v and %v", str1, num1)
		}

		if str2 != "hello" || num2 != 42 {
			t.Errorf("Expected cached 'hello' and 42, got %v and %v", str2, num2)
		}

		if callCount != 1 {
			t.Errorf("Expected function to be called once, got %d times", callCount)
		}
	})
}

func TestOnceFunc(t *testing.T) {
	t.Run("OnceFunc", func(t *testing.T) {
		var callCount int32

		doSomething := OnceFunc(func() {
			atomic.AddInt32(&callCount, 1)
		})

		doSomething()
		doSomething()
		doSomething()

		if callCount != 1 {
			t.Errorf("Expected function to be called once, got %d times", callCount)
		}
	})

	t.Run("OnceFuncWithArg", func(t *testing.T) {
		var callCount int32
		var receivedArg string

		doWithArg := OnceFuncWithArg(func(arg string) {
			atomic.AddInt32(&callCount, 1)
			receivedArg = arg
		})

		doWithArg("first")
		doWithArg("second") // This should be ignored

		if callCount != 1 {
			t.Errorf("Expected function to be called once, got %d times", callCount)
		}

		if receivedArg != "first" {
			t.Errorf("Expected 'first', got %v", receivedArg)
		}
	})
}

func TestOnceResult(t *testing.T) {
	t.Run("OnceResult success", func(t *testing.T) {
		var callCount int32

		getResult := OnceResult(func() (string, error) {
			atomic.AddInt32(&callCount, 1)
			return "success", nil
		})

		result1, err1 := getResult()
		result2, err2 := getResult()

		if result1 != "success" || err1 != nil {
			t.Errorf("Expected 'success' and no error, got %v and %v", result1, err1)
		}

		if result2 != "success" || err2 != nil {
			t.Errorf("Expected cached 'success' and no error, got %v and %v", result2, err2)
		}

		if callCount != 1 {
			t.Errorf("Expected function to be called once, got %d times", callCount)
		}
	})

	t.Run("OnceResult error", func(t *testing.T) {
		var callCount int32
		expectedErr := errors.New("test error")

		getResult := OnceResult(func() (string, error) {
			atomic.AddInt32(&callCount, 1)
			return "", expectedErr
		})

		result1, err1 := getResult()
		result2, err2 := getResult()

		if result1 != "" || err1 != expectedErr {
			t.Errorf("Expected empty string and test error, got %v and %v", result1, err1)
		}

		if result2 != "" || err2 != expectedErr {
			t.Errorf("Expected cached empty string and test error, got %v and %v", result2, err2)
		}

		if callCount != 1 {
			t.Errorf("Expected function to be called once, got %d times", callCount)
		}
	})
}

func TestLazyValue(t *testing.T) {
	t.Run("Basic lazy initialization", func(t *testing.T) {
		var initCount int32

		lazy := NewLazyValue(func() string {
			atomic.AddInt32(&initCount, 1)
			return "initialized"
		})

		// Check not initialized initially
		if lazy.IsInitialized() {
			t.Error("Expected lazy value to not be initialized initially")
		}

		// Get value - should initialize
		value1 := lazy.Get()
		if value1 != "initialized" {
			t.Errorf("Expected 'initialized', got %v", value1)
		}

		// Check initialized now
		if !lazy.IsInitialized() {
			t.Error("Expected lazy value to be initialized after Get()")
		}

		// Get again - should not reinitialize
		value2 := lazy.Get()
		if value2 != "initialized" {
			t.Errorf("Expected cached 'initialized', got %v", value2)
		}

		if initCount != 1 {
			t.Errorf("Expected init function to be called once, got %d times", initCount)
		}
	})

	t.Run("Lazy value reset", func(t *testing.T) {
		var initCount int32

		lazy := NewLazyValue(func() int {
			atomic.AddInt32(&initCount, 1)
			return 42
		})

		value1 := lazy.Get()
		if value1 != 42 {
			t.Errorf("Expected 42, got %v", value1)
		}

		lazy.Reset(func() int {
			atomic.AddInt32(&initCount, 1)
			return 100
		})

		value2 := lazy.Get()
		if value2 != 100 {
			t.Errorf("Expected 100 after reset, got %v", value2)
		}

		if initCount != 2 {
			t.Errorf("Expected init function to be called twice, got %d times", initCount)
		}
	})

	t.Run("Lazy value with nil initializer", func(t *testing.T) {
		lazy := NewLazyValue[string](nil)

		value := lazy.Get()
		if value != "" {
			t.Errorf("Expected zero value, got %v", value)
		}
	})
}

func TestConcurrency(t *testing.T) {
	t.Run("Once concurrent access", func(t *testing.T) {
		var once Once[int]
		var callCount int32
		const numGoroutines = 100

		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		results := make([]int, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(index int) {
				defer wg.Done()
				results[index] = once.Do(func() int {
					atomic.AddInt32(&callCount, 1)
					time.Sleep(10 * time.Millisecond) // Simulate work
					return 42
				})
			}(i)
		}

		wg.Wait()

		// Check all goroutines got the same result
		for i, result := range results {
			if result != 42 {
				t.Errorf("Goroutine %d got %v, expected 42", i, result)
			}
		}

		if callCount != 1 {
			t.Errorf("Expected function to be called once, got %d times", callCount)
		}
	})

	t.Run("LazyValue concurrent access", func(t *testing.T) {
		var initCount int32
		const numGoroutines = 100

		lazy := NewLazyValue(func() string {
			atomic.AddInt32(&initCount, 1)
			time.Sleep(10 * time.Millisecond) // Simulate work
			return "computed"
		})

		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		results := make([]string, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(index int) {
				defer wg.Done()
				results[index] = lazy.Get()
			}(i)
		}

		wg.Wait()

		// Check all goroutines got the same result
		for i, result := range results {
			if result != "computed" {
				t.Errorf("Goroutine %d got %v, expected 'computed'", i, result)
			}
		}

		if initCount != 1 {
			t.Errorf("Expected init function to be called once, got %d times", initCount)
		}
	})
}
