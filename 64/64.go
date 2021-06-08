package main

import "fmt"

func minPathSum(grid [][]int) int {
	var (
		rows = len(grid)
		cols = len(grid[0])
		dp   = make([][]int, rows)
	)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < cols; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	fmt.Println(dp)
	return dp[rows-1][cols-1]

}

func min(i, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}
func main() {
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}
