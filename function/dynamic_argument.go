package main

import (
	"fmt"
	"strings"
)

func dynamicParameter1(number int, params ...string) {
	fmt.Println(params)
}

func dynamicParameter2(params ...interface{}) {
	fmt.Println(params)
}

func main() {

	dynamicParameter1(2, "Hey", "khobi")
	dynamicParameter2("Hey", func(a string) (output string) { output = strings.ReplaceAll(a, "-", ""); return })

}
