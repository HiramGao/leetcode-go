package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	res := 0
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				res = 1
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}
	return res * res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '1', '1', '1'}}
	fmt.Println(maximalSquare(matrix))
}
