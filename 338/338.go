package main

import (
	"fmt"
)

func countBits(num int) []int {
	dp := make([]int, num+1)

	for i := 0; i <= num; i++ {
		dp[i] = dp[i>>1] + i&1
	}

	return dp
}

func main() {
	fmt.Print(countBits(15))
}
