package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("hello from a goroutine!")
}

func main() {
	wg.Add(1)
	defer wg.Wait()
	go hello()
}
