package main

import (
	"fmt"
	"sort"
)

func minDays(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if n < m*k {
		return -1
	}

	r := 0

	for _, v := range bloomDay {

		if v > r {
			r = v
		}
	}

	return sort.Search(r, func(days int) bool {
		fmt.Println(days)
		adjacentFlower, bouquets := 0, 0
		for _, v := range bloomDay {
			if v > days {
				adjacentFlower = 0
			} else {
				adjacentFlower++
				if adjacentFlower == k {
					bouquets++
					adjacentFlower = 0
				}
			}
		}
		return bouquets >= m
	})
}
func main() {
	fmt.Print(minDays([]int{7, 7, 7, 7, 12, 7, 7}, 2, 3))
}
