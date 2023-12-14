package main

import "fmt"

func main() {
	var brands = []string{
		"BMW", "HYNDAY", "KOWASAKI", "JET", "BOING",
	}
	fmt.Printf("%T\n", append(brands[:1], brands[2:]...))
	fmt.Printf("%v\n", append(brands[:1], brands[2:]...))
	fmt.Printf("%s\n", append(brands[:1], brands[2:]...))
}
