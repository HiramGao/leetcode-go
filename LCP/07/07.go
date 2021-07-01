package main

import "fmt"

func numWays(n int, relation [][]int, k int) int {
	edges := make([][]int, n)

	for _, v := range relation {
		x, y := v[0], v[1]
		edges[x] = append(edges[x], y)
	}
	var dfs func(from, step int)
	res := 0
	dfs = func(from, step int) {
		if step == k {
			if from == n-1 {
				res++
			}
			return
		}

		for _, to := range edges[from] {
			dfs(to, step+1)
		}
	}
	dfs(0, 0)
	return res
}
func main() {
	n := 5
	relation := [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}
	k := 3
	fmt.Println(numWays(n, relation, k))

	n = 3
	relation = [][]int{{0, 2}, {2, 1}}
	k = 2
	fmt.Println(numWays(n, relation, k))
}
