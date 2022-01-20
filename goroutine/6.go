package main

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
			oreChan <- item //send
		}
	}(theMine)

	// Ore Breaker
	go func(ch chan string) {
		for i := 0; i < 3; i++ {
			foundOre := <-ch //receive
			fmt.Println("Miner: Received " + foundOre + " from finder")
		}
	}(oreChan)

	fmt.Println(<-oreChan)
	<-time.After(time.Second * 5)
}
