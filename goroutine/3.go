package goroutine

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	channel := make(chan string)

	go func() {
		time.Sleep(time.Second * time.Duration(3))
		channel <- "ping"
	}()

	msg := <-channel
	fmt.Println(msg)

}
