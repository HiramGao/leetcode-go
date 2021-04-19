package main

import (
	. "../ListNode"
)

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head==nil{
		return head
	}
	slow,fast := head,head

	for nil != fast {
		if k<=0{
			slow = slow.Next
		}
		fast = fast.Next
		k--
	}
	return slow
}

func main() {
	headArray := []int{1,2,3,4,5,6}
	head := &ListNode{}
	for _, v := range headArray {
		InsertListNode(head, &ListNode{Val: v, Next: nil})
	}
	result := getKthFromEnd(head.Next,2)
	ListPrint(result)
}
