package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	//dp[i][j] i到j是否为回文串
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	begin := 0
	maxLen := 1

	for r := 0; r < n; r++ {
		for l := 0; l < r; l++ {
			if s[l] != s[r] {
				dp[l][r] = false
			} else {
				if r-l <= 2 {
					dp[l][r] = true
				} else {
					dp[l][r] = dp[l+1][r-1]
				}
			}
			if dp[l][r] {
				curlen := r - l + 1
				if curlen > maxLen {
					maxLen = curlen
					begin = l
				}
			}
		}
	}

	return s[begin : begin+maxLen]
}

/**
中心扩展法
*/
func longestPalindrome1(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	l, r := 0, 0
	for i := range s {
		l1, r1 := expandAroundCenter(s, i, i)
		l2, r2 := expandAroundCenter(s, i, i+1)
		if r1-l1 > r-l {
			l, r = l1, r1
		}
		if r2-l2 > r-l {
			l, r = l2, r2
		}

	}
	return s[l+1 : r]
}

func expandAroundCenter(s string, l int, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l, r
}
func main() {
	fmt.Println(longestPalindrome("aaaaa"))
}
