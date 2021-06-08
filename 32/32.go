package main

import "fmt"

func longestValidParentheses(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	res := 0
	dp := make([]int, n)

	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			res = max(res, dp[i])
		}
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses(""))
	fmt.Println(longestValidParentheses("()(())"))
}
