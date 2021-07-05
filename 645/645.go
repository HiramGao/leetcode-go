package main

import "fmt"

func findErrorNums(nums []int) []int {
	ans := make([]int, 2)
	hm := map[int]int{}
	for _, v := range nums {
		hm[v]++
	}
	for i := 1; i <= len(nums); i++ {
		if num := hm[i]; num == 2 {
			ans[0] = i
		} else if num == 0 {
			ans[1] = i
		}
	}
	return ans
}

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
	fmt.Println(findErrorNums([]int{1, 1}))
}
