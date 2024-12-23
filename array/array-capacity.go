package main

import "fmt"

func main() {
	// Create an array
	arr := []string{"This", "is", "the", "tutorial", "of", "Go", "language"}

	// Create a slice from the array
	myslice := arr[1:6]
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Slice: %v\n", myslice)
	fmt.Printf("Length of the slice: %d\n", len(myslice))
	fmt.Printf("Capacity of the slice: %d\n", cap(myslice))
}
