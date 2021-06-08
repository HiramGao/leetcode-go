package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	res := 0

	hashMap := [128]int{}
	for i := range hashMap {
		hashMap[i] = -1
	}
	left := 0
	for right, value := range []byte(s) {
		left = max(left, hashMap[value]+1)
		res = max(res, right-left+1)
		hashMap[value] = right
	}
	return res
}

func max(j, i int) int {
	if j > i {
		return j
	}
	return i
}

func main() {
	fmt.Println(lengthOfLongestSubstring("bbbbbb"))
}
