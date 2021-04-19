package main

import (
	"fmt"
	"sort"
	"strings"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return compare(nums[i], nums[j])
	})
	var str strings.Builder
	for _, v := range nums {
		str.WriteString(fmt.Sprintf("%d", v))
	}
	return str.String()
}

func compare(a, b int) bool {
	str1 := fmt.Sprintf("%d%d", a, b)
	str2 := fmt.Sprintf("%d%d", b, a)
	if str1 < str2 {
		return true
	}
	return false
}

func main() {
	println(minNumber([]int{3, 30, 34, 5, 9}))
}
