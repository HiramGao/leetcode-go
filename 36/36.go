package main

import "fmt"

func isValidSudoku(board [][]byte) bool {

	for _, rows := range board {
		hashMap := [9]int{}

		for _, v := range rows {
			if '1' <= v && v <= '9' {
				if hashMap[v-'1'] == 1 {
					return false
				}
				hashMap[v-'1'] = 1
			}
		}
	}

	for i := 0; i < 9; i = i + 1 {
		hashMap := [9]int{}
		for j := 0; j < 9; j = j + 1 {
			v := board[j][i]
			if '1' <= v && v <= '9' {
				if hashMap[v-'1'] == 1 {
					return false
				}
				hashMap[v-'1'] = 1
			}
		}
	}

	for i := 0; i < 9; i = i + 3 {
		for j := 0; j < 9; j = j + 3 {
			hashMap := [9]int{}
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					v := board[x][y]
					if '1' <= v && v <= '9' {
						if hashMap[v-'1'] == 1 {
							return false
						}
						hashMap[v-'1'] = 1
					}
				}
			}
		}
	}

	return true
}

func isValidSudoku1(board [][]byte) bool {
	rowHashMap := [9][9]int{}
	colHashMap := [9][9]int{}
	boxHashMap := [9][9]int{}
	for i, rows := range board {

		for j, v := range rows {
			if v == '.' {
				continue
			}
			boxIndex := (i/3)*3 + j/3
			index := v - '1'
			if rowHashMap[i][index] == 1 || colHashMap[j][index] == 1 || boxHashMap[boxIndex][index] == 1 {
				return false
			}
			rowHashMap[i][index] = 1
			colHashMap[j][index] = 1
			boxHashMap[boxIndex][index] = 1
		}
	}

	return true
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

	fmt.Println(isValidSudoku1(board))
}
