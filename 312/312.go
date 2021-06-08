package main

import "fmt"

func maxCoins(nums []int) int {
	n := len(nums)
	val := make([]int, n+2)
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	val[0], val[n+1] = 1, 1
	//[i][j] i到j的的最多硬币数
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i] * val[k] * val[j]
				sum += dp[i][k] + dp[k][j]
				dp[i][j] = max(dp[i][j], sum)
			}
		}
	}
	return dp[0][n+1]
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
func main() {
	nums := []int{3, 1, 5, 8}
	fmt.Println(maxCoins(nums))
}
