package main

import (
	"fmt"
	"sort"
)

func searchMatrix(matrix [][]int, target int) bool {

	row := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)

	return col < len(matrix[row]) && matrix[row][col] == target
}
func main() {
	m := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	t := 12
	fmt.Println(searchMatrix(m, t))
}
