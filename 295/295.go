package main

import (
	"container/heap"
	"fmt"
)

type MinHeap struct {
	hp []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (m MinHeap) Len() int { return len(m.hp) }

func (m *MinHeap) Less(i, j int) bool { return m.hp[i] < m.hp[j] }

func (m *MinHeap) Swap(i, j int) { m.hp[i], m.hp[j] = m.hp[j], m.hp[i] }

func (m *MinHeap) Push(x interface{}) { m.hp = append(m.hp, x.(int)) }

func (m *MinHeap) Pop() interface{} {
	x := m.hp[m.Len()-1]
	m.hp = m.hp[:m.Len()-1]
	return x
}

func (m MinHeap) Top() int { return m.hp[0] }

type MaxHeap struct{ hp []int }

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}
func (m MaxHeap) Len() int { return len(m.hp) }

func (m *MaxHeap) Less(i, j int) bool { return m.hp[i] > m.hp[j] }

func (m *MaxHeap) Swap(i, j int) { m.hp[i], m.hp[j] = m.hp[j], m.hp[i] }

func (m *MaxHeap) Push(x interface{}) { m.hp = append(m.hp, x.(int)) }

func (m *MaxHeap) Pop() interface{} {
	x := m.hp[m.Len()-1]
	m.hp = m.hp[:m.Len()-1]
	return x
}

func (m MaxHeap) Top() int { return m.hp[0] }

type MedianFinder struct {
	maxH *MaxHeap
	minH *MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	maxH, minH := NewMaxHeap(), NewMinHeap()
	heap.Init(maxH)
	heap.Init(minH)
	return MedianFinder{maxH: maxH, minH: minH}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxH, num)
	heap.Push(this.minH, heap.Pop(this.maxH))
	for this.maxH.Len() < this.minH.Len() {
		heap.Push(this.maxH, heap.Pop(this.minH))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxH.Len() > this.minH.Len() {
		return float64(this.maxH.Top())
	}
	return float64(this.maxH.Top()+this.minH.Top()) * 0.5
}

func main() {

	mf := Constructor()
	mf.AddNum(-1)
	fmt.Println(mf.FindMedian())
	mf.AddNum(-2)
	fmt.Println(mf.FindMedian())
	mf.AddNum(-3)
	fmt.Println(mf.FindMedian())
	mf.AddNum(-4)
	fmt.Println(mf.FindMedian())
	mf.AddNum(-5)
	fmt.Println(mf.FindMedian())
	mf.AddNum(-6)
	fmt.Println(mf.FindMedian())
}
