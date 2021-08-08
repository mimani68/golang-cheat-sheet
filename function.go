package main

import (
	"fmt"
)

func app(i interface{}) interface{} {
	return i
}

func main() {

	if true {
		a := app("salam")
		fmt.Println(a)
	}

	// Call function imidiatly
	func(a string) {
		fmt.Println(a)
	}("salam")

}
