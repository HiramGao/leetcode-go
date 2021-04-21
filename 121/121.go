package main

func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	min := prices[0]
	maxDiff := 0

	for i := 1; i < n; i++ {
		if min > prices[i] {
			min = prices[i]
		} else {
			tmp := prices[i] - min
			if tmp > maxDiff {
				maxDiff = tmp
			}
		}
	}
	return maxDiff
}
func main() {
	println(maxProfit([]int{7, 6, 4, 3, 1}))
}
