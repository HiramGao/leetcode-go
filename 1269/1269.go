package main

import (
	"fmt"
)

/**
i步后，在下标j 的方案数
dp[i][j] when 0<=i<=steps 0<=j<=min(steps,arrLen-1)
dp[0][0] = 1
dp[0][j] = 0  when 1<=j<=min(steps,arrLen-1)
			i-1 Right      i-1 Stay     i-1 Left
dp[i][j] = dp[i-1][j-1] + dp[i-1][j] + dp[i-1][j+1]
*/
func numWays(steps int, arrLen int) int {
	const mod = 1e9 + 7
	minJ := steps
	if arrLen-1 < minJ {
		minJ = arrLen - 1
	}
	dp := make([][]int, steps+1)
	for i := range dp {
		dp[i] = make([]int, minJ+1)
	}
	dp[0][0] = 1

	for i := 1; i <= steps; i++ {
		for j := 0; j <= minJ; j++ {
			dp[i][j] = dp[i-1][j]
			if j-1 >= 0 {
				dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % mod
			}
			if j+1 <= minJ {
				dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % mod
			}
		}
	}
	return dp[steps][0]
}

func numWays2(steps int, arrLen int) int {
	const mod = 1e9 + 7
	minJ := steps
	if arrLen-1 < minJ {
		minJ = arrLen - 1
	}
	dp := make([]int, minJ+1)
	dp[0] = 1

	for i := 1; i <= steps; i++ {
		newDp := make([]int, minJ+1)
		for j := 0; j <= minJ; j++ {
			newDp[j] = dp[j]
			if j-1 >= 0 {
				newDp[j] = (newDp[j] + dp[j-1]) % mod
			}
			if j+1 <= minJ {
				newDp[j] = (newDp[j] + dp[j+1]) % mod
			}
		}
		dp = newDp
	}
	return dp[0]
}
func main() {
	fmt.Println(numWays2(4, 2))
}
