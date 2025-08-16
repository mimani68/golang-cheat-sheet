package main

import "fmt"

func main() {
	privateFunc(map[string]string{
		"id":   "id2384569",
		"name": "Reza",
	}, map[string]string{
		"id":   "id9394",
		"name": "Ali",
	})
}

func privateFunc(a ...interface{}) {
	fmt.Println(a)
	fmt.Println(a...)
	for index, param := range a {
		fmt.Printf("[%d] Type=%T %v \n", index, param, param)
	}
}
