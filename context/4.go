package main

import (
	"context"
	"fmt"
	"time"
)

func operation(ctx context.Context) {
	// We use a similar pattern to the HTTP server
	// that we saw in the earlier example
	select {
	case <-time.After(10 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("halted operation2")
	}
}

func main() {
	// Create a new context
	ctx := context.Background()

	// Create a new context, with its cancellation function
	// from the original context
	ctx, cancel := context.WithCancel(ctx)

	// Run operation
	operation(ctx)

	// cancel next operation
	if true {
		cancel()
	}

	// Run operation
	operation(ctx)
}
