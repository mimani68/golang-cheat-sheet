## Question:
How can you use goroutines and channels to implement a concurrent task that involves multiple stages, ensuring that each stage completes before the next one starts, and how can you handle errors and panics in such a setup?

## Answer:

To implement a concurrent task with multiple stages using goroutines and channels, you can use a pipeline approach where each stage communicates with the next through channels. Here’s how you can do it:

### Using Channels for Pipeline Communication

You can create a series of goroutines, each representing a stage in the pipeline. Each goroutine receives input from a channel, processes it, and then sends the output to another channel for the next stage.

### Handling Errors and Panics

To handle errors and panics, you can use additional channels to communicate error messages or use a context to cancel the pipeline if an error occurs.

### Example

Here is an example of a pipeline with three stages: `generate`, `process`, and `output`. Each stage runs in a separate goroutine and communicates through channels.

```go
package main

import (
    "fmt"
    "time"
)

// Stage 1: Generate numbers
func generate(ch chan int) {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch) // Signal that no more values will be sent
}

// Stage 2: Process numbers
func process(in chan int, out chan int) {
    for num := range in {
        // Simulate some processing time
        time.Sleep(100 * time.Millisecond)
        out <- num * 2
    }
    close(out) // Signal that no more values will be sent
}

// Stage 3: Output results
func output(ch chan int) {
    for num := range ch {
        fmt.Println(num)
    }
}

func main() {
    // Create channels
    genCh := make(chan int)
    procCh := make(chan int)

    // Start goroutines
    go generate(genCh)
    go process(genCh, procCh)
    go output(procCh)

    // Wait for all goroutines to finish (in this case, main will block until output finishes)
    time.Sleep(2 * time.Second) // Adjust the sleep time based on the actual processing time
}
```

### Handling Errors

To handle errors, you can add error channels and check for errors in each stage.

```go
package main

import (
    "fmt"
    "time"
)

// Stage 1: Generate numbers
func generate(ch chan int, errCh chan error) {
    defer close(errCh)
    for i := 0; i < 10; i++ {
        if i == 5 { // Simulate an error
            errCh <- fmt.Errorf("Error generating number %d", i)
            return
        }
        ch <- i
    }
    close(ch) // Signal that no more values will be sent
}

// Stage 2: Process numbers
func process(in chan int, out chan int, errCh chan error) {
    defer close(errCh)
    for num := range in {
        // Simulate some processing time
        time.Sleep(100 * time.Millisecond)
        if num == 6 { // Simulate an error
            errCh <- fmt.Errorf("Error processing number %d", num)
            return
        }
        out <- num * 2
    }
    close(out) // Signal that no more values will be sent
}

// Stage 3: Output results
func output(ch chan int, errCh chan error) {
    defer close(errCh)
    for num := range ch {
        fmt.Println(num)
    }
}

func main() {
    // Create channels
    genCh := make(chan int)
    procCh := make(chan int)
    errCh := make(chan error)

    // Start goroutines
    go generate(genCh, errCh)
    go process(genCh, procCh, errCh)
    go output(procCh, errCh)

    // Check for errors
    go func() {
        if err := <-errCh; err != nil {
            fmt.Println("Error:", err)
            // You can also use context to cancel the pipeline here
        }
    }()

    // Wait for all goroutines to finish (in this case, main will block until output finishes)
    time.Sleep(2 * time.Second) // Adjust the sleep time based on the actual processing time
}
```

### Handling Panics

To handle panics, you can use `defer` and `recover` within each goroutine.

```go
package main

import (
    "fmt"
    "time"
)

// Stage 1: Generate numbers
func generate(ch chan int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in generate:", r)
        }
        close(ch) // Signal that no more values will be sent
    }()
    for i := 0; i < 10; i++ {
        if i == 5 { // Simulate a panic
            panic("Panic generating number")
        }
        ch <- i
    }
}

// Stage 2: Process numbers
func process(in chan int, out chan int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in process:", r)
        }
        close(out) // Signal that no more values will be sent
    }()
    for num := range in {
        // Simulate some processing time
        time.Sleep(100 * time.Millisecond)
        if num == 6 { // Simulate a panic
            panic("Panic processing number")
        }
        out <- num * 2
    }
}

// Stage 3: Output results
func output(ch chan int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in output:", r)
        }
    }()
    for num := range ch {
        fmt.Println(num)
    }
}

func main() {
    // Create channels
    genCh := make(chan int)
    procCh := make(chan int)

    // Start goroutines
    go generate(genCh)
    go process(genCh, procCh)
    go output(procCh)

    // Wait for all goroutines to finish (in this case, main will block until output finishes)
    time.Sleep(2 * time.Second) // Adjust the sleep time based on the actual processing time
}
```

This example demonstrates how to use goroutines and channels to implement a multi-stage pipeline, handle errors, and recover from panics in each stage.