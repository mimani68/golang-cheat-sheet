package main

import (
	"fmt"
	"time"
)

func main() {
	list := []string{"Ali", "Reza", "Sina"}
	for index, value := range list {
		fmt.Printf("%d %s \n", index, value)
	}

	object := map[string]string{"username": "ali224", "age": "15"}
	for index, value := range object {
		fmt.Printf("%s %s \n", index, value)
	}

	for index := range object {
		fmt.Printf("%s %s\n", index, object[index])
	}

	// Unlimited loop
	for {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Unlimited loop")
	}
}
