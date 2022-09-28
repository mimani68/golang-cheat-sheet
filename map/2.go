package main

import (
	"fmt"
)

func main() {
	a := make(map[string]int)
	a["k1"] = 7
	a["k2"] = 13

	fmt.Println("map:", a)
	fmt.Println("len:", len(a))

	value, isExists := a["k1"]
	fmt.Println("value:", value)
	fmt.Println("isExists:", isExists)

	delete(a, "k2")
	fmt.Println("map:", a)
}
