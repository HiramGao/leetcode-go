package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	hashMap := map[int]int{}

	for _, v := range nums1 {
		for _, w := range nums2 {
			hashMap[v+w]++
		}
	}
	res := 0
	for _, v := range nums3 {
		for _, w := range nums4 {
			res += hashMap[-(v + w)]
		}
	}
	return res
}
func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(fourSumCount(A, B, C, D))
}
