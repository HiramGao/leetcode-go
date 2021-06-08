package main

import (
	"fmt"
	"strings"
)

func longestSubstring(s string, k int) int {
	if s == "" {
		return 0
	}
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	var split byte
	res := 0
	for i, v := range cnt {
		if 0 < v && v < k {
			split = 'a' + byte(i)
			break
		}
	}
	if split == 0 {
		return len(s)
	}

	for _, subStr := range strings.Split(s, string(split)) {
		res = max(res, longestSubstring(subStr, k))
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(longestSubstring("aaabb", 3))
	fmt.Println(longestSubstring("ababbc", 2))
}
