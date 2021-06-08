package main

import "fmt"

func countSubstrings(s string) int {
	n := len(s)
	// i...j是否为回文
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	res := n
	for r := range s {
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
				res++
			}
		}
	}

	return res
}

func main() {
	fmt.Println(countSubstrings("aaa"))
}
