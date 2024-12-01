package main

import (
	"fmt"
	"time"
)

// Stage 1: Generate numbers
func generate(ch chan int, errCh chan error) {
	defer close(errCh)
	for i := 0; i < 10; i++ {
		if i == 7 { // Simulate an error
			errCh <- fmt.Errorf("Error generating number %d", i)
			// return
		}
		fmt.Println("Generate:", i)
		ch <- i
	}
	close(ch) // Signal that no more values will be sent
}

// Stage 2: Process numbers
func process(in chan int, out chan int, errCh chan error) {
	defer close(errCh)
	for num := range in {
		// Simulate some processing time
		time.Sleep(100 * time.Millisecond)
		if num == 2 { // Simulate an error
			errCh <- fmt.Errorf("Error processing number %d", num)
			// return
		}
		// out <- num * 2
		fmt.Println("Process:", num)
		out <- num
	}
	close(out) // Signal that no more values will be sent
}

// Stage 3: Output results
func output(ch chan int, errCh chan error) {
	// defer close(errCh)
	for num := range ch {
		fmt.Println("Output:", num)
	}
}

// Last Stage: Check for errors
func errorhandler(errCh chan error) {
	if err := <-errCh; err != nil {
		fmt.Println("Error:", err)
		// You can also use context to cancel the pipeline here
	}
}

func main() {
	// Create channels
	genCh := make(chan int)
	errGenCh := make(chan error)
	procCh := make(chan int)
	errProCh := make(chan error)

	// Start goroutines
	go generate(genCh, errGenCh)
	go process(genCh, procCh, errProCh)
	go output(procCh, errGenCh)
	go errorhandler(errGenCh)
	go errorhandler(errProCh)

	// Wait for all goroutines to finish (in this case, main will block until output finishes)
	time.Sleep(2 * time.Second) // Adjust the sleep time based on the actual processing time
}
