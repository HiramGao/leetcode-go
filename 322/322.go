package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {

	var dp = make([]int, amount+1)

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func main() {
	fmt.Print(coinChange([]int{1, 2, 5}, 11))
}
