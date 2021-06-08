package main

import (
	"fmt"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	emailToIndex := map[string]int{}
	emailToName := map[string]string{}

	for _, account := range accounts {
		name := account[0]
		for _, email := range account[1:] {
			if _, ok := emailToIndex[email]; !ok {
				emailToIndex[email] = len(emailToIndex)
				emailToName[email] = name
			}
		}
	}
	parent := make([]int, len(emailToIndex))

	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if x != parent[x] {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x, y int) {
		x, y = find(x), find(y)
		if x == y {
			return
		}
		parent[x] = parent[y]
	}

	for _, account := range accounts {
		firstIndex := emailToIndex[account[1]]
		for _, email := range account[2:] {
			union(emailToIndex[email], firstIndex)
		}
	}
	indexToEmails := map[int][]string{}

	for email, index := range emailToIndex {
		index = find(index)
		indexToEmails[index] = append(indexToEmails[index], email)
	}
	ans := [][]string{}
	for _, emails := range indexToEmails {
		sort.Strings(emails)
		account := append([]string{emailToName[emails[0]]}, emails...)
		ans = append(ans, account)
	}
	return ans
}
func main() {
	acounts := [][]string{{"John", "johnsmith@mail.com", "john00@mail.com"}, {"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"Mary", "mary@mail.com"}}
	fmt.Println(accountsMerge(acounts))
}
