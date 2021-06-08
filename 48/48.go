package main

import "fmt"

func rotate(matrix [][]int) {

	n := len(matrix)
	//水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
func main() {

	matrix := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}

	rotate(matrix)

	fmt.Println(matrix)
}
