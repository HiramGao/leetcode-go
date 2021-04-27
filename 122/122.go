package main

func maxProfit(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][1] = -prices[0]
	n := len(prices)

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(i, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
func main() {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
