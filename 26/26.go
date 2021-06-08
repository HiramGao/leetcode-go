package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, 0
	for r < len(nums) {
		if nums[l] == nums[r] {
			r++
			continue
		}
		nums[l+1] = nums[r]
		l += 1
		r++
	}
	return l + 1
}
func main() {
	nums := []int{0, 1, 2, 3}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
