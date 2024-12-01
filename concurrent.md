### Advanced Golang Interview Question:

**Question:**
How would you implement a concurrent program in Go using goroutines and channels to perform a task that involves multiple steps, such as processing a list of items in parallel? Provide an example to illustrate your approach.

### Answer:

To implement a concurrent program in Go using goroutines and channels, you can leverage Go's built-in support for concurrency. Here’s an example of how you can process a list of items in parallel:

#### Example: Processing a List of Items in Parallel

In this example, we will create a program that processes a list of integers in parallel. Each integer will be processed by a separate goroutine, and the results will be collected using a channel.

```go
package main

import (
    "fmt"
    "time"
)

// Function to process an item
func processItem(item int, resultChan chan int) {
    // Simulate some processing time
    time.Sleep(100 * time.Millisecond)
    resultChan <- item * 2 // Send the processed item back through the channel
}

func main() {
    // List of items to process
    items := []int{1, 2, 3, 4, 5}

    // Channel to collect results
    resultChan := make(chan int)

    // Start a goroutine for each item
    for _, item := range items {
        go processItem(item, resultChan)
    }

    // Collect results from the channel
    for i := 0; i < len(items); i++ {
        result := <-resultChan
        fmt.Printf("Processed item: %d\n", result)
    }

    // Close the channel to indicate no more results
    close(resultChan)
}
```

### Explanation:

1. **Goroutines**:
   - The `processItem` function is run as a goroutine for each item in the list. This allows the processing of items to happen concurrently.
   - The `go` keyword is used to start a new goroutine.

2. **Channels**:
   - A channel `resultChan` is created to collect the results from each goroutine.
   - Each goroutine sends its processed item back through the channel using `resultChan <- item * 2`.

3. **Synchronization**:
   - The main goroutine waits for the results by receiving from the channel in a loop.
   - The loop runs `len(items)` times to ensure all results are collected.

4. **Channel Closure**:
   - Although not strictly necessary in this example, closing the channel (`close(resultChan)`) is a good practice to indicate that no more values will be sent.

### Output:

When you run this program, it will output the processed items (each item multiplied by 2) in an order that may vary due to the concurrent nature of the processing:

```
Processed item: 2
Processed item: 4
Processed item: 6
Processed item: 8
Processed item: 10
```

This example demonstrates how to use goroutines and channels to achieve concurrency in Go, allowing for efficient parallel processing of tasks[4][5].