package main

import (
	"fmt"
	"sort"
)

func kClosest1(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		q, p := points[i], points[j]
		return q[0]*q[0]+q[1]*q[1] < p[0]*p[0]+p[1]*p[1]
	})
	return points[:k]
}
func kClosest2(points [][]int, k int) [][]int {

	var quickSort func(left, right int)
	quickSort = func(left, right int) {
		if left == right {
			return
		}
		pivot := points[right]
		lessCount := left
		for i := left; i < right; i++ {
			if less(points[i], pivot) {
				points[i], points[lessCount] = points[lessCount], points[i]
				lessCount++
			}
		}
		points[right], points[lessCount] = points[lessCount], points[right]
		if lessCount+1 == k {
			return
		} else if lessCount+1 < k {
			quickSort(lessCount+1, right)
		} else {
			quickSort(left, lessCount-1)
		}
	}
	quickSort(0, len(points)-1)
	return points[:k]
}

func less(q []int, p []int) bool {
	return q[0]*q[0]+q[1]*q[1] < p[0]*p[0]+p[1]*p[1]
}
func main() {
	fmt.Printf("%v", kClosest2([][]int{
		{3, 3}, {5, -1}, {-2, 4},
	}, 2))
}
