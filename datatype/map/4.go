package main

import "fmt"

func main() {

	var a map[interface{}]interface{}
	a["id"] = 1.4
	fmt.Println(a)

}

// OUTPUT
// panic: assignment to entry in nil map
