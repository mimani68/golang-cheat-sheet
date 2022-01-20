package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("1")
	go func() {
		fmt.Println("2")
	}()
	fmt.Println("3")
	time.Sleep(2 * time.Second)
	// result
	// 1
	// 3
	// 2
}
