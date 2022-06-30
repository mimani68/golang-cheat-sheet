package main

import (
	"context"
	"fmt"
	"time"
)

// heavy process function
func fn(ctx context.Context, c chan context.Context) {
	// ctx1, _ := context.WithTimeout(ctx, 2*time.Second)
	time.Sleep(1 * time.Second)
	ctx.Done()
	c <- ctx
	// return ctx.Done()
}

func main() {
	c := make(chan context.Context)
	ctx := context.Background()
	context.WithValue(ctx, "rid", "9f44f4ba-f84e-11ec-9c6d-6fc79d5e25bf")
	go fn(ctx, c)
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case <-c:
			fmt.Println("Context finished using channel")
		case <-ctx.Done():
			fmt.Println("Context finished")
			return
		default:
			fmt.Println("-")
		}
	}
	fmt.Println(ctx.Value("rid"))
}
