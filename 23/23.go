package main

import (
	. "../ListNode"
	"container/heap"
	"fmt"
)

type IntHeap []*ListNode

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i int, j int) bool {
	return (*h[i]).Val < (*h[j]).Val
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	h := &IntHeap{}
	heap.Init(h)
	for _, list := range lists {
		node := list
		for node != nil {
			heap.Push(h, node)
			node = node.Next
		}
	}
	var dummp = &ListNode{}
	var node *ListNode
	for h.Len() != 0 {
		if node == nil {
			dummp.Next = heap.Pop(h).(*ListNode)
			node = dummp.Next
		} else {
			node.Next = heap.Pop(h).(*ListNode)
			node = node.Next
		}
	}
	if node != nil {
		node.Next = nil
	}

	return dummp.Next
}

func mergeKList1(lists []*ListNode) *ListNode {
	n := len(lists)
	var merge func(int, int) *ListNode
	var mergeTwoLists func(a, b *ListNode) *ListNode

	merge = func(l, r int) *ListNode {
		if l == r {
			return lists[l]
		}
		if l > r {
			return nil
		}
		mid := (l + r) >> 1

		return mergeTwoLists(merge(l, mid), merge(mid+1, r))
	}
	mergeTwoLists = func(a, b *ListNode) *ListNode {
		if a == nil {
			return b
		}
		if b == nil {
			return a
		}
		dummy := &ListNode{}
		var node *ListNode
		nodeA, nodeB := a, b
		for nodeA != nil && nodeB != nil {
			var curNode *ListNode
			if nodeA.Val < nodeB.Val{
				curNode = nodeA
				nodeA = nodeA.Next
			}else {
				curNode = nodeB
				nodeB = nodeB.Next
			}
			if node == nil{
				dummy.Next = curNode
				node = dummy.Next
			}else {
				node.Next = curNode
				node = node.Next
			}
		}
		if nodeA ==nil && node!=nil{
			node.Next = nodeB
		}
		if nodeB ==nil && node!=nil{
			node.Next = nodeA
		}
		return dummy.Next
	}

	return merge(0, n-1)
}
func main() {
	list1 := BuildList([]int{1,4,5})
	list2 := BuildList([]int{1,3,4})
	list3 := BuildList([]int{2,6})
	list := []*ListNode{list1, list2, list3}
	ans := mergeKList1(list)
	arr := ListPrint(ans)
	fmt.Println(arr)
}
