package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var total = 0

// No changes needed here
func isEvenNumber(number int) bool {
	return number%2 == 0
}

// OPTIMIZED: Uses a for-range loop to process all values from the input channel.
func pip1(input <-chan int, output chan<- int) {
	defer wg.Done()
	// Process all data from the input channel until it's closed.
	for v := range input {
		if isEvenNumber(v) {
			output <- v
		}
	}
	// After the loop, close the output channel to signal the next stage that we're done.
	close(output)
}

// OPTIMIZED: This function was already mostly correct.
func pip2(input <-chan int, output chan<- int) {
	defer wg.Done()
	for v := range input {
		sq := v * v
		if sq < 1e5 {
			output <- sq
		}
	}
	close(output)
}

// OPTIMIZED: Uses a for-range loop and a mutex to safely update the shared total.
func pip3(input <-chan int) { // No output channel needed, it just updates the total.
	defer wg.Done()
	// Process all data from the input channel until it's closed.
	for v := range input {
		mu.Lock()
		total += v
		mu.Unlock()
	}
}

func main() {
	listOfInput := []int{1, 2, 3, 4, 5, 6, 7, 7, 7, 10, 100, 1000, 2000, 3000}

	// OPTIMIZED: Create channels and goroutines ONCE.
	inputCh := make(chan int)
	middlePipeCh1 := make(chan int)
	middlePipeCh2 := make(chan int)

	wg.Add(3)
	go pip1(inputCh, middlePipeCh1)
	go pip2(middlePipeCh1, middlePipeCh2)
	go pip3(middlePipeCh2) // pip3 now just consumes data and updates the total.

	// Use a separate goroutine to feed data into the pipeline to avoid a deadlock.
	// The main goroutine can then wait for the pipeline to finish.
	go func() {
		for _, value := range listOfInput {
			inputCh <- value
		}
		close(inputCh) // Signal the start of the pipeline that no more data is coming.
	}()

	wg.Wait() // Wait for all three pipeline stages to complete.

	fmt.Printf("Final total is %d\n", total)
}
