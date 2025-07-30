# stdx

A comprehensive standard library extension for Go, providing commonly used data structures and functional programming utilities.

## Packages

### Core Data Structures

- **`listx`** - List interfaces and implementations
  - `listx/linked` - Linked list implementation
  - `listx/slices` - Slice-based implementation  
  - `listx/hash` - Hash-based implementation
  - Includes: List, Deque, Stack, Queue interfaces

- **`mapx`** - Map interfaces and implementations
  - `mapx/hashmap` - Hash map implementation
  - `mapx/concurrentmap` - Thread-safe concurrent map

- **`setx`** - Set interfaces and implementations
  - `setx/hashset` - Hash set implementation
  - `setx/concurrentset` - Thread-safe concurrent set

### Functional Programming

- **`option`** - Rust-inspired Option[T] type for handling nullable values
- **`result`** - Rust-inspired Result[T, E] type for error handling

## Features

- **Generic Types**: All data structures use Go 1.18+ generics for type safety
- **Multiple Implementations**: Different backing implementations for different use cases
- **Comprehensive Testing**: Full test coverage with factory patterns
- **Rust-Inspired**: Functional programming constructs inspired by Rust
- **Thread Safety**: Concurrent implementations where needed

## Usage

```go
package main

import (
    "fmt"
    "github.com/gosuda/stdx/listx/linked"
    "github.com/gosuda/stdx/option"
    "github.com/gosuda/stdx/result"
)

func main() {
    // Using List
    list := linked.New[int]()
    list.Add(1)
    list.Add(2)
    list.Add(3)
    
    // Using Option
    opt := option.Some(42)
    value := opt.UnwrapOr(0)
    
    // Using Result
    res := result.Ok[int, string](100)
    if res.IsOk() {
        fmt.Println("Success:", res.Unwrap())
    }
}
```