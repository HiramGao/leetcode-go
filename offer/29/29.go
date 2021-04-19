package main

import "fmt"

//[
//	[1, 2, 3, 4],
//	[5, 6, 7, 8],
//	[9,10,11,12]
//]

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, cols               = len(matrix), len(matrix[0])
		left, right, top, bottom = 0, cols - 1, 0, rows - 1
		index                    = 0
		res                      = make([]int, rows*cols)
	)

	for left <= right && top <= bottom {
		println(index, left, right, top, bottom)
		for col := left; col <= right; col++ {
			res[index] = matrix[top][col]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			res[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for col := right - 1; col > left; col-- {
				res[index] = matrix[bottom][col]
				index++
			}
			for row := bottom; row > top; row-- {
				res[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return res
}

func main() {
	fmt.Printf("%v", spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
}
