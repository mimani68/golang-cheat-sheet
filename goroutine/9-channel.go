package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// create a unbuffered channel
	// with a undetermined capacity
	channelInstance := make(chan string)
	go func(ch chan string) {
		max := 3
		min := 1
		randomNumber := rand.Intn(max-min) + min
		time.Sleep(time.Duration(randomNumber/2) * time.Second)
		ch <- fmt.Sprintf("calculated random number: %d", randomNumber)
	}(channelInstance)
	fmt.Println(<-channelInstance)

	// create a buffered channel
	// with a capacity of 2. (fixed capacity)
	// Which means only two time could add new message the channel.
	// Code is blocked since the channel has exceeded its capacity and program reaches deadlock situation.
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "hello world"         // => correct
	ch <- "hello my compatriot" // => fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
