package main

import "fmt"

func app(i interface{}) interface{} {
	return i
}

func main() {
	fmt.Println(app("Hey"))
	fmt.Println(app(1))
	fmt.Println(app(map[string]string{
		"username": "Mahdi",
	}))
}
