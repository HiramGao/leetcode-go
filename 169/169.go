package main

import "fmt"

func majorityElement(nums []int) int {
	count, candidate := 0, 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if candidate == num {
			count += 1
		} else {
			count -= 1
		}
	}
	return candidate
}
func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
