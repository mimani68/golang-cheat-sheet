package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffers := make([]*bytes.Buffer, 5)
	for i := 0; i < 5; i++ {
		buffers[i] = bytes.NewBuffer(make([]byte, 0, 1024))
		buffers[i].Reset()
		buffers[i].WriteString(fmt.Sprintf("Buffer %d", i))
		fmt.Printf("   Got buffer %d: %s\n", i, buffers[i].String())
	}
}
