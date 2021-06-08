package main

import (
	. "../../ListNode"
)

//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	newNode := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//	return newNode
//}
func reverseList(head *ListNode) *ListNode {
	var preNode *ListNode;
	curr:=head
	for curr!=nil{
		next := curr.Next
		curr.Next = preNode
		preNode = curr
		curr = next
	}
	return preNode
}

func main() {
	headArray := []int{1, 2, 3, 4, 5}
	head := &ListNode{}
	for _, v := range headArray {
		InsertListNode(head, &ListNode{Val: v, Next: nil})
	}
	//ListPrint(head)
	ListPrint(reverseList(head.Next))
}
