package goroutine

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channelInstance := make(chan string)
	go func(ch chan string) {
		max := 3
		min := 1
		randomNumber := rand.Intn(max-min) + min
		time.Sleep(time.Duration(randomNumber/2) * time.Second)
		ch <- fmt.Sprintf("salam %d", randomNumber)
	}(channelInstance)
	fmt.Println(<-channelInstance)
}
