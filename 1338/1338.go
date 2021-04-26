package main

import (
	"sort"
)

func minSetSize(arr []int) int {
	hashMap := map[int]int{}
	var sortArr []int
	for _, v := range arr {
		hashMap[v] += 1
	}
	for i := range hashMap {
		sortArr = append(sortArr, i)

	}
	sort.Slice(sortArr, func(i, j int) bool {
		return hashMap[sortArr[i]] > hashMap[sortArr[j]]
	})
	half := len(arr) / 2
	res := 0
	for _, v := range sortArr {
		half -= hashMap[v]
		res++
		if half <= 0 {
			return res
		}
	}
	return 0
}
func main() {
	println(minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))
}
