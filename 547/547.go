package main

import "fmt"

/**
dfs
*/
func findCircleNum(isConnected [][]int) int {

	isVis := make([]bool, len(isConnected))
	ans := 0
	var dfs func(from int)

	dfs = func(from int) {
		isVis[from] = true
		for to, con := range isConnected[from] {
			if !isVis[to] && con == 1 {
				dfs(to)
			}
		}
	}

	for i, v := range isVis {
		if !v {
			dfs(i)
			ans++
		}
	}

	return ans
}

/**
bfs
*/
func findCircleNum2(isConnected [][]int) int {

	isVis := make([]bool, len(isConnected))
	ans := 0

	for from, v := range isVis {
		if !v {
			ans++
			Q := []int{from}
			for len(Q) != 0 {
				from = Q[0]
				Q = Q[1:]
				isVis[from] = true
				for to, con := range isConnected[from] {
					if !isVis[to] && con == 1 {
						Q = append(Q, to)
					}
				}
			}
		}
	}

	return ans
}

/**
并查集
*/
func findCircleNum3(isConnected [][]int) int {

	n := len(isConnected)
	parent := make([]int, n)
	ans := 0
	for i := range parent {
		parent[i] = i
	}
	var find func(i int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}
	var union func(from, to int)
	union = func(from, to int) {
		parent[find(from)] = find(to)
	}

	for i, row := range isConnected {
		for j := i + 1; j < n; j++ {
			if row[j] == 1 {
				union(i, j)
			}
		}
	}

	for i, v := range parent {
		if i == v {
			ans++
		}
	}

	return ans
}
func main() {
	fmt.Println(findCircleNum3([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
}

//  a  b  c
//a{1, 1, 0},
//b{1, 1, 0},
//c{0, 0, 1}
