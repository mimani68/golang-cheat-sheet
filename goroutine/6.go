package goroutine

import (
	"fmt"
	"time"
)

func main() {
	theMine := [3]string{"ore1", "ore2", "ore3"}
	oreChan := make(chan string)
	// Finder
	go func(mine [3]string) {
		for _, item := range mine {
			time.Sleep(1 * time.Second)
			oreChan <- item //send
		}
	}(theMine)

	// Ore Breaker
	for i := 0; i < 3; i++ {
		go func(ch chan string) {
			foundOre := <-ch //receive
			fmt.Println("Miner: Received " + foundOre + " from finder")
		}(oreChan)
	}

	<-time.After(time.Second * 5)
}
