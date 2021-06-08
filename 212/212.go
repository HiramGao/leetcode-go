package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	word     string
}
type direction struct {
	row, col int
}

//上下左右
var directions = []direction{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	m, n := len(board), len(board[0])
	root := &TrieNode{[26]*TrieNode{}, ""}

	for _, word := range words {
		node := root
		for _, b := range []byte(word) {
			if oldNode := node.children[b-'a']; oldNode != nil {
				node = oldNode
			} else {
				newNode := &TrieNode{[26]*TrieNode{}, ""}
				node.children[b-'a'] = newNode
				node = newNode
			}
		}
		node.word = word
	}

	var backtracking func(row, col int, node *TrieNode)

	backtracking = func(row, col int, currNode *TrieNode) {
		if currNode == nil {
			return
		}
		currByte := board[row][col]
		if currNode.word != "" {
			res = append(res, currNode.word)
			currNode.word = ""
		}
		board[row][col] = '#'

		for _, d := range directions {
			newRow, newCol := row+d.row, col+d.col
			if 0 <= newRow && newRow < m && 0 <= newCol && newCol < n {
				nextByte := board[newRow][newCol]
				if nextByte != '#' {
					if nextNode := currNode.children[nextByte-'a']; nextNode != nil {
						backtracking(newRow, newCol, nextNode)
					}
				}

			}

		}
		board[row][col] = currByte
	}

	for i, bytes := range board {
		for j, b := range bytes {
			if node := root.children[b-'a']; node != nil {
				backtracking(i, j, node)
			}
		}
	}
	return res
}

func main() {
	board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	words := []string{"oath", "pea", "eat", "rain", "oathi", "oathk", "oathf", "oate", "oathii", "oathfi", "oathfii"}
	fmt.Println(findWords(board, words))
}
