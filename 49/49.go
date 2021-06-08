package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	hashMap := map[string][]string{}

	for _, str := range strs {
		s1 := []byte(str)
		sort.Slice(s1, func(i, j int) bool {
			return s1[i] < s1[j]
		})
		s := string(s1)
		hashMap[s] = append(hashMap[s], str)
	}
	res := [][]string{}
	for _, strings := range hashMap {
		res = append(res, strings)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
