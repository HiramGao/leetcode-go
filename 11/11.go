package main

import "fmt"

func maxArea(height []int) int {
	res := 0

	for left, right := 0, len(height)-1; left < right; {
		width := right - left
		leftHeight, rightHeight := height[left], height[right]
		area := width * min(leftHeight, rightHeight)
		res = max(res, area)
		if leftHeight < rightHeight {
			left++
		} else {
			right--
		}
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
