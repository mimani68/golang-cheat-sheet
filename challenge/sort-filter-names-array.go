package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	userNames := make([]string, 30)
	for i := range userNames {
		userNames[i] = fmt.Sprintf("user%d", rand.Intn(100))
	}

	freq := make(map[string]int)
	for _, name := range userNames {
		freq[name]++
	}

	type user struct {
		name  string
		count int
	}

	userList := make([]user, len(freq))
	i := 0
	for name, count := range freq {
		userList[i] = user{name, count}
		i++
	}

	sort.Slice(userList, func(i, j int) bool {
		return userList[i].count > userList[j].count
	})

	topUsers := userList[:4]

	fmt.Println("Top 4 users:")
	for _, u := range topUsers {
		fmt.Printf("%s - %d occurrences\n", u.name, u.count)
	}
}
