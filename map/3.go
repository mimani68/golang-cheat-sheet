package main

import "fmt"

func main() {

	type sampleMapSchema map[string]int
	a := sampleMapSchema{
		"A": 1,
	}

	fmt.Println("o:", a)
}
