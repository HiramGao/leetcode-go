package main

import (
	"fmt"
	"math"
)

func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}
	//dp[i][j][k]
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			for m := 2; m <= k; m++ {
				dp[i][j][m] = math.MaxInt64
			}
		}
	}

	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + stones[i-1]
	}

	for l := 2; l <= n; l++ {
		for i := 1; i+l-1 <= n; i++ {
			j := i + l - 1
			for m := 2; m <= k; m++ {
				for p := i; p < j; p += k - 1 {
					dp[i][j][m] = min(dp[i][j][m], dp[i][p][1]+dp[p+1][j][m-1])
				}
			}
			dp[i][j][1] = dp[i][j][k] + sum[j] - sum[i-1]
		}
	}

	return dp[1][n][1]
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}
func main() {
	stones := []int{3, 5, 1, 2, 6}
	fmt.Println(mergeStones(stones, 3))

}
