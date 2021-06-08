package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	for i, num := range nums {
		n ^= i ^ num
	}
	return n
}
func main() {
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
