package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	sender("Hey", ch)
	msg := reciver(ch)
	fmt.Println(msg)
}

func sender(a string, ch chan<- string) {
	ch <- a + "!"
}

func reciver(ch <-chan string) string {
	a := <-ch
	return a
}
