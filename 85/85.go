package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	heights := make([]int, len(matrix[0]))
	res := 0

	for _, row := range matrix {
		for j, v := range row {
			if v == '1' {
				heights[j] += 1
			} else {
				heights[j] = 0
			}
		}
		res = max(res, largestRectangleArea(heights))
	}

	return res
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}
	stack := []int{}
	push := func(i int) {
		stack = append(stack, i)
	}
	peek := func() int {
		return stack[len(stack)-1]
	}
	pop := func() int {
		x := peek()
		stack = stack[:len(stack)-1]
		return x
	}
	res := 0
	for i := 0; i < n; i++ {

		for len(stack) != 0 && heights[peek()] > heights[i] {
			currHeight := heights[pop()]
			for len(stack) != 0 && currHeight == heights[peek()] {
				pop()
			}
			var currWidth int
			if len(stack) == 0 {
				currWidth = i
			} else {
				currWidth = i - peek() - 1
			}
			res = max(res, currWidth*currHeight)
		}
		push(i)
	}
	for len(stack) != 0 {
		currHeight := heights[pop()]
		for len(stack) != 0 && currHeight == heights[peek()] {
			pop()
		}
		var currWidth int
		if len(stack) == 0 {
			currWidth = n
		} else {
			currWidth = n - peek() - 1
		}
		res = max(res, currWidth*currHeight)
	}
	return res
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
func main() {
	matrix := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	fmt.Println(maximalRectangle(matrix))
}
