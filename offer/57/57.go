package main

import "fmt"

func twoSum(nums []int, target int) []int {
	ahead, behind := 0, len(nums)-1

	for ahead < behind {
		sum := nums[ahead] + nums[behind]
		if sum > target {
			behind--
		} else if sum < target {
			ahead++
		} else {
			return []int{nums[ahead], nums[behind]}
		}
	}
	return []int{}
}
func main() {
	fmt.Printf("%v", twoSum([]int{2, 7, 11, 15}, 9))
}
