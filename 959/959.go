package main

import "fmt"

func regionsBySlashes(grid []string) int {
	n := len(grid)
	count := 4 * n * n
	parent := make([]int, count)
	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	var union func(x, y int)
	find = func(x int) int {
		if x != parent[x] {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union = func(from, to int) {
		x, y := find(from), find(to)
		if x == y {
			return
		} else if x < y {
			parent[y] = x
		} else {
			parent[x] = y
		}
		count--
	}

	for i, s := range grid {

		for j, b := range []byte(s) {
			index := 4 * (i*n + j)
			if b == '/' {
				union(index, index+3)
				union(index+1, index+2)
			} else if b == '\\' {
				union(index, index+1)
				union(index+2, index+3)
			} else {
				union(index, index+1)
				union(index+1, index+2)
				union(index+2, index+3)
			}
			//向右合并
			if j+1 < n {
				union(index+1, 4*(i*n+j+1)+3)
			}
			if i+1 < n {
				union(index+2, 4*((i+1)*n+j))
			}
			fmt.Println(parent)
		}
	}

	return count
}

func main() {
	grid := []string{"  \\", "/\\/", "/\\\\"}
	fmt.Println(regionsBySlashes(grid))
}
