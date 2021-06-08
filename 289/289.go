package main

import "fmt"

func gameOfLife(board [][]int) {

	rows, cols := len(board), len(board[0])
	neighbors := []int{0, 1, -1}
	roundLiveCount := func(x int, y int) int {
		liveNeighbors := 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if neighbors[i] == 0 && neighbors[j] == 0 {
					continue
				}
				newX, newY := x+neighbors[i], y+neighbors[j]
				if 0 <= newX && newX < rows && 0 <= newY && newY < cols && abs(board[newX][newY]) == 1 {
					liveNeighbors++
				}
			}
		}
		return liveNeighbors
	}

	for i, rows := range board {
		for j, cellLive := range rows {
			aroundLiv := roundLiveCount(i, j)
			if cellLive == 1 && (aroundLiv < 2 || aroundLiv > 3) {
				board[i][j] = -1 //之前是活着现在已经死了
			}
			if cellLive == 0 && aroundLiv == 3 {
				board[i][j] = 2 //之前是死 现在是活着
			}
		}
	}

	for i, rows := range board {
		for j, cellLive := range rows {
			if cellLive == -1 {
				board[i][j] = 0
			} else if cellLive == 2 {
				board[i][j] = 1
			}
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
	fmt.Printf("%v", board)
}
