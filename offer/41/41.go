package main

import (
	"container/heap"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i int, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Peek() int {
	return (*h)[0]
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Peek() int {
	return (*h)[0]
}

type MedianFinder struct {
	Right *MinHeap
	Left  *MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	Min := &MinHeap{}
	heap.Init(Min)
	Max := &MaxHeap{}
	heap.Init(Max)

	return MedianFinder{
		Right: Min,
		Left:  Max,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.Right.Len() == this.Left.Len() {
		if this.Left.Len() > 0 && this.Left.Peek() > num {
			heap.Push(this.Left, num)
			num = heap.Pop(this.Left).(int)
		}
		heap.Push(this.Right, num)
	} else {
		if this.Right.Len() > 0 && this.Right.Peek() < num {
			heap.Push(this.Right, num)
			num = heap.Pop(this.Right).(int)
		}
		heap.Push(this.Left, num)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.Right.Len() != this.Left.Len() {
		return float64(this.Right.Peek())
	} else {
		return (float64(this.Left.Peek()) + float64(this.Right.Peek())) / 2
	}
}
func main() {
	obj := Constructor()
	obj.AddNum(2)

	println(obj.FindMedian())
	obj.AddNum(3)
	println(obj.FindMedian())
}
