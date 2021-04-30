package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		num := (nums[i] - 1) % n
		nums[num] += n
	}
	for i, num := range nums {
		if num <= n {
			res = append(res, i+1)
		}
	}
	return res
}
func main() {
	res := findDisappearedNumbers([]int{1, 1})
	fmt.Printf("%v", res)
}
