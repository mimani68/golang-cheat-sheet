package main

import "fmt"

var key = "name"

func main() {
	class := []interface{}{"first"}
	class = append(class, string("second"))
	fmt.Println(class)

	userOne := map[string]interface{}{}
	userOne[key] = "ali"
	class = append(class, userOne)
	fmt.Println(class)

	userTwo := struct {
		Name   string
		Family string
	}{
		Name:   "Mahdi",
		Family: "Imani",
	}
	class = append(class, userTwo)
	fmt.Println(class)

	printInterfaceName(class)
}

// func printInterfaceName(list ...interface{}) {
func printInterfaceName(list []interface{}) {
	fmt.Println(list[1])
	fmt.Println(list...)
	for index, _ := range list {
		fmt.Println(list[index])
	}
}
