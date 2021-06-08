package main

import "fmt"

func lengthOfLIS(nums []int) (ans int) {
	var (
		n  = len(nums)
		dp = make([]int, n)
	)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	ans = dp[0]
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		ans = max(ans, dp[i])
	}
	return
}

func max(i, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
func main() {
	fmt.Print(lengthOfLIS([]int{0}))
}
