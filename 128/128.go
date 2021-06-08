package main

import "fmt"

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	res := 0

	for num := range numSet {
		if !numSet[num-1] {
			currNum := num
			currLen := 1
			for numSet[currNum+1] {
				currNum++
				currLen++
			}
			if res < currLen {
				res = currLen
			}
		}
	}
	return res
}
func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
