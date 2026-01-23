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
	ch := make(chan string, 2) // NOTE:  increasing channel capacity is not a good solution
	ch <- "hello"
	ch <- "hello world"
	fmt.Println(<-ch)           // <-- FIX: Receive BEFORE third send to free channel space
	ch <- "hello my compatriot" // => fatal error: all goroutines are asleep - deadlock! without free capacity
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
