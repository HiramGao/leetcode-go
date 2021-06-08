package main

import "fmt"

func solveSudoku(board [][]byte) {
	rowHashMap := [9][9]bool{}
	colHashMap := [9][9]bool{}
	boxHashMap := [9][9]bool{}
	space := [][2]int{}
	for i, rows := range board {

		for j, v := range rows {
			if v == '.' {
				space = append(space, [2]int{i, j})
				continue
			}
			boxIndex := (i/3)*3 + j/3
			index := v - '1'
			rowHashMap[i][index] = true
			colHashMap[j][index] = true
			boxHashMap[boxIndex][index] = true
		}
	}
	var dfs func(p int) bool
	dfs = func(p int) bool {
		if p == len(space) {
			return true
		}
		x, y := space[p][0], space[p][1]
		boxIndex := (x/3)*3 + y/3
		for d := 0; d <= 8; d++ {
			if !(rowHashMap[x][d] || colHashMap[y][d] || boxHashMap[boxIndex][d]) {
				rowHashMap[x][d] = true
				colHashMap[y][d] = true
				boxHashMap[boxIndex][d] = true
				board[x][y] = byte(d + '1')
				if dfs(p + 1) {
					return true
				}
				rowHashMap[x][d] = false
				colHashMap[y][d] = false
				boxHashMap[boxIndex][d] = false
			}
		}

		return false
	}

	dfs(0)
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
	fmt.Println(board)
}
