package main

import "fmt"

func permute(nums []int) (ans [][]int) {
	var (
		n         = len(nums)
		backtrack func(i int)
	)

	backtrack = func(i int) {
		if i == n {
			ans = append(ans, append([]int(nil), nums...))
		}
		for j := i; j < n; j++ {
			nums[j], nums[i] = nums[i], nums[j]
			backtrack(i + 1)
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	backtrack(0)
	return ans
}
func main() {
	fmt.Print(permute([]int{1, 2, 3}))
}
