package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// Create a slice
	s := make([]int, 10, 20)
	fmt.Printf("Initial Length: %d, Capacity: %d\n", len(s), cap(s))

	// Manipulate the slice header using unsafe
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	hdr.Len = 5
	hdr.Cap = 15
	fmt.Printf("After Manipulation: Length: %d, Capacity: %d\n", len(s), cap(s))
}

// Note: This example uses reflect.SliceHeader which is not directly exposed in the standard library.
// You would typically use the reflect package for such manipulations.
