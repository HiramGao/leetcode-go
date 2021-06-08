package main

import "fmt"

type pair struct {
	x, y int
}

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	if m == 0 && n == 0 {
		return 0
	}
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}
	res := 0
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if visited[x][y] != 0 {
			return visited[x][y]
		}
		visited[x][y]++
		for _, p := range dir4 {
			newX, newY := x+p.x, y+p.y
			if 0 <= newX && newX < m && 0 <= newY && newY < n && matrix[newX][newY] > matrix[x][y] {
				visited[x][y] = max(visited[x][y], dfs(newX, newY)+1)
			}

		}
		return visited[x][y]
	}

	for x, rows := range matrix {
		for y := range rows {
			res = max(res, dfs(x, y))
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	matrix := [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}
	fmt.Println(longestIncreasingPath(matrix))
	matrix = [][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}
	fmt.Println(longestIncreasingPath(matrix))
	matrix = [][]int{{1}}
	fmt.Println(longestIncreasingPath(matrix))
}
