package main

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) {
	n := len(nums)
	right, mid := n, (n+1)>>1
	copyNums := make([]int, n)
	copy(copyNums, nums)
	sort.Ints(copyNums)
	for left := range nums {
		if left&1 == 1 {
			right--
			nums[left] = copyNums[right]
		} else {
			mid--
			nums[left] = copyNums[mid]
		}
	}
}
func main() {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	fmt.Println(nums)
	nums = []int{1, 3, 2, 2, 3, 1}
	wiggleSort(nums)
	fmt.Println(nums)
}
