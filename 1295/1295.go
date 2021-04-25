package main

import "strconv"

func findNumbers(nums []int) int {
	res := 0
	for _, v := range nums {
		if len(strconv.Itoa(v))&1 == 0 {
			res++
		}
	}
	return res
}
func main() {
	println(findNumbers([]int{12, 345, 2, 6, 7896}))
}
