package main

import (
	"fmt"
)

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]
	check := func(m int) bool {
		x, y := n-1, 0
		num := 0
		for x >= 0 && y < n {
			if matrix[x][y] <= m {
				num += x + 1
				y++
			} else {
				x--
			}
		}
		return num >= k
	}
	for l < r {
		m := l + (r-l)/2
		if check(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(kthSmallest([][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15}}, 8))
}
