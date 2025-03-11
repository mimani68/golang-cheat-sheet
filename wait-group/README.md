Here's a quick overview of Golang WaitGroups and some common use cases:

WaitGroup Overview:
- A WaitGroup waits for a collection of goroutines to finish executing
- It's part of the sync package
- Has three main methods:
  - Add(delta int) - Adds delta to the WaitGroup counter
  - Done() - Decrements the WaitGroup counter by 1 
  - Wait() - Blocks until the WaitGroup counter is 0

Key points:
- Used to synchronize multiple goroutines
- Allows main goroutine to wait for other goroutines to complete
- Counter starts at 0 and is incremented by Add() calls
- Counter is decremented by Done() calls
- Wait() blocks until counter reaches 0

Common Use Cases:

1. Waiting for multiple goroutines to complete:

```go
var wg sync.WaitGroup
for i := 0; i < 5; i++ {
  wg.Add(1)
  go func() {
    defer wg.Done()
    // Do work
  }()
}
wg.Wait() // Wait for all goroutines to finish
```

2. Parallel processing of data:

```go
var wg sync.WaitGroup
data := []int{1,2,3,4,5}
for _, item := range data {
  wg.Add(1)
  go func(i int) {
    defer wg.Done()
    processItem(i)
  }(item)
}
wg.Wait()
```

3. Coordinating startup of multiple services:

```go
var wg sync.WaitGroup
wg.Add(3)
go startDatabase(&wg)
go startWebServer(&wg) 
go startCacheService(&wg)
wg.Wait() // Wait for all services to start
```

4. Fan-out/fan-in patterns:

```go
func fanOut(tasks []Task) {
  var wg sync.WaitGroup
  for _, task := range tasks {
    wg.Add(1)
    go func(t Task) {
      defer wg.Done()
      processTask(t)
    }(task)
  }
  wg.Wait()
}
```

5. Graceful shutdown:

```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
  defer wg.Done()
  // Run service
}()

// On shutdown signal
wg.Wait() // Wait for service to finish
```

WaitGroups are useful whenever you need to wait for multiple concurrent operations to complete before proceeding. They help coordinate goroutines and ensure proper synchronization in concurrent Go programs.