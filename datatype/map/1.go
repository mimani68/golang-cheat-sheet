package main

import "fmt"

func main() {

	var a map[string]string
	a = map[string]string{}
	a["username"] = "ali"
	a["username"] = "hoda"
	fmt.Printf("%p\n", a)

	b := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(b)

	// Generate mao using Make
	c := make(map[string]string, 0)
	c["content"] = "Lorem ipsum"
	fmt.Println(c)

	// To remove all key/value pairs from a map, use the clear builtin.
	d := make(map[string]string, 0)
	d["k1"] = "Ops"
	d["k2"] = "Yup"
	fmt.Println(d)

	// Remove one item from Map
	delete(d, "k2")
	fmt.Println(d)

	// Remove all items from Map
	clear(d)
	fmt.Println(d)

	// Map inside Map
	e := make(map[string]map[string]string, 0)
	e["user"] = map[string]string{"name": "reza"}
	// d["user"]["name"] = "reza" // error => panic: assignment to entry in nil map
	fmt.Println(e)

}
