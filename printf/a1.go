package main

import "fmt"

func main() {
	fmt.Printf("%s\n", "hello")    // prints hello
	fmt.Printf("%q\n", "hello")    // prints "hello"
	fmt.Printf("%s\n", "hello\n;") // prints hello
	//; \n is not escaped
	fmt.Printf("%q\n", "hello\n;") // prints "hello\n;" \n is escaped here

	a1 := fmt.Sprintf("Name: %s, Age: %d, Score: %.2f", "John", 30, 95.5)
	fmt.Println(a1) // Name: John, Age: 30, Score: 95.50

	fmt.Printf("%T", map[int]interface{}{1: "salam"}) // map[int]interface {}
}
