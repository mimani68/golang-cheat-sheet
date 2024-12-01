package main

import "fmt"

func main() {

	ch := make(chan int, 4)

	go func() {
		fmt.Println("1")
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		ch <- 6
		fmt.Println("2")
		close(ch)
	}()

	go func() {
		fmt.Println("3")
		for num := range ch {
			fmt.Println("4")
			fmt.Println(num)
		}
	}()
}
