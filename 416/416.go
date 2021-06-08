package main

import "fmt"

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum, maxVal := 0, 0
	for _, num := range nums {
		sum += num
		if num > maxVal {
			maxVal = num
		}
	}
	if sum%2 != 0 || maxVal > sum/2 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[target]
}
func main() {
	fmt.Print(canPartition([]int{2, 2, 3, 5}))
}
