package main

import "fmt"

func main() {
	user := new(string)
	*user = "Mahdi"
	fmt.Printf("%t\n", user)
}
