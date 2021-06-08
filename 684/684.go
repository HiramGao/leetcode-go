package main

import "fmt"

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)

	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}
	var union func(i, j int) bool

	union = func(i, j int) bool {
		x, y := find(i), find(j)
		if x == y {
			return false
		}
		parent[x] = y
		return true
	}

	for _, v := range edges {
		if !union(v[0], v[1]) {
			return v
		}
	}
	return nil
}
func main() {
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	ans := findRedundantConnection(edges)
	fmt.Println(ans)
}
