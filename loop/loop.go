package main

import (
	"fmt"
)

func main() {
	list := []string{"Ali", "Reza", "Sina"}
	fmt.Println("=== List ===")
	for index, value := range list {
		fmt.Printf("%d %s \n", index, value)
	}

	object := map[string]string{"username": "ali224", "age": "15"}
	fmt.Println("=== Object ===")
	for index, value := range object {
		fmt.Printf("%s %s \n", index, value)
	}

	fmt.Println("=== Object ===")
	for index := range object {
		fmt.Printf("%s %s\n", index, object[index])
	}

	// Unlimited loop
	// fmt.Println("=== Infinite loop ===")
	// for {
	// 	time.Sleep(500 * time.Millisecond)
	// 	fmt.Println("Unlimited loop")
	// }
}
