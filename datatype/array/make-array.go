package main

import (
	"fmt"
)

func main() {
	var a = make([]int, 1, 10)
	a = append(a, 11, 12)
	fmt.Println(a)
}

// OUTPUT
// [0 11 12]
