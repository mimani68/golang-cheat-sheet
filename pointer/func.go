package main

import (
	"fmt"
)

func main() {
	value := *getPointer()
	memoryAddress := getPointer()
	fmt.Printf("Value is %d and address is %v\n", value, memoryAddress)
}

func getPointer() (myPointer *int) {
	a := 234
	return &a
}
