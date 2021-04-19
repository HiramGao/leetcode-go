package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func getLeastNumbers1(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	sort.Ints(arr)
	return arr[:k]
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i int, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Peek() int {
	return (*h)[0]
}

func getLeastNumbers2(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	h := &IntHeap{}
	res := []int{}
	heap.Init(h)

	for _, v := range arr {
		if h.Len() < k {
			heap.Push(h, v)
			continue
		}
		if h.Peek() > v {
			heap.Pop(h)
			heap.Push(h, v)
		}
	}

	for i := 0; i < k; i++ {
		res = append([]int{heap.Pop(h).(int)}, res...)
	}

	return res
}
func main() {
	//[0,0,1,1,2,2,2,3]
	fmt.Printf("%v", getLeastNumbers2([]int{3, 2, 1}, 2))
}
