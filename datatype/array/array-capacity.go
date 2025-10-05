package main

import "fmt"

func main() {
	// Create an array
	arr := []string{"This", "is", "the", "tutorial", "of", "Go", "language", "!"}

	// Create a slice from the array
	myslice := arr[2:6] // Low=2 , High=6
	fmt.Printf("Array: %v\n", arr)
	// Array: [This is the tutorial of Go language !]
	fmt.Printf("Length: %d\n", len(arr))
	// Length: 8
	fmt.Printf("Slice: %v\n", myslice)
	// Slice: [the tutorial of Go]
	fmt.Printf("Length of the slice: %d\n", len(myslice))
	// len(myslice) = 6(High) - 2 (Low) = 4
	// Length of the slice: 4
	fmt.Printf("Capacity of the slice: %d\n", cap(myslice))
	// cap(myslice) = 8 (Underlying array length) - 2 (Low) = 6 → from index 2 to the end of the underlying array there are 8 elements (["the", "tutorial", "of", "Go", "language", "!"])
	// Capacity of the slice: 6
}
