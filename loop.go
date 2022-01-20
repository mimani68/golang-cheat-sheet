package main

import (
	. "fmt"
)

func main() {
	list := []string{"Ali", "Reza", "Sina"}
	for index, value := range list {
		Printf("%d %s \n", index, value)
	}
}
