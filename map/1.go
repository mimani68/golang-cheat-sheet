package main

import "fmt"

func main() {

	var a map[string]string
	a = map[string]string{}
	a["username"] = "ali"
	a["username"] = "hoda"
	fmt.Printf("%p\n", a)

	b := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", b)

}
