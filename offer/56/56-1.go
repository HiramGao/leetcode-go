package main

import "fmt"

func singleNumbers(nums []int) []int {
	count := 0
	res := make([]int, 2)
	for _, v := range nums {
		count ^= v
	}
	count = count & (-count)
	println(count)
	for _, v := range nums {
		if v&count == 0 {
			res[1] ^= v
		} else {
			res[0] ^= v
		}
	}
	return res
}

func main() {
	fmt.Printf("%v", singleNumbers([]int{2, 4, 3, 6, 3, 2, 5, 5}))
}
