package main

import "math"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	min := prices[0]
	maxDiff := prices[1] - min

	for i := 2; i < len(prices); i++ {
		if prices[i-1] < min {
			min = prices[i-1]
		}
		currentDiff := prices[i] - min
		if currentDiff > maxDiff {
			maxDiff = currentDiff
		}
	}
	if maxDiff < 0 {
		return 0
	}
	return maxDiff
}
func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	min := math.MaxInt32
	maxDiff := 0

	for _, v := range prices {
		if v < min {
			min = v
		} else {
			currDiff := v - min
			if currDiff > maxDiff {
				maxDiff = currDiff
			}
		}
	}
	return maxDiff
}
func main() {
	println(maxProfit2([]int{7, 6, 4, 3, 1}))
}
