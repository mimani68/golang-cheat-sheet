### Question:
How can you ensure concurrent consumption of data from a buffered channel to avoid deadlocks, and what are the key differences between buffered and unbuffered channels in this context?

### Answer:

To ensure concurrent consumption of data from a buffered channel and avoid deadlocks, you need to separate the producer and consumer into different goroutines. Here’s why and how:

#### Buffered vs Unbuffered Channels:

- **Buffered Channels**: These allow sending and receiving to occur independently up to the limit of the buffer. If the buffer is full, the sender will block until space is available. If the buffer is empty, the receiver will block until data is available. Buffered channels can help prevent deadlocks by allowing producers to send data even if the consumers are not immediately ready to receive it.

- **Unbuffered Channels**: These require both send and receive operations to be ready at the same time. If one side is not ready, the operation will block until the other side is ready.

#### Avoiding Deadlocks:
To avoid deadlocks when using a buffered channel, ensure that the producer and consumer run in separate goroutines. Here is an example:

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 4) // Create a buffered channel with capacity 4

    // Start a goroutine to send values into the channel
    go func() {
        for i := 1; i <= 5; i++ {
            ch <- i
        }
        close(ch) // Close the channel to indicate no more data
    }()

    // Consume values from the channel in the main goroutine
    for num := range ch {
        fmt.Println(num)
    }
}
```

### Explanation:
- **Separate Goroutines**: By using `go func() { ... }()` to send values into the channel, we ensure that the sending operation runs in a separate goroutine from the main goroutine, which consumes the values. This prevents the channel from getting full and blocking further sends, thus avoiding a deadlock[2][3].
- **Concurrent Consumption**: The `range` loop in the main goroutine continuously reads from the channel, ensuring that the channel does not get full and block the sending goroutine.
- **Channel Closure**: The `close(ch)` statement after sending all values indicates that no more data will be sent, allowing the range loop to exit cleanly.

### Example Output:
This code will output all the numbers from 1 to 5 without encountering a deadlock.

```plaintext
1
2
3
4
5
```

### Incorrect Scenario Leading to Deadlock:
If you remove the anonymous goroutine and send integers directly within the `main` function, it will lead to a deadlock because the channel will get full before it gets consumed, blocking further operations[2].

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 4)
    ch <- 1
    ch <- 2
    ch <- 3
    ch <- 4
    ch <- 5 // This will block and cause a deadlock
    close(ch)
    for num := range ch {
        fmt.Println(num)
    }
}
```

This scenario will result in a deadlock error:

```
fatal error: all goroutines are asleep - deadlock!
```