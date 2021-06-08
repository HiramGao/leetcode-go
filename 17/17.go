package main

import "fmt"

func letterCombinations(digits string) (ans []string) {

	if digits == "" {
		return []string{}
	}

	digByte := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	n := len(digits)
	set := ""
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == n {
			ans = append(ans, set)
			return
		}
		for _, v := range digByte[digits[i]] {
			set = set + string(v)
			backtrack(i + 1)
			set = set[:len(set)-1]
		}
	}

	backtrack(0)
	return ans
}
func main() {
	fmt.Print(letterCombinations("23"))
}
