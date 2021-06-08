package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	minStrLen := math.MaxInt32
	left, right := 0, 0
	res := ""
	hashMap := map[byte]int{}
	for _, v := range []byte(t) {
		hashMap[v]++
	}
	check := func() bool {
		for _, v := range hashMap {
			if v > 0 {
				return false
			}
		}
		return true
	}

	for right < len(s) {
		hashMap[s[right]]--
		for check() {
			if right-left+1 < minStrLen {
				minStrLen = right - left + 1
				res = s[left : right+1]
			}
			hashMap[s[left]]++
			left++
		}
		right++
	}

	return res
}
func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
