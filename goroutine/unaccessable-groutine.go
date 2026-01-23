package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 4)

	go func() {
		fmt.Println("Start")
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		ch <- 6
		fmt.Println("End")
		close(ch)
	}()

	go func() {
		fmt.Println("Inner Goroutine Start")
		for num := range ch {
			fmt.Println(num)
		}
	}()

	time.Sleep(time.Second * 2)
}
