package main

import "fmt"

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
	heights := []int{2, 4}
	fmt.Println(largestRectangleArea(heights))
}
