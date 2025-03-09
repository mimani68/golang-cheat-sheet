Certainly! I'll provide an example of a race condition in Go and then show how to resolve it using best practices. Here's an example of a race condition:

```go
package main

import (
    "fmt"
    "sync"
)

var counter int

func increment(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 1000; i++ {
        counter++
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go increment(&wg)
    go increment(&wg)

    wg.Wait()
    fmt.Println("Final counter value:", counter)
}
```

In this example, we have a shared `counter` variable that two goroutines are trying to increment concurrently. This creates a race condition because the increment operation (`counter++`) is not atomic, and both goroutines may read and write to the counter at the same time, leading to unpredictable results.

To resolve this issue, we can use several best practices:

1. Use a mutex to protect the shared resource:

```go
package main

import (
    "fmt"
    "sync"
)

var (
    counter int
    mutex   sync.Mutex
)

func increment(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 1000; i++ {
        mutex.Lock()
        counter++
        mutex.Unlock()
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go increment(&wg)
    go increment(&wg)

    wg.Wait()
    fmt.Println("Final counter value:", counter)
}
```

2. Use atomic operations for simple counters:

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

var counter int64

func increment(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 1000; i++ {
        atomic.AddInt64(&counter, 1)
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go increment(&wg)
    go increment(&wg)

    wg.Wait()
    fmt.Println("Final counter value:", atomic.LoadInt64(&counter))
}
```

3. Use channels for communication between goroutines:

```go
package main

import (
    "fmt"
)

func increment(ch chan<- int) {
    for i := 0; i < 1000; i++ {
        ch <- 1
    }
    close(ch)
}

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go increment(ch1)
    go increment(ch2)

    counter := 0
    for {
        select {
        case _, ok := <-ch1:
            if !ok {
                ch1 = nil
            } else {
                counter++
            }
        case _, ok := <-ch2:
            if !ok {
                ch2 = nil
            } else {
                counter++
            }
        }
        if ch1 == nil && ch2 == nil {
            break
        }
    }

    fmt.Println("Final counter value:", counter)
}
```

These solutions resolve the race condition by ensuring that access to the shared resource (the counter) is properly synchronized. Here are some additional best practices to avoid race conditions:

1. Use the `-race` flag when testing and building your Go programs to detect race conditions:
   ```
   go test -race ./...
   go build -race
   go run -race main.go
   ```

2. Minimize shared state between goroutines.

3. Prefer using channels for communication between goroutines instead of shared memory and locks.

4. When using locks, keep the critical sections as small as possible to reduce contention.

5. Consider using sync.Once for one-time initialization that needs to be thread-safe.

6. Use tools like go vet and static analysis tools to catch potential issues early.

By following these best practices, you can write concurrent Go programs that are less prone to race conditions and more reliable.