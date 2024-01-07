package main

import "fmt"

// Global variable
var globalVariable int = 100

func main() {
	// Local variable
	var localVariable int = 10

	fmt.Println("Local variable:", localVariable)
	fmt.Println("Global variable:", globalVariable)

	// Updating the local variable
	localVariable = 20

	fmt.Println("Updated local variable:", localVariable)

	// Updating the global variable
	globalVariable = 200

	fmt.Println("Updated global variable:", globalVariable)
}
