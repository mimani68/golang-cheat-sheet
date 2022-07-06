package main

import (
	"fmt"
	. "fmt"
	"time"
)

func main() {
	list := []string{"Ali", "Reza", "Sina"}
	for index, value := range list {
		Printf("%d %s \n", index, value)
	}

	// Unlimit loop
	for {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Unlimit loop")
	}
}
