package main

import (
	"sort"
)

func isStraight1(nums []int) bool {
	var set = map[int]bool{}
	maxV, minV := 0, 14
	for _, v := range nums {
		if v == 0 {
			continue
		}
		maxV = max(maxV, v)
		minV = min(minV, v)
		if set[v] {
			return false
		}
		set[v] = true
	}
	return maxV-minV < 5
}

func max(v int, v2 int) int {
	if v > v2 {
		return v
	}
	return v2
}
func min(v int, v2 int) int {
	if v < v2 {
		return v
	}
	return v2
}

func isStraight2(nums []int) bool {
	joker := 0
	sort.Ints(nums)
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[joker] < 5
}

func main() {
	println(isStraight2([]int{1, 0, 3, 0, 5}))
}
