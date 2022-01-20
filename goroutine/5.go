package goroutine

import (
	"fmt"
	"time"
)

func main() {
	heroes := []string{"Marvel", "Flash", "Thanos", "Eagle", "Hulk", "Thor"}
	spaceships := []string{"Battlecruiser", "Battleship", "Cruiseship", "Her Majesty's Ship", "Imperial Spaceship"}
	go externalAPI(heroes, "hero")
	go externalAPI(spaceships, "spaceship")

	<-time.After(time.Second * 10)
}

func externalAPI(items []string, label string) {
	for i := range items {
		fmt.Printf("%s: %s\n", label, items[i])
		time.Sleep(time.Second)
	}
}
