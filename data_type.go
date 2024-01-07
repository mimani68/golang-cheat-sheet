package main

import "fmt"

func main() {
	// Integer types
	var i int = 10
	var ui uint = 20
	var i8 int8 = -30
	var ui8 uint8 = 40
	var i16 int16 = -50
	var ui16 uint16 = 60
	var i32 int32 = -70
	var ui32 uint32 = 80
	var i64 int64 = -90
	var ui64 uint64 = 100

	fmt.Println("Int:", i)
	fmt.Println("Unsigned Int:", ui)
	fmt.Println("Int8:", i8)
	fmt.Println("Unsigned Int8:", ui8)
	fmt.Println("Int16:", i16)
	fmt.Println("Unsigned Int16:", ui16)
	fmt.Println("Int32:", i32)
	fmt.Println("Unsigned Int32:", ui32)
	fmt.Println("Int64:", i64)
	fmt.Println("Unsigned Int64:", ui64)

	// Floating-point types
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793238
	fmt.Println("Float32:", f32)
	fmt.Println("Float64:", f64)

	// Boolean type
	var b bool = true
	fmt.Println("Boolean:", b)

	// String type
	var s string = "Hello, Go!"
	fmt.Println("String:", s)

	// Rune type
	var r rune = 'A'
	fmt.Println("Rune:", r)

	// Array type
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Slice type
	var slice []int = []int{1, 2, 3}
	fmt.Println("Slice:", slice)

	// Map type
	var m map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Map:", m)

	// Struct type
	type Person struct {
		Name string
		Age  int
	}
	var p Person = Person{Name: "Alice", Age: 30}
	fmt.Println("Struct:", p)
}
