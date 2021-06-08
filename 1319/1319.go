package main

import "fmt"

func makeConnected(n int, connections [][]int) int {

	if len(connections) < n-1 {
		return -1
	}
	need := n
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if x == parent[x] { // 当节点x的父节点是它自己，返回它本身
			return x
		}
		parent[x] = find(parent[x]) // 递归一直找父亲，直到遇到递归出口
		return parent[x]

	}

	var union func(int, int)
	union = func(from int, to int) {
		x, y := find(from), find(to)
		if x != y {
			parent[x] = y
			need--
		}

	}

	for _, connection := range connections {
		union(connection[0], connection[1])
	}

	return need - 1
}

func makeConnected1(n int, connections [][]int) int {

	if len(connections) < n-1 {
		return -1
	}
	G := make([][]int, n)
	for _, connection := range connections {
		from, to := connection[0], connection[1]
		G[to] = append(G[to], from)
		G[from] = append(G[from], to)
	}
	isVis := make([]bool, n)
	var dfs func(int)

	dfs = func(from int) {
		isVis[from] = true
		for _, to := range G[from] {
			if !isVis[to] {
				dfs(to)
			}
		}
	}

	res := 0
	for i, vis := range isVis {
		if !vis {
			res++
			dfs(i)
		}
	}

	return res - 1
}

func main() {
	fmt.Println(makeConnected1(12, [][]int{{1, 5}, {1, 7}, {1, 2}, {1, 4}, {3, 7}, {4, 7}, {3, 5}, {0, 6}, {0, 1}, {0, 4}, {2, 6}, {0, 3}, {0, 2}}))
}
