package main

import (
	"fmt"
	"strings"
)

const target = "123450"

var neighbors = [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

type pair struct {
	status string
	step   int
}

func slidingPuzzle(board [][]int) int {

	startByte := []byte{}
	for _, row := range board {
		for _, b := range row {
			startByte = append(startByte, '0'+byte(b))
		}
	}
	start := string(startByte)
	if start == target {
		return 0
	}

	nextStatus := func(status string) (next []string) {
		s := []byte(status)
		x := strings.IndexByte(status, '0')
		for _, y := range neighbors[x] {
			s[x], s[y] = s[y], s[x]
			next = append(next, string(s))
			s[x], s[y] = s[y], s[x]
		}
		return
	}

	queue := []pair{{start, 0}}
	visited := map[string]bool{start: true}
	for len(queue) > 0 {
		currentStatus := queue[0]
		queue = queue[1:]
		for _, next := range nextStatus(currentStatus.status) {
			if !visited[next] {
				if next == target {
					return currentStatus.step + 1
				}
				visited[next] = true
				queue = append(queue, pair{next, currentStatus.step + 1})
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(slidingPuzzle([][]int{{1, 2, 3}, {4, 0, 5}}))
	fmt.Println(slidingPuzzle([][]int{{1, 2, 3}, {5, 4, 0}}))
	fmt.Println(slidingPuzzle([][]int{{4, 1, 2}, {5, 0, 3}}))
	fmt.Println(slidingPuzzle([][]int{{3, 2, 4}, {1, 5, 0}}))
}
