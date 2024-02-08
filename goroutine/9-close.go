package main

import "fmt"

func main() {

	// queue := make(chan string, 1) // fatal error: all goroutines are asleep - deadlock!
	// queue := make(chan string, 2) // OK
	queue := make(chan string, 3) // OK
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

// OUTPUT
// one
// two
