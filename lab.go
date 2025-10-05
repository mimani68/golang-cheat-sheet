package main

import (
	"fmt"
)

func main() {
	fmt.Printf("fmt.Sprintf(42, a): \"42 \" (%T)", fmt.Sprintf("%s", 42))
}
