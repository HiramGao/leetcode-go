package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)

	dp := make(map[int]int)
	dp[nums[0]]++
	dp[-nums[0]]++

	for i := 1; i < n; i++ {
		nextDp := make(map[int]int)

		for sum := range dp {
			nextDp[sum+nums[i]] += dp[sum]
			nextDp[sum-nums[i]] += dp[sum]
		}
		dp = nextDp
	}
	fmt.Println(dp)
	if target > 1000 {
		return 0
	}
	return dp[target]
}
func main() {
	fmt.Println(findTargetSumWays([]int{0, 1}, 1))
}
