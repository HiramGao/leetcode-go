package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	rows := len(matrix)

	if rows == 0 {
		return false
	}

	row := 0
	column := len(matrix[0]) - 1

	for row < rows && column >= 0 {
		println(matrix[row][column], row, column)
		if matrix[row][column] == target {
			return true
		} else if matrix[row][column] > target {
			column--
		} else {
			row++
		}
	}

	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
		{31, 32, 40, 41, 60},
	}
	target := 18
	result := findNumberIn2DArray(matrix, target)
	println(result)
}
