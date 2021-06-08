package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(_rob(nums[:n-1]), _rob(nums[1:]))
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
func _rob(nums []int) int {
	n := len(nums)
	a, b := nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		a, b = b, max(a+nums[i], b)
	}
	return b
}
func main() {
	fmt.Print(rob([]int{2, 3, 2}))
}
