package main

import "fmt"
import "../utils"

type direction struct {
	x, y int
}

//上下左右
var dir4 = []direction{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	rows, cols := len(board), len(board[0])

	var dfs func(row, col int)

	dfs = func(row, col int) {
		if 0 <= row && row < rows && 0 <= col && col < cols && board[row][col] == 'O' {
			board[row][col] = 'A'
			for _, d := range dir4 {
				newRow, newCol := row+d.x, col+d.y
				dfs(newRow, newCol)
			}
		}
	}
	for i := 0; i < rows; i++ {
		dfs(i, 0)
		dfs(i, cols-1)
	}
	for i := 0; i < cols; i++ {
		dfs(0, i)
		dfs(rows-1, i)
	}
	for i, bytes := range board {
		for j, b := range bytes {
			if b == 'A'{
				board[i][j] = 'O'
			}else if b =='O'{
				board[i][j] = 'X'
			}
		}
	}
}
func main() {
	board := utils.Resolve2(`[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]`)
	solve(board)
	fmt.Println(board)
	fmt.Printf("%v", board)
}
