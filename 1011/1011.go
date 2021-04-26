package main

import "sort"

func shipWithinDays(weights []int, D int) int {
	left, right := 0, 0
	for _, v := range weights {
		if v > left {
			left = v
		}
		right += v
	}

	res := left + sort.Search(right-left, func(x int) bool {
		x += left
		day := 1
		sum := 0
		for _, w := range weights {
			if sum+w > x {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= D
	})
	return res
}
func main() {
	println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}
