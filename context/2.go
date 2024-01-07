package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, "key", "value")
	go retrieveValues(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("Took too long!")
	}
	time.Sleep(10 * time.Second)
}

func retrieveValues(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		default:
			value := ctx.Value("key")
			fmt.Println(value)
		}
		time.Sleep(4 * time.Second)
	}
}
