package main

import "fmt"

func main() {

	queue := make(chan string, 3)
	close(queue)
	queue <- "one"
	queue <- "two"

	for elem := range queue {
		fmt.Println(elem)
	}
}

// OUTPUT
// panic: send on closed channel
