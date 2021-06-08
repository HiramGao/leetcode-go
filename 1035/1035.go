package main

import "fmt"

func maxUncrossedLines(nums1 []int, nums2 []int) int {

	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	dp[0] = make([]int, n+1)

	for i, n1 := range nums1 {
		dp[i+1] = make([]int, n+1)
		for j, n2 := range nums2 {
			if n1 == n2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[m][n]
}
func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
func main() {
	nums1 := []int{2, 5, 1, 2, 5}
	nums2 := []int{10, 5, 2, 1, 5, 2}
	fmt.Println(maxUncrossedLines(nums1, nums2))
}
