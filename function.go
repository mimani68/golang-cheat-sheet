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

	//
	// Call function immediately
	//
	func(a string) {
		fmt.Println(a)
	}("salam")

	//
	// Dynamic argument in method
	//
	dynamicParameter(2, "salam", "khobi")

	//
	// Anonymouse function with initial mode
	//
	(func(id int) {
		fmt.Println(id)
	})(12)

	//
	// Callback function
	//
	sayHello := func(cb func(message string)) {
		cb("Hello")
	}
	sayHello(func(message string) {
		fmt.Println(message)
	})
}
