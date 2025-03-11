package main

import (
	"fmt"
	"sync"
)

type StateHolder struct {
	State sync.Map `json:"state"`
}

var stateHolder = StateHolder{}

func main() {
	stateHolder.State.Store("setting-debug", false)
	value, _ := stateHolder.State.Load("setting-debug")
	fmt.Println(value)
}
