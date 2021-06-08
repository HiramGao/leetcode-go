package main

import "fmt"

func trap(height []int) int {

	stack := []int{}
	res := 0

	for i, v := range height {

		for len(stack) > 0 && v > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			w := i - left - 1
			h := min(height[left], v) - height[top]
			res += w * h

		}
		stack = append(stack, i)
	}

	return res
}

func min(i int, v int) int {
	if i < v {
		return i
	}
	return v
}

func trap2(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	res := 0

	leftMax := make([]int, n)
	rightMax := make([]int, n)
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i, v := range height {
		res += min(leftMax[i], rightMax[i]) - v
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func trap3(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	res := 0
	left, right := 0, n-1
	lMax, rMax := 0, 0
	for left < right {
		lMax = max(lMax, height[left])
		rMax = max(rMax, height[right])
		if height[left] < height[right] {
			res += lMax - height[left]
			left++
		} else {
			res += rMax - height[right]
			right--
		}
	}

	return res
}

func main() {
	fmt.Println(trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
