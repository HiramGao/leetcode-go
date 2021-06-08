package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {

	var (
		n          = len(nums)
		small, mid = math.MaxInt64, math.MaxInt64
	)
	if n < 3 {
		return false
	}

	for _, num := range nums {
		if num <= small {
			small = num
		} else if num <= mid {
			mid = num
		} else if num > mid {
			return true
		}
	}

	return false
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(increasingTriplet(nums))

	nums = []int{5, 4, 3, 2, 1}
	fmt.Println(increasingTriplet(nums))

	nums = []int{2, 1, 5, 0, 4, 6}
	fmt.Println(increasingTriplet(nums))
}
