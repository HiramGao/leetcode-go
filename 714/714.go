package main

import "fmt"

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	sell, buy := 0, -prices[0]

	for i := 1; i < n; i++ {
		sell = max(sell, buy+prices[i]-fee)
		buy = max(buy, sell-prices[i])
	}
	return sell
}
func max(i, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println(maxProfit(prices, fee))
}
