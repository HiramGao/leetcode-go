package main

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l] < target {
			l++
		} else if nums[r] > target {
			r--
		} else {
			break
		}
	}
	if nums[l] != target || nums[r] != target {
		return []int{-1, -1}
	}
	return []int{l, r}
}

func searchRange1(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}
func main() {
	nums := []int{5, 6, 6, 6, 6, 6, 7}
	fmt.Println(searchRange1(nums, 6))
}
