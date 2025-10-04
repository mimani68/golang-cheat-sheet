## Golang generator

In Go, there is no `yield` keyword like in Python or other languages. However, you can achieve similar functionality using **channels** and **goroutines** to create a generator-like pattern. Here are examples demonstrating how to mimic the behavior of `yield` in Go:

---

### Example 1: Basic "Generator" with Channels
This example creates a channel that "yields" numbers from 1 to 5:

```go
package main

import "fmt"

// Define a function that returns a channel of integers.
func numberGenerator() <-chan int {
    ch := make(chan int)
    go func() {
        for i := 1; i <= 5; i++ {
            ch <- i // "Yield" the value by sending to the channel
        }
        close(ch) // Close the channel after sending all values
    }()
    return ch
}

func main() {
    ch := numberGenerator()
    for num := range ch { // Receive values until the channel is closed
        fmt.Println(num)
    }
}
```

**Output:**
```
1
2
3
4
5
```

---

### Example 2: Fibonacci Sequence Generator
This example generates Fibonacci numbers indefinitely until a stop signal:

```go
package main

import "fmt"

// Define a Fibonacci generator that stops when the 'stop' channel is closed.
func fibonacciGenerator(stop <-chan struct{}) <-chan int {
    ch := make(chan int)
    go func() {
        a, b := 0, 1
        for {
            select {
            case <-stop:
                close(ch)
                return
            case ch <- a:
                a, b = b, a+b
            }
        }
    }()
    return ch
}

func main() {
    stop := make(chan struct{})
    fib := fibonacciGenerator(stop)

    // Print first 10 Fibonacci numbers
    for i := 0; i < 10; i++ {
        fmt.Println(<-fib)
    }
    close(stop) // Signal to stop the generator
}
```

**Output:**
```
0
1
1
2
3
5
8
13
21
34
```


## Generator with Stop Condition

```go
package main

import "fmt"

func rangeGenerator(start, end int) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for i := start; i <= end; i++ {
            ch <- i
        }
    }()
    return ch
}

func main() {
    for num := range rangeGenerator(1, 5) {
        fmt.Println(num)
    }
}
```

## Using Function Closures (Alternative Pattern)

```go
package main

import "fmt"

func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    next := counter()
    fmt.Println(next()) // 1
    fmt.Println(next()) // 2
    fmt.Println(next()) // 3
}
```

## Iterator Pattern with State

```go
package main

import "fmt"

type Iterator struct {
    current int
    max     int
}

func NewIterator(max int) *Iterator {
    return &Iterator{current: 0, max: max}
}

func (it *Iterator) Next() (int, bool) {
    if it.current >= it.max {
        return 0, false
    }
    value := it.current
    it.current++
    return value, true
}

func main() {
    iter := NewIterator(5)
    for {
        val, ok := iter.Next()
        if !ok {
            break
        }
        fmt.Println(val)
    }
}
```

## Infinite Generator with Context for Cancellation

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func infiniteSequence(ctx context.Context) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        i := 0
        for {
            select {
            case <-ctx.Done():
                return
            case ch <- i:
                i++
            }
        }
    }()
    return ch
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    for num := range infiniteSequence(ctx) {
        fmt.Println(num)
        time.Sleep(500 * time.Millisecond)
    }
}
```
---

### Key Concepts:
1. **Channels**: Used to communicate values between goroutines. The generator sends values (`ch <- value`), and the receiver reads them (`<-ch`).
2. **Goroutines**: The generator runs in its own goroutine to avoid blocking the main program.
3. **Closing Channels**: Closing the channel (`close(ch)`) signals that no more values will be sent, allowing the receiver to exit the loop.

---

### Why No `yield` Keyword?
Go does not have a `yield` keyword because its concurrency model is built around **channels** and **goroutines**. This design choice enforces a clear separation of concerns and avoids the complexity of pausing/resuming functions (like in Python's `yield`).

---

### Use Cases for "Yield"-Like Behavior:
- Generating sequences (e.g., numbers, Fibonacci).
- Lazy evaluation of data (process elements as needed).
- Streaming large datasets without loading everything into memory.

By using channels and goroutines, Go provides a flexible way to achieve similar patterns to `yield`, while staying true to its design principles.

