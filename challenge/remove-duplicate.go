package main

import "fmt"

// please remove duplicate and save order
func main() {
	items := []int{7, 7, 1, 2, 3, 1, 5, 2, 2, 3, 4, 2}
	newItemsMap := make(map[int]bool)
	newItems := []int{}
	for _, item := range items {
		if !newItemsMap[item] {
			newItems = append(newItems, item)
		}
		newItemsMap[item] = true
	}
	fmt.Println(newItems)
}
