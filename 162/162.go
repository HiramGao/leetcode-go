package main

import "fmt"

func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	for i, num := range nums {
		var prev, next int
		if i == 0 {
			prev = nums[len(nums)-1]
		} else {
			prev = nums[i-1]
		}
		if i == n-1 {
			next = nums[0]
		} else {
			next = nums[i+1]
		}

		if prev < num && num > next {
			return i
		}

	}
	return 0
}
func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
