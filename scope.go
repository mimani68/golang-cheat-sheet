package main

import "fmt"

func main() {
	a := 1
	{
		a := 2
		fmt.Printf("* %d\n", a)
	}
	fmt.Printf("** %d\n", a)
}
