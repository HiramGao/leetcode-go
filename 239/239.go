package main

import (
	"container/heap"
	"fmt"
	"sort"
)

var copyNums []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return copyNums[h.IntSlice[i]] > copyNums[h.IntSlice[j]] }

func (h *hp) Push(x interface{}) { h.IntSlice = append(h.IntSlice, x.(int)) }

func (h *hp) Pop() interface{} {
	v := h.IntSlice[h.IntSlice.Len()-1]
	h.IntSlice = h.IntSlice[:h.IntSlice.Len()-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	copyNums = nums
	h := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		h.IntSlice[i] = i
	}
	heap.Init(h)
	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[h.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(h, i)
		for h.IntSlice[0] <= i-k {
			heap.Pop(h)
		}
		ans = append(ans, nums[h.IntSlice[0]])
	}

	return ans
}

func maxSlidingWindow1(nums []int, k int) []int {
	q := []int{}

	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		fmt.Println(q)
	}
	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}
func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
	fmt.Println(maxSlidingWindow1(nums, 3))
}
