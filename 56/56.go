package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) (ans [][]int) {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for _, interval := range intervals {
		l, r := interval[0], interval[1]

		if len(ans) == 0 || ans[len(ans)-1][1] < l {
			ans = append(ans, interval)
		} else {
			if r > ans[len(ans)-1][1] {
				ans[len(ans)-1][1] = r
			}

		}
	}

	return
}
func main() {
	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
}
