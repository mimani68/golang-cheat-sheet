package main

import (
	"fmt"
	"time"
)

func DateFn() {
	t1 := time.Now()
	fmt.Println(t1.Format(time.RFC3339))

	t2 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
	fmt.Println(t2.Format(time.RFC3339))

	t3, _ := time.Parse(time.RFC3339, "2018-10-17T07:26:33Z")
	fmt.Println(t3.Year())
	fmt.Println(t3.Format(time.RFC3339))
}
