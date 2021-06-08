package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	if nums[0] > nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}

	for i := 2; i < n; i++ {
		if nums[i]+dp[i-2] > dp[i-1] {
			dp[i] = nums[i] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n-1]
}
func main() {
	fmt.Print(rob([]int{2, 1}))
}
