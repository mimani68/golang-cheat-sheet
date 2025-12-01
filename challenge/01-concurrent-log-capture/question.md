Here’s a Go code challenge that focuses on high-volume disk I/O using built-in packages and concurrency. It’s designed to test your understanding of goroutines, channels, file operations, and synchronization.

---

## 🧠 Go Code Challenge: Concurrent Disk Logger

### 📝 Task Description
Build a concurrent logging system in Go that handles high-frequency write and read requests to a log file. The system should:
- Accept log messages from multiple sources (simulated with goroutines).
- Write each message to disk using built-in packages (`os`, `bufio`, etc.).
- Periodically read the log file to count the number of entries.
- Use channels and goroutines to manage concurrency safely.

---

### 📌 Requirements
1. **Concurrency**: Use goroutines to simulate multiple sources sending log messages.
2. **Channels**: Use channels to queue log messages.
3. **Disk I/O**: Use `os` and `bufio` to write/read from a file.
4. **Synchronization**: Ensure safe access to the file (e.g., using `sync.Mutex`).
5. **Performance**: Handle at least 1000 log messages per second.

---

### 🧪 Example Behavior
- Start 5 goroutines, each sending 200 log messages.
- A logger goroutine writes messages to `logs.txt`.
- A monitor goroutine reads the file every 2 seconds and prints the number of log entries.

---

### 🚀 Starter Code Skeleton

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "sync"
    "time"
)

const logFile = "logs.txt"

func main() {
    logChan := make(chan string, 1000)
    var mu sync.Mutex

    // Start logger goroutine
    go func() {
        file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            panic(err)
        }
        defer file.Close()
        writer := bufio.NewWriter(file)

        for msg := range logChan {
            mu.Lock()
            fmt.Fprintln(writer, msg)
            writer.Flush()
            mu.Unlock()
        }
    }()

    // Start monitor goroutine
    go func() {
        for {
            time.Sleep(2 * time.Second)
            mu.Lock()
            count := countLines(logFile)
            mu.Unlock()
            fmt.Printf("Log entries: %d\n", count)
        }
    }()

    // Simulate multiple sources
    for i := 0; i < 5; i++ {
        go func(id int) {
            for j := 0; j < 200; j++ {
                logChan <- fmt.Sprintf("Source %d: message %d", id, j)
                time.Sleep(5 * time.Millisecond)
            }
        }(i)
    }

    time.Sleep(10 * time.Second)
    close(logChan)
}

func countLines(filename string) int {
    file, err := os.Open(filename)
    if err != nil {
        return 0
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    count := 0
    for scanner.Scan() {
        count++
    }
    return count
}
```

---

Would you like to extend this challenge with error handling, performance metrics, or log rotation?