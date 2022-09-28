package main

import "fmt"

func main() {
	value := new(int)
	fmt.Printf("%T\n", value)
	fmt.Printf("%v\n", value)
	fmt.Printf("%d\n", *value)
	*value = 18
	fmt.Printf("%T\n", value)
	fmt.Printf("%v\n", value)
	fmt.Printf("%d\n", *value)
}
