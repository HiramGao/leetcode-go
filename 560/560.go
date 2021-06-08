package main

import "fmt"

func subarraySum(nums []int, k int) int {
	ans := 0
	hashMap := map[int]int{0: 1}
	prevCount := 0
	for _, num := range nums {
		prevCount += num
		ans += hashMap[prevCount-k]
		hashMap[prevCount] += 1
	}

	return ans
}
func main() {
	fmt.Println(subarraySum([]int{1}, 0))
}
