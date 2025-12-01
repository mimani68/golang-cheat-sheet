package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var total = 0

func isEvenNumber(number int) bool {
	return number%2 == 0
}

func pip1(input chan int, output chan int) {
	defer wg.Done()
	data := <-input
	if isEvenNumber(data) {
		output <- data
	}
	close(output)
}

func pip2(input chan int, output chan int) {
	defer wg.Done()
	for v := range input {
		sq := v * v
		if sq < 1e5 {
			output <- sq
		}
	}
	close(output)
}

func pip3(input chan int, output chan int) {
	defer wg.Done()
	data, ok := <-input
	if !ok {
		close(output)
		return
	}
	mu.Lock()
	total += data
	output <- total
	mu.Unlock()
	close(output)
}

func main() {
	listOfInput := []int{1, 2, 3, 4, 5, 6, 7, 7, 7, 10, 100}

	for _, value := range listOfInput {
		inputCh := make(chan int)
		middlePipeCh1 := make(chan int, 1)
		middlePipeCh2 := make(chan int, 1)
		middlePipeCh3 := make(chan int, 1)

		wg.Add(3)
		go pip1(inputCh, middlePipeCh1)
		go pip2(middlePipeCh1, middlePipeCh2)
		go pip3(middlePipeCh2, middlePipeCh3)

		inputCh <- value
		close(inputCh)

		if response, ok := <-middlePipeCh3; ok {
			fmt.Printf("Number: %d - Result: %v\n", value, response)
		}

		wg.Wait()
	}
	fmt.Printf("Total number is %d\n", total)

}
