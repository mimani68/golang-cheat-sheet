package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(t1.Format(time.RFC3339))

	t2 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
	fmt.Println(t2.Format(time.RFC3339))

	t3, _ := time.Parse(time.RFC3339, "2018-10-17T07:26:33Z")
	fmt.Println(t3.Year())
	fmt.Println(t3.Format(time.RFC3339))

	t4 := time.Now().Format(time.RFC3339)
	fmt.Printf("%T\n", t4)
	fmt.Println(t4 >= "2025-02-00T00:00:00Z")
	fmt.Println("2021-02-02T00:00:00Z" >= "2025-01-27T00:00:00Z")

	hour := 5
	minutes := 10
	seconds := 0
	t5 := time.Now().Add(time.Hour*time.Duration(hour) + time.Minute*time.Duration(minutes) + time.Second*time.Duration(seconds))
	fmt.Println(t5)
}
