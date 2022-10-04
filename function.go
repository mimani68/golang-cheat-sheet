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

	(func(id int, users ...string) {
		for index := 0; index < len(users); index++ {
			fmt.Printf("Hello %s \n", users[index])
		}
	})(12, "Matt", "Daneil", "Alison")

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
