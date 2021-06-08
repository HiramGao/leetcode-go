package main

import "fmt"

func rotate(nums []int, k int) {
	k %= len(nums)

	reverse := func(nums []int) {
		n := len(nums)
		m := n / 2
		for i := 0; i < m; i++ {
			nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
		}
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
	nums = []int{-1, -100, 3, 99}
	rotate(nums, 2)
	fmt.Println(nums)
}
