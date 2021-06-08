package main

import "fmt"

func maxProduct(nums []int) int {
	n := len(nums)
	maxValue, minValue, ans := nums[0], nums[0], nums[0]

	for i := 1; i < n; i++ {
		num := nums[i]
		maxVal, minVal := maxValue, minValue
		maxValue = max(num, max(maxVal*num, minVal*num))
		minValue = min(num, min(maxVal*num, minVal*num))
		ans = max(ans, maxValue)
	}
	return ans
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
func min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}
func main() {
	fmt.Print(maxProduct([]int{2, 3, -2, -4}))
}
