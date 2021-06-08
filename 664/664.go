package main

import "fmt"

func strangePrinter(s string) int {

	//dp[i][j] s[i:j]打印的最少次数
	//1. dp[i][j] = 1+dp[i+1][j] i 单独打印 s[i+1][j] 另外打印
	//2. dp[i][j] = min(dp[i][j],dp[i+1][k]+dp[k+1][j]) i 单独打印 s[i+1][j] 另外打印 when s[i]==s[k]
	//3. dp[i][j] = min(dp[i][j],dp[i+1][j]) dp[i + 1][j]代表将i放入[j + 1, i]一起打印 when s[i] == s[j]
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = dp[i+1][j] + 1
			for k := i + 1; k < j; k++ {
				if s[i] == s[k] {
					dp[i][j] = min(dp[i][j], dp[i+1][k]+dp[k+1][j])
				}
			}
			if s[i] == s[j] {
				dp[i][j] = min(dp[i][j], dp[i+1][j])
			}
		}
	}
	return dp[0][n-1]
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}
func main() {
	s := "aaaaaaaa"
	fmt.Println(strangePrinter(s))
}
