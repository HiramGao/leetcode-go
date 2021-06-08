package main

import "fmt"

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	//dp[i][0] 手上持有股票的最大收益
	//dp[i][1] 手上未持有股票 并且处于冷冻期
	//dp[i][2] 手上未持有股票 并且不处于冷冻期
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[n-1][1], dp[n-1][2])
}
func max(i, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
func maxProfit1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	//f[0] 手上持有股票的最大收益
	//f[1] 手上未持有股票 并且处于冷冻期
	//f[2] 手上未持有股票 并且不处于冷冻期

	f := [3]int{-prices[0], 0, 0}
	for i := 1; i < n; i++ {
		newF := [3]int{
			max(f[0], f[2]-prices[i]),
			f[0] + prices[i],
			max(f[1], f[2]),
		}
		f = newF
	}
	return max(f[1], f[2])
}
func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfit1(prices))
}
