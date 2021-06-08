package main

import (
	"fmt"
)

func canJump(nums []int) bool {

	n := len(nums)
	maxStep := 0

	for i, num := range nums {
		if maxStep >= i {
			if num+i > maxStep {
				maxStep = num + i
			}
			if maxStep >= n-1 {
				return true
			}
		}
	}
	return false
}
func main() {
	fmt.Print(canJump([]int{3, 2, 1, 0, 4}))
}
