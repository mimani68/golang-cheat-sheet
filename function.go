package main

import (
	"fmt"
)

func app(i interface{}) interface{} {
	return i
}

func dynamicParameter(number int, parametersList ...string) {
	fmt.Println(parametersList)
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

	dynamicParameter(2, "salam", "khobi")

}
