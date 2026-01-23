package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	oreChan := make(chan string)

	theMine := [3]string{"ore1", "ore2", "ore3"}

	// Miner
	go func(mine [len(theMine)]string) {
		for _, item := range mine {
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Println("Miner: Find " + item)
			oreChan <- item
		}
	}(theMine)

	// Extractor
	for i := 0; i < len(theMine); i++ {
		go func(ch chan string) {
			foundOre := <-ch //receive
			fmt.Println("Extractor: Received " + foundOre + " from finder")
		}(oreChan)
	}

	<-time.After(time.Second * 6)
}

//
// OUTPUT
//
// Miner: Find ore1
// Miner: Find ore2
// Extractor: Received ore1 from finder
// Extractor: Received ore2 from finder
// Miner: Find ore3
// Extractor: Received ore3 from finder
