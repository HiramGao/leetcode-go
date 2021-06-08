package main

import "fmt"

type direction struct {
	row, col int
}

//上下左右
var directions = []direction{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {

	rows, cols, n := len(board), len(board[0]), len(word)

	isVis := make([][]bool, rows)
	for i := 0; i < len(isVis); i++ {
		isVis[i] = make([]bool, cols)
	}

	var backtrack func(row, col, l int) bool
	backtrack = func(row, col, l int) bool {
		if board[row][col] != word[l] {
			return false
		}
		if l == n-1 {
			return true
		}
		isVis[row][col] = true
		defer func() {
			isVis[row][col] = false
		}()
		for _, direction := range directions {
			if newRow, newCol := row+direction.row, col+direction.col; 0 <= newRow && newRow < rows && 0 <= newCol && newCol < cols && !isVis[newRow][newCol] {
				if backtrack(newRow, newCol, l+1) {
					return true
				}
			}
		}
		return false
	}

	for i, b := range board {
		for j := range b {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}
func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
}
