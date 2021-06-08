package main

import (
	"fmt"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	distNums := make([]int, n)
	copy(distNums, nums)
	sort.Slice(distNums, func(i, j int) bool { return distNums[i] < distNums[j] })

	l, r := 0, n-1
	for l < r {
		if distNums[l] == nums[l] {
			l++
		} else if distNums[r] == nums[r] {
			r--
		} else {
			break
		}
	}
	if r == l {
		return 0
	}
	return r - l + 1
}

func findUnsortedSubarray1(nums []int) int {
	n := len(nums)
	stack := []int{}
	l, r := n, 0

	for i := 0; i < n; i++ {
		for len(stack) != 0 && nums[stack[len(stack)-1]] > nums[i] {
			l = min(l, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i] {
			r = max(r, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	if r-l > 0 {
		return r - l + 1
	}
	return 0
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray1([]int{2, 6, 4, 8, 10, 9, 15}))
}
