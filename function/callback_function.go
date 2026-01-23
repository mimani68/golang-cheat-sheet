package main

import "fmt"

func main() {

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
