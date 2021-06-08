package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := range result {
		result[i] = 1
	}

	left, right := 1, 1

	for i := range nums {
		result[i] *= left
		left *= nums[i]

		result[n-i-1] *= right
		right *= nums[n-i-1]
	}
	return result
}
func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
