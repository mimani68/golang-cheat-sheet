package main

import "fmt"

func main() {
	value := new(string)
	fmt.Printf("%T\n", value)
	fmt.Printf("%v\n", value)
	fmt.Printf("%s\n", *value)
	*value = "Hello dear my friends"
	fmt.Printf("%T\n", value)
	fmt.Printf("%v\n", value)
	fmt.Printf("%s\n", *value)
}
