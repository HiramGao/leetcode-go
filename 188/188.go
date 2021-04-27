package main

import "math"

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	k = min(k, n/2)
	buy := make([][]int, n)
	sell := make([][]int, n)
	for i, _ := range buy {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}
	buy[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt64
		buy[0][i] = math.MinInt64
	}
	for i := 1; i < n; i++ {
		buy[i][0] = max(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}
	maxValue := sell[n-1][0]

	for _, v := range sell[n-1] {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

func min(i, i2 int) int {
	if i > i2 {
		return i2
	}
	return i
}
func max(i, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
func main() {
	println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}
