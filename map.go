package main

import "fmt"

func main() {

	/**
	 *
	 * 01
	 *
	 */
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	/**
	 *
	 * 02
	 *
	 */
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	/**
	 *
	 * 03
	 *
	 */
	type sampleMapSchema map[string]int
	o := sampleMapSchema{
		"A": 1,
	}

	fmt.Println("o:", o)
}
