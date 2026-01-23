package main

import (
	"fmt"

	"test.io/api/v1/goroutine/mathmodule"
)

func main() {
	// Pass goroutine and channel between modules
	chIns := make(chan mathmodule.Result)
	go mathmodule.Calc(5, chIns)
	res := <-chIns
	fmt.Printf("state:\"%t\" and result:%d\n", res.Error, res.Result)
}
