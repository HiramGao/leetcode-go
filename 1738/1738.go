package main

import (
	"fmt"
	"sort"
)

func kthLargestValue(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	res := []int{}
	prev := make([][]int, m+1)
	for i := range prev {
		prev[i] = make([]int, n+1)
	}
	for i, ma := range matrix {
		for j, v := range ma {
			prev[i+1][j+1] = prev[i+1][j] ^ prev[i][j+1] ^ prev[i][j] ^ v
			res = append(res, prev[i+1][j+1])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(res)))
	return res[k-1]
}
func main() {
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 4))
}
