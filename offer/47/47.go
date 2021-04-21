package main

func maxValue1(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])
	f := make([][]int, rows)
	for i := range f {
		f[i] = make([]int, cols)
	}
	for i, rows := range grid {
		for j, v := range rows {
			up := 0
			left := 0
			if i > 0 {
				up = f[i-1][j]
			}
			if j > 0 {
				left = f[i][j-1]
			}
			f[i][j] = max(left, up) + v
		}
	}
	return f[rows-1][cols-1]
}
func maxValue2(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	cols := len(grid[0])
	f := make([]int, cols)

	for i, rows := range grid {
		for j, v := range rows {
			up := 0
			left := 0
			if i > 0 {
				up = f[j]
			}
			if j > 0 {
				left = f[j-1]
			}
			f[j] = max(left, up) + v
		}
	}
	return f[cols-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	println(maxValue2([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}
