package main

import "fmt"

func PointerFn() {
	a := 1
	fmt.Println(&a)    // 0xc42000e1f8
	fmt.Println(*(&a)) // 1

	b := struct{ Id string }{Id: "hf94ytgjn"}
	fmt.Println(&b)    // ${hf94ytgjn}
	fmt.Println(*(&b)) // {hf94ytgjn}
}
