package main

import "fmt"

func application_void() {
	fmt.Println("Void return applications")
}

func application() (message string) {
	// message := "Hi"  // variable already instantiated
	message = "Hi"
	return
}

func main() {
	response := application()
	fmt.Println(response)

	// empty_response := application_void()  // application_void() (no value) used as value
	// fmt.Println(empty_response)
}
