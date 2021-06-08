package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
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

func furthestBuilding(heights []int, bricks int, ladders int) int {
	heightsHeap := &IntHeap{}
	heap.Init(heightsHeap)

	for i := 1; i < len(heights); i++ {
		diff := heights[i] - heights[i-1]
		if diff <= 0 {
			continue
		}
		heap.Push(heightsHeap, diff)
		if heightsHeap.Len() > ladders {
			dif := heap.Pop(heightsHeap).(int)
			if dif > bricks {
				return i - 1
			}
			bricks -= dif
		}
	}
	return len(heights) - 1
}
func main() {
	fmt.Println(furthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2))
}

//4
//4 2
//4 2(5)7
//4 2(5)7 6
//4 2(5)7 6(3)9
//4 2(5)7 6(3)9(5)14
