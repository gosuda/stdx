# stdx

A comprehensive standard library extension for Go, providing commonly used data structures, functional programming utilities, and advanced synchronization primitives.

## 📦 Packages

### 🗃️ Core Data Structures

#### **`listx`** - List Interfaces and Implementations
- **`listx/linked`** - Doubly linked list implementation
- **`listx/slices`** - Slice-based implementation  
- **`listx/hash`** - Hash table-based implementation
- **Interfaces**: `List[T]`, `Deque[T]`, `Stack[T]`, `Queue[T]`

#### **`mapx`** - Map Interfaces and Implementations
- **`mapx/hashmap`** - Standard hash map implementation
- **`mapx/concurrentmap`** - Thread-safe concurrent map using `sync.Map`
- **Interface**: `Map[K, V]` with advanced operations

#### **`setx`** - Set Interfaces and Implementations
- **`setx/hashset`** - Hash-based set implementation
- **`setx/concurrentset`** - Thread-safe concurrent set
- **Interface**: `Set[T]` with set operations (union, intersection, difference)

### 🧠 Functional Programming

#### **`option`** - Rust-Inspired Optional Values
- `Option[T]` type for handling nullable values safely
- Methods: `Some()`, `None()`, `Map()`, `Filter()`, `UnwrapOr()`
- JSON serialization support

#### **`result`** - Rust-Inspired Error Handling
- `Result[T, E]` type for error propagation
- Methods: `Ok()`, `Err()`, `Map()`, `FlatMap()`, `Match()`
- Conversion utilities like `Try()` and `TryWith()`

#### **`cond`** - Conditional Expressions
- Lisp-style conditional expressions
- `Cond[T]` for complex branching logic
- Switch-like pattern matching with `Switch[T, U]`

### 🧮 Tuple Types

#### **`tuple`** - Generic Tuple Types
- **`Pair[T, U]`** - Two-element tuples with functional operations
- **`Triple[T, U, V]`** - Three-element tuples
- Operations: `Map()`, `Swap()`, `Apply()`, currying/uncurrying
- Type-safe transformations with `MapFirst()`, `MapSecond()`

### 🔄 Synchronization

#### **`synx`** - Advanced Synchronization Primitives
- **`Once[T]`** - Generic wrapper around `sync.Once` with return values
- **`Pool[T]`** - Type-safe object pooling with `sync.Pool`
- **`LazyValue[T]`** - Lazy initialization with thread safety
- **`OnceValue[T]`**, **`OnceFunc`** - Function-based once execution

## ✨ Features

- **🎯 Type Safety**: All data structures use Go 1.18+ generics for compile-time type safety
- **⚡ Multiple Implementations**: Different backing implementations optimized for different use cases
- **🧪 Comprehensive Testing**: Full test coverage with factory patterns for implementation testing
- **🦀 Rust-Inspired**: Functional programming constructs inspired by Rust's type system
- **⚙️ Thread Safety**: Concurrent implementations where needed with proper synchronization
- **📋 Rich APIs**: Extensive method sets with functional programming patterns
- **🔗 Composable**: Designed for easy composition and chaining of operations

## 🚀 Quick Start

```go
package main

import (
    "fmt"
    "github.com/gosuda/stdx/listx/linked"
    "github.com/gosuda/stdx/option"
    "github.com/gosuda/stdx/result"
    "github.com/gosuda/stdx/setx/hashset"
    "github.com/gosuda/stdx/tuple"
)

func main() {
    // Using List
    list := linked.New[int]()
    list.Add(1)
    list.Add(2)
    list.Add(3)
    
    // Safe indexing with Option
    value := list.Get(1)
    if value.IsSome() {
        fmt.Println("Value at index 1:", value.Unwrap())
    }
    
    // Using Set operations
    set1 := hashset.New[int]()
    set1.Add(1)
    set1.Add(2)
    
    set2 := hashset.New[int]()
    set2.Add(2)
    set2.Add(3)
    
    union := set1.Union(set2)
    fmt.Println("Union:", union.ToSlice()) // [1, 2, 3]
    
    // Using Option for safe operations
    opt := option.Some(42)
    doubled := option.Map(opt, func(x int) int { return x * 2 })
    fmt.Println("Doubled:", doubled.UnwrapOr(0)) // 84
    
    // Using Result for error handling
    res := result.Try(divide(10, 2))
    res.Match(
        func(value int) { fmt.Println("Success:", value) },
        func(err error) { fmt.Println("Error:", err) },
    )
    
    // Using Tuples
    pair := tuple.NewPair("hello", 42)
    swapped := pair.Swap()
    fmt.Printf("Swapped: (%v, %v)\n", swapped.First(), swapped.Second())
}

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

## 📚 Detailed Examples

### Lists, Stacks, and Queues

```go
// Different list implementations
linkedList := linked.New[string]()
hashList := hash.New[string]() 
sliceList := slices.New[string]()

// Stack operations (LIFO)
stack := linked.NewStack[int]()
stack.Push(1)
stack.Push(2)
value := stack.Pop() // Result[int, error]
if value.IsOk() {
    fmt.Println("Popped:", value.Unwrap())
}

// Queue operations (FIFO)
queue := linked.NewQueue[string]()
queue.Enqueue("first")
queue.Enqueue("second")
item := queue.Dequeue() // Result[string, error]

// Deque (double-ended queue)
deque := linked.NewDeque[int]()
deque.AddFirst(1)
deque.AddLast(2)
first := deque.RemoveFirst() // Result[int, error]
last := deque.RemoveLast()   // Result[int, error]
```

### Maps with Advanced Operations

```go
// Hash map
hashMap := hashmap.New[string, int]()
hashMap.Put("key1", 100)
hashMap.Put("key2", 200)

// Safe retrieval
value := hashMap.Get("key1") // Option[int]
if value.IsSome() {
    fmt.Println("Found:", value.Unwrap())
}

// Try operations that return Results
removed := hashMap.TryRemove("key1") // Result[int, error]
removed.Match(
    func(val int) { fmt.Println("Removed:", val) },
    func(err error) { fmt.Println("Not found:", err) },
)

// Concurrent map for thread-safe operations
concMap := concurrentmap.New[string, int]()
// Same interface as hashmap but thread-safe
```

### Sets with Set Theory Operations

```go
set1 := hashset.New[int]()
set1.Add(1)
set1.Add(2)
set1.Add(3)

set2 := hashset.New[int]()
set2.Add(2)
set2.Add(3)
set2.Add(4)

// Set operations
union := set1.Union(set2)           // {1, 2, 3, 4}
intersection := set1.Intersection(set2) // {2, 3}
difference := set1.Difference(set2)     // {1}

// Predicate-based operations
evenNumbers := set1.Filter(func(x int) bool { return x%2 == 0 })
found := set1.Find(func(x int) bool { return x > 2 }) // Option[int]

// Safe removal
removed := set1.TryRemove(2) // Result[int, error]
```

### Option and Result Patterns

```go
// Option chaining
result := option.Some("hello").
    Map(func(s string) string { return s + " world" }).
    Filter(func(s string) bool { return len(s) > 5 }).
    UnwrapOr("default")

// Result error handling
func riskyOperation() result.Result[int, string] {
    if rand.Float32() < 0.5 {
        return result.Ok[int, string](42)
    }
    return result.Err[int, string]("something went wrong")
}

// Chaining results
finalResult := result.FlatMap(riskyOperation(), func(x int) result.Result[int, string] {
    return result.Ok[int, string](x * 2)
})

// Pattern matching
finalResult.Match(
    func(value int) { fmt.Println("Success:", value) },
    func(err string) { fmt.Println("Error:", err) },
)
```

### Conditional Expressions

```go
// Cond expressions (like Lisp's cond)
score := 85
grade := cond.New[string]().
    When(func() bool { return score >= 90 }, func() string { return "A" }).
    When(func() bool { return score >= 80 }, func() string { return "B" }).
    When(func() bool { return score >= 70 }, func() string { return "C" }).
    ElseValue("F").
    MustEval()

// Switch-like expressions
day := "Monday"
mood := cond.Switch(day, func(a, b string) bool { return a == b }).
    CaseValue("Monday", "😴").
    CaseValue("Friday", "🎉").
    CaseValue("Saturday", "😎").
    Default(func() string { return "😐" }).
    MustEval()
```

### Tuple Operations

```go
// Pair operations
pair := tuple.NewPair(10, "hello")
doubled := tuple.MapFirst(pair, func(x int) int { return x * 2 })
// Result: (20, "hello")

// Applying functions to pairs
sum := tuple.Apply(tuple.NewPair(3, 4), func(a, b int) int { return a + b })
// Result: 7

// Currying and uncurrying
add := func(a, b int) int { return a + b }
curriedAdd := tuple.Curry(add)
add5 := curriedAdd(5)
result := add5(3) // 8

// Triple operations
triple := tuple.NewTriple(1, "hello", true)
rotated := triple.RotateLeft() // ("hello", true, 1)
```

### Synchronization Utilities

```go
// Once with return value
var expensiveComputation synx.Once[string]
result := expensiveComputation.Do(func() string {
    // This will only run once
    time.Sleep(time.Second)
    return "computed value"
})

// Type-safe object pooling
pool := synx.NewPool(func() *bytes.Buffer {
    return new(bytes.Buffer)
})

buffer := pool.Get()
buffer.WriteString("some data")
// ... use buffer ...
buffer.Reset()
pool.Put(buffer)

// Lazy initialization
lazy := synx.NewLazyValue(func() string {
    // Expensive initialization
    return "lazy value"
})

// This will initialize only on first call
value := lazy.Get()
```

## 🔧 Installation

```bash
go get github.com/gosuda/stdx
```

## 📖 Documentation

Each package includes comprehensive documentation and examples. The codebase follows Go best practices and includes extensive test coverage.

## 🤝 Contributing

Contributions are welcome! Please ensure all tests pass and follow the existing code style.

## 📄 License

This project is licensed under the MIT License.