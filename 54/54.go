package main

import "fmt"

func spiralOrder(matrix [][]int) []int {

	var (
		rows, cols               = len(matrix), len(matrix[0])
		ans                      = make([]int, rows*cols)
		index                    = 0
		left, top, right, bottom = 0, 0, cols - 1, rows - 1
	)

	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			ans[index] = matrix[top][col]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			ans[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for col := right - 1; col > left; col-- {
				ans[index] = matrix[bottom][col]
				index++
			}
			for row := bottom; row > top; row-- {
				ans[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}

	return ans
}
func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))
	matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix))
	matrix = [][]int{{1}}
	fmt.Println(spiralOrder(matrix))
}
