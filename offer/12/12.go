package main

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])
	wordLens := len(word)

	var dfs func(i int, j int, k int) bool

	dfs = func(i int, j int, k int) bool {
		if i < 0 || i >= rows || j < 0 || j >= cols || board[i][j] != word[k] {
			return false
		}
		if k == wordLens-1 {
			return true
		}
		board[i][j] = '0'
		nextK := k + 1
		res := dfs(i+1, j, nextK) || dfs(i-1, j, nextK) || dfs(i, j-1, nextK) || dfs(i, j+1, nextK)
		board[i][j] = word[k]
		return res
	}

	for i := range board {
		for j := range board[i] {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCEDASS"))
}
