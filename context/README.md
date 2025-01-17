Let's break down the common scenarios for using `context.Context` in Go, focusing on `context.Background()`, `context.TODO()`, `context.WithCancel()`, `context.WithDeadline()`, `context.WithTimeout()`, and `context.WithValue()`.  I'll explain their intended use, differences, and best practices.

**1. `context.Background()`**

*   **Purpose:**  The root of every context tree.  It's an empty context; it has no values, no deadlines, and is never canceled.
*   **Use Case:**
    *   **Main Function/Initialization:**  The very first context you create in your program, typically in your `main()` function or during service startup.  This is the ancestor of all other contexts.
    *   **Top-Level Request Handlers:** When you receive a new request (e.g., in an HTTP server), you often start with `context.Background()` and then derive more specific contexts (like those with timeouts).
    *   **Tests:** When unit testing functions that accept a context, `context.Background()` provides a safe, neutral starting point.
*   **Key Feature:** Never canceled, no deadlines, no values. It is the empty context.
* **Example**

```go
package main

import (
        "context"
        "fmt"
        "time"
)

func main() {
        // Start with a background context
        ctx := context.Background()

        // Pass it to a function
        doSomething(ctx)
}

func doSomething(ctx context.Context) {
        // Simulate some work
        for i := 0; i < 5; i++ {
                select {
                case <-ctx.Done():
                        fmt.Println("Context cancelled, stopping work")
                        return
                default:
                        fmt.Println("Doing some work...", i)
                        time.Sleep(500 * time.Millisecond)
                }
        }
}
```

**2. `context.TODO()`**

*   **Purpose:** A placeholder context.  It signifies that you *know* you need a context, but you haven't yet decided which one is appropriate.  It's a temporary marker.
*   **Use Case:**
    *   **During Development:** When you're writing code and realize a function should accept a context, but you're not ready to determine the exact context hierarchy or cancellation strategy.
    *   **Unclear Context:** When it is not clear which context to use.
*   **Key Feature:**  Functionally identical to `context.Background()`. The difference is purely in the intent it conveys to other developers (and your future self).
*   **Best Practice:**  Replace `context.TODO()` with a more specific context (`context.Background()`, `context.WithCancel()`, etc.) as soon as you understand how the context should be used.  Don't leave `context.TODO()` in production code.  Think of it like a `// TODO:` comment, but for contexts.
* **Example**
```go
package main

import (
        "context"
        "fmt"
        "time"
)

func main() {
    // Use context.TODO() as a placeholder
    ctx := context.TODO()

    // Call a function that needs a context
    result, err := fetchData(ctx, "some_data_id")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Result:", result)
}

func fetchData(ctx context.Context, dataID string) (string, error) {
    // Simulate fetching data with a potential timeout
    select {
    case <-time.After(2 * time.Second): // Simulate a 2-second operation
        return "Data for " + dataID, nil
    case <-ctx.Done():
        return "", ctx.Err() // Return the context's error (e.g., canceled or deadline exceeded)
    }
}
```

**3. `context.WithCancel(parent Context)`**

*   **Purpose:** Creates a new context that can be explicitly canceled.  Returns the derived context and a `cancel` function.  Calling the `cancel` function signals to all child contexts that they should stop their work.
*   **Use Case:**
    *   **Goroutine Management:**  When you launch a goroutine that performs a long-running operation, use `context.WithCancel()` to create a context that you can cancel if the operation needs to be stopped prematurely (e.g., the user closes a connection, or another part of the system requests a shutdown).
    *   **Resource Cleanup:**  When a request is canceled, you often want to release resources (close connections, stop timers, etc.).  The `cancel` function triggers this cleanup.
    *   **Error Propagation (Optional):** You can optionally pass an error to the cancel function. This error will then be available via `ctx.Err()`. This can be used to explain *why* the context was canceled.
*   **Key Feature:** Provides a `cancel()` function to explicitly stop operations.  Propagates cancellation to child contexts.
* **Example**

```go
package main

import (
        "context"
        "fmt"
        "time"
)

func main() {
        // Create a context with cancel
        ctx, cancel := context.WithCancel(context.Background())

        // Start a goroutine that uses the context
        go worker(ctx)

        // Wait for a while, then cancel the context
        time.Sleep(3 * time.Second)
        fmt.Println("Cancelling context...")
        cancel()

        // Wait for the goroutine to finish (or for a timeout)
        time.Sleep(2 * time.Second)
        fmt.Println("Main goroutine exiting")
}

func worker(ctx context.Context) {
        for i := 0; ; i++ {
                select {
                case <-ctx.Done():
                        fmt.Println("Worker: Context cancelled, stopping")
                        return
                default:
                        fmt.Println("Worker: Doing some work...", i)
                        time.Sleep(1 * time.Second)
                }
        }
}
```

**4. `context.WithDeadline(parent Context, d time.Time)`**

*   **Purpose:**  Creates a new context that will be automatically canceled at a specific time (`d`).
*   **Use Case:**
    *   **Time-Bound Operations:**  When an operation must complete by a certain deadline. For example, you might have a service that needs to respond to a request within 5 seconds.  If the deadline is exceeded, the context is canceled, and you can return an error to the client.
*   **Key Feature:** Automatic cancellation at a fixed point in time.
* **Example**

```go
package main

import (
        "context"
        "fmt"
        "time"
)

func main() {
        // Set a deadline for 2 seconds from now
        deadline := time.Now().Add(2 * time.Second)
        ctx, cancel := context.WithDeadline(context.Background(), deadline)
        defer cancel() // Important: Always call cancel, even if the deadline passes!

        // Simulate work that might take longer than the deadline
        err := doSomethingWithDeadline(ctx)
        if err != nil {
                fmt.Println("Error:", err)
        }
}

func doSomethingWithDeadline(ctx context.Context) error {
        select {
        case <-time.After(3 * time.Second): // Simulate work taking 3 seconds
                fmt.Println("Work completed successfully")
                return nil
        case <-ctx.Done():
                return ctx.Err() // Returns context.DeadlineExceeded
        }
}
```

**5. `context.WithTimeout(parent Context, timeout time.Duration)`**

*   **Purpose:** Creates a new context that will be automatically canceled after a specified duration (`timeout`).  This is a convenience function that's equivalent to using `context.WithDeadline` with `time.Now().Add(timeout)`.
*   **Use Case:**  Same as `context.WithDeadline()`, but expressed in terms of a duration rather than an absolute time. This is generally preferred for readability.
*   **Key Feature:** Automatic cancellation after a specified time interval.
* **Example**

```go
package main

import (
        "context"
        "fmt"
        "time"
)

func main() {
        // Create a context with a 2-second timeout
        ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
        defer cancel() // Good practice: Always call cancel, even if the timeout is reached.

        // Do some work that might take longer than the timeout
        err := doSomethingWithTimeout(ctx)
        if err != nil {
                fmt.Println("Error:", err)
        }
}

func doSomethingWithTimeout(ctx context.Context) error {
    select {
    case <-time.After(3 * time.Second): // Simulate work taking 3 seconds
        fmt.Println("Work completed (but timed out)")
        return nil
    case <-ctx.Done():
        return ctx.Err() // Returns context.DeadlineExceeded
    }
}
```

**6. `context.WithValue(parent Context, key, val interface{})`**

*   **Purpose:** Creates a new context that carries a key-value pair.  This allows you to associate data with a context.
*   **Use Case:**
    *   **Request-Scoped Data:**  Passing information that's specific to a particular request (e.g., user ID, request ID, authentication token).
    *   **Middleware:**  In web frameworks, middleware can use `context.WithValue()` to add data to the context that's then available to downstream handlers.
*   **Key Feature:**  Stores arbitrary key-value pairs.  Values are retrieved using `ctx.Value(key)`.
*   **Best Practices:**
    *   **Use Defined Types for Keys:** Don't use built-in types like `string` or `int` for context keys.  Create custom types (usually unexported) to avoid key collisions between different packages.
    *   **Minimize Use:**  Avoid overusing `context.WithValue()`.  It can make code harder to understand and debug.  If you find yourself storing many values in the context, it might be a sign that you should refactor your code to pass data more explicitly.
    *   **Immutable Data:** The data stored in the context should be immutable.  If you need to modify data, create a new context with the updated value.
    *   **Request-Scoped Only:** Only for data that's directly related to the processing of a single request.  Don't use it for global configuration or shared state.
* **Example**
```go
package main

import (
    "context"
    "fmt"
)

// Define a custom type for the context key
type keyType int

const userIDKey keyType = 0

func main() {
    // Create a context with a user ID
    ctx := context.WithValue(context.Background(), userIDKey, 123)

    // Pass the context to a function
    processRequest(ctx)
}

func processRequest(ctx context.Context) {
    // Retrieve the user ID from the context
    userID, ok := ctx.Value(userIDKey).(int)
    if !ok {
        fmt.Println("User ID not found in context")
        return
    }

    fmt.Println("Processing request for user ID:", userID)
}
```

**Key Differences and Summary Table**

| Context Type             | Purpose                                    | Cancellation                   | Deadline          | Values               | Best Use Cases                                                                                                                            |
| ------------------------- | ------------------------------------------ | ------------------------------ | ----------------- | --------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `context.Background()`   | Root context, empty.                       | Never.                        | None.             | None.                | `main()` function, top-level handlers, tests.                                                                                           |
| `context.TODO()`         | Placeholder, functionally same as Background. | Never.                        | None.             | None.                | During development, when the context's role is not yet clear.  *Replace ASAP*.                                                        |
| `context.WithCancel()`   | Creates a cancellable context.               | Explicitly via `cancel()`.     | None.             | Inherits from parent. | Goroutine management, resource cleanup, stopping long-running operations.                                                             |
| `context.WithDeadline()` | Creates a context with a deadline.         | Automatically at deadline.     | Specified time.  | Inherits from parent. | Operations that must complete by a specific time.                                                                                     |
| `context.WithTimeout()`  | Creates a context with a timeout.          | Automatically after timeout.   | Duration.        | Inherits from parent. | Operations that must complete within a time interval. (Preferred over `WithDeadline` for readability.)                               |
| `context.WithValue()`   | Creates a context with a key-value pair.    | Inherits from parent.          | Inherits from parent. | Key-value pairs.    | Passing request-scoped data (user ID, request ID, etc.). Use sparingly and with defined key types.                                     |

**Important Considerations**

*   **Context Hierarchy:** Contexts form a tree.  When a parent context is canceled, all of its child contexts are also canceled.
*   **`defer cancel()`:** When using `context.WithCancel()`, `context.WithDeadline()`, or `context.WithTimeout()`, it's *crucial* to call the `cancel()` function, even if the context is already canceled (e.g., the deadline passes).  This ensures that any resources associated with the context are released.  Use `defer cancel()` to guarantee this happens.
*   **`ctx.Done()`:**  The `ctx.Done()` channel is closed when the context is canceled (either explicitly or due to a deadline/timeout).  Use `select` with `ctx.Done()` to detect cancellation.
* **Error with cancellation:** With `context.WithCancelCause(parent Context)` you can pass an error that will be show with `context.Cause(ctx)`
*   **`ctx.Err()`:** Returns an error indicating *why* the context was canceled.  This will be `context.Canceled` if the context was explicitly canceled, `context.DeadlineExceeded` if the deadline was reached, or the error provided to `cancel` function if one is called with a non-nil error.
*  **Do not store Contexts inside a struct type;** instead, pass a Context explicitly to each function that needs it. The Context should be the first parameter, typically named ctx.

By understanding these different context types and their intended uses, you can write more robust, reliable, and maintainable Go code, especially when dealing with concurrency and time-sensitive operations. Remember to use `context.TODO()` sparingly and replace it with a more appropriate context type as soon as possible.  Always `defer cancel()` when using contexts with cancellation capabilities.  And use `context.WithValue()` judiciously, preferring to pass data explicitly whenever possible.