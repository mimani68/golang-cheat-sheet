package main

import (
	"context"
	"fmt"
	"time"
)

func operation(ctx context.Context, opsName string) {
	// We use a similar pattern to the HTTP server
	// that we saw in the earlier example
	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Printf("[%s] done\n", opsName)
	case <-ctx.Done():
		fmt.Printf("[%s] halted\n", opsName)
	}
}

func main() {
	// Create a new context
	ctx := context.Background()

	// Create a new context, with its cancellation function
	// from the original context
	ctx, cancel := context.WithCancel(ctx)


	// Run operation one
	operation(ctx, "operation-one")

	// cancel next operation
	if true {
		cancel()
	}

	// Run operation two
	operation(ctx, "operation-two")
}
