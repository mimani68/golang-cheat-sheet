// Go version 1.23+

package main

import (
	"fmt"
	"slices"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Gopher", 13},
		{"Alice", 20},
		{"Bob", 5},
		{"Vera", 24},
		{"Zac", 15},
	}

	// Chunk people into slices of 2 elements each
	for _, c := range slices.Chunk(people, 2) {
		fmt.Println(c)
	}
}
