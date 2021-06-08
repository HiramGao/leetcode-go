package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	s = strings.ToLower(s)
	isalnum := func(ch byte) bool {
		return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')

	}
	for l, r := 0, n-1; l < r; {
		for l < r && !isalnum(s[l]) {
			l++
		}
		for l < r && !isalnum(s[r]) {
			r--
		}
		if l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
	}
	return true
}
func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
