package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	hm := map[int]int{}
	res := []int{}
	for _, i := range nums1 {
		hm[i]++
	}
	for _, i := range nums2 {
		if v, ok := hm[i]; ok && v > 0 {
			res = append(res, i)
			hm[i]--
		}
	}
	return res
}
func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}
