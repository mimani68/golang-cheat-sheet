package main

import "fmt"

func main() {

	var i interface{} = 23
	fmt.Printf("%v\n", i)

	var value interface{} = map[string]int{
		"id":      1,
		"message": 0,
	}
	fmt.Printf("%v\n", value)

	var simpleObject interface{} = struct {
		Id     int
		Messag string
	}{
		Id:     int(18),
		Messag: "hello guys",
	}
	fmt.Printf("%v\n", simpleObject)
}
