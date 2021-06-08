package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {

	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt64
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}
func main() {
	fmt.Print(numSquares(12))
}
