package main

import "fmt"

type UnionFindSet struct {
	parent, size []int
}

func (s UnionFindSet) Union(x, y int) {
	x, y = s.Find(x), s.Find(y)
	if x == y {
		return
	}
	s.parent[x] = y
	s.size[y] += s.size[x]
}

func (s UnionFindSet) GetSize(x int) int {
	x = s.Find(x)
	return s.size[x]
}

func (s UnionFindSet) Find(x int) int {
	if s.parent[x] != x {
		s.parent[x] = s.Find(s.parent[x])
	}
	return s.parent[x]
}

func NewUnionFindSet(n int) *UnionFindSet {

	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFindSet{parent, size}
}

type pair struct {
	x, y int
}

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func hitBricks(grid [][]int, hits [][]int) []int {
	rows := len(grid)
	cols := len(grid[0])
	//1 按照hits敲掉所有砖块
	copyGrid := make([][]int, rows)
	for i := range copyGrid {
		copyGrid[i] = append([]int(nil), grid[i]...)
	}
	for _, hit := range hits {
		copyGrid[hit[0]][hit[1]] = 0
	}
	//2 建图，n是rows*cols，也表示虚拟为屋顶
	n := rows * cols
	set := NewUnionFindSet(n + 1)
	//下标0的一行与屋顶连接
	for i := 0; i < cols; i++ {
		if copyGrid[0][i] == 1 {
			set.Union(i, n)
		}
	}
	//其余网格进行向上向左的合并
	for i := 1; i < rows; i++ {
		for j := 0; j < cols; j++ {

			if copyGrid[i][j] == 1 {
				if copyGrid[i-1][j] == 1 {
					set.Union((i-1)*cols+j, i*cols+j)
				}
				if j > 0 && copyGrid[i][j-1] == 1 {
					set.Union(i*cols+j-1, i*cols+j)
				}
			}

		}
	}
	//3 按照hits的逆序在copy补回砖块
	res := []int{}
	for i := len(hits) - 1; i >= 0; i-- {
		x, y := hits[i][0], hits[i][1]
		if grid[x][y] == 0 {
			continue
		}
		org := set.GetSize(n)
		if x == 0 {
			set.Union(y, n)
		}
		//上下左右合并
		for _, d := range dir4 {
			if newX, newY := x+d.x, y+d.y; 0 <= newX && newX < rows && 0 <= newY && newY < cols {
				set.Union(x*cols+y, newX*cols+newY)
			}
		}
		current := set.GetSize(n)
		res = append(res, max(0, current-org-1))

		copyGrid[x][y] = 1
	}
	return res
}

func max(i, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
func main() {
	grid := [][]int{{1, 0, 0, 0}, {1, 1, 1, 0}}
	hits := [][]int{{1, 0}}
	fmt.Println(hitBricks(grid, hits))
}
