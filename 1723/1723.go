package main

import (
	"fmt"
	"sort"
)

func minimumTimeRequired(jobs []int, k int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(jobs)))
	var (
		n = len(jobs)
		l = jobs[0]
		r = 0
	)

	for _, job := range jobs {
		r += job
	}

	return l + sort.Search(r-l, func(limit int) bool {
		limit += l
		workloads := make([]int, k)
		var backtrace func(index int) bool

		backtrace = func(index int) bool {
			if index == n {
				return true
			}
			cur := jobs[index]
			for i := 0; i < k; i++ {
				if workloads[i]+cur <= limit {
					workloads[i] += cur
					if backtrace(index + 1) {
						return true
					}
					workloads[i] -= cur
				}
				if workloads[i] == 0 || workloads[i]+cur == limit {
					break
				}
			}

			return false
		}

		return backtrace(0)
	})
}
func main() {
	fmt.Print(minimumTimeRequired([]int{15, 2, 3, 2, 3}, 2))
}
