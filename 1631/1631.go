package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

type point struct{ x, y, maxDiff int }
type hp []point

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].maxDiff < h[j].maxDiff }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(point)) }
func (h *hp) Pop() (v interface{}) {
	a := *h
	*h, v = a[:len(a)-1], a[len(a)-1]
	return
}

type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumEffortPath(heights [][]int) int {
	rows := len(heights)
	cols := len(heights[0])

	maxDiff := make([][]int, rows)
	for i := range maxDiff {
		maxDiff[i] = make([]int, cols)
		for j := range maxDiff[i] {
			maxDiff[i][j] = math.MaxInt64
		}
	}
	maxDiff[0][0] = 0

	h := &hp{{}}
	heap.Init(h)
	for {
		p := heap.Pop(h).(point)
		if p.x == rows-1 && p.y == cols-1 {
			return p.maxDiff
		}
		if maxDiff[p.x][p.y] < p.maxDiff {
			continue
		}

		for _, d := range dir4 {
			if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < rows && 0 <= y && y < cols {
				if diff := max(p.maxDiff, abs(heights[x][y]-heights[p.x][p.y])); diff < maxDiff[x][y] {
					maxDiff[x][y] = diff
					heap.Push(h, point{x, y, diff})
				}
			}
		}
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minimumEffortPath1(heights [][]int) int {

	type Edge struct {
		from, to, diff int
	}

	m, n := len(heights), len(heights[0])

	edges := make([]Edge, 0, m*n*2)
	for r, row := range heights {
		for c, h := range row {
			id := r*n + c
			if r > 0 {
				edges = append(edges, Edge{from: id, to: (r-1)*n + c, diff: abs(h - heights[r-1][c])})
			}
			if c > 0 {
				edges = append(edges, Edge{from: id, to: r*n + c - 1, diff: abs(h - heights[r][c-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].diff < edges[j].diff
	})

	parent := make([]int, n*m)

	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		for x != parent[x] {
			x, parent[x] = parent[x], parent[parent[x]]
		}
		return x
	}
	var union func(x, y int)
	union = func(x, y int) {
		x, y = find(x), find(y)
		if x == y {
			return
		}
		parent[x] = y
	}
	startId, endId := 0, m*n-1
	for _, edge := range edges {
		union(edge.from, edge.to)
		if find(startId) == find(endId) {
			return edge.diff
		}
	}

	return 0
}

func main() {

	heights := [][]int{{4, 3, 4, 10, 5, 5, 9, 2}, {10, 8, 2, 10, 9, 7, 5, 6}, {5, 8, 10, 10, 10, 7, 4, 2}, {5, 1, 3, 1, 1, 3, 1, 9}, {6, 4, 10, 6, 10, 9, 4, 6}}
	fmt.Println(minimumEffortPath1(heights))
}
