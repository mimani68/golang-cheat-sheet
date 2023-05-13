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
		a := app("Hey")
		fmt.Println(a)
	}

	//
	// Call function immediately
	//
	func(a string) {
		fmt.Println(a)
	}("Hey")

	//
	// Dynamic argument in method
	//
	dynamicParameter(2, "Hey", "khobi")

	//
	// Anonymous function with initial mode
	//
	(func(id int) {
		fmt.Println(id)
	})(12)

	//
	// Callback function
	//
	sayHello := func(cb func(message string) string) {
		result := cb("Hello")
		fmt.Println(result)
	}
	sayHello(func(message string) string {
		return "[INFO]" + message
	})
}
