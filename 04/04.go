package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	totalLeft := (m + n + 1) / 2
	// nums1中 [0:m] 查找分割线
	// 使得 nums1[i-1] <= nums2[j] && nums2[j-1]<=nums1[i]
	left, right := 0, m

	for left < right {
		i := left + (right-left+1)/2
		j := totalLeft - i

		if nums1[i-1] > nums2[j] {
			right = i - 1
		} else {
			left = i
		}
	}
	i, j := left, totalLeft-left
	nums1LMax, nums1RMin := math.MinInt64, math.MaxInt64
	nums2LMax, nums2RMin := math.MinInt64, math.MaxInt64
	if i != 0 {
		nums1LMax = nums1[i-1]
	}
	if i != m {
		nums1RMin = nums1[i]
	}
	if j != 0 {
		nums2LMax = nums2[j-1]
	}
	if j != n {
		nums2RMin = nums2[j]
	}
	if (m+n)%2 == 1 {
		return float64(max(nums1LMax, nums2LMax))
	} else {
		return (float64(max(nums1LMax, nums2LMax)) + float64(min(nums1RMin, nums2RMin))) / 2
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	fmt.Print(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
