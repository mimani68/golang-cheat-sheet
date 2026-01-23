package main

import "fmt"

func main() {

	//
	// Anonymous function with initial mode
	//
	(func(id int) {
		fmt.Println(id)
	})(12)
}
