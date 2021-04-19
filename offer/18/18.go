package main

import (
	. "../ListNode"
)
//
//func deleteNode(head *ListNode, val int) *ListNode {
//
//	node := head
//	for node!=nil {
//		if node.Val==val {
//			if node.Next == nil {
//				preNode:=head
//				for node.Val != preNode.Next.Val {
//					preNode = preNode.Next
//				}
//				preNode.Next = nil
//			}else {
//				node.Val = node.Next.Val
//				node.Next = node.Next.Next
//			}
//			break
//		}
//		node = node.Next
//	}
//
//	return head
//}

func deleteNode(head *ListNode, val int) *ListNode {
	dummy:=&ListNode{}
	dummy.Next = head
	preNode := dummy
	node := head

	for node!=nil {
		if  node.Val == val{
			preNode.Next = node.Next
			break
		}

		preNode = preNode.Next
		node = node.Next
	}
	return dummy.Next
}

func main() {
	headArray := []int{4,2,3}
	head := &ListNode{}
	for _, v := range headArray {
		InsertListNode(head, &ListNode{Val: v, Next: nil})
	}
	result := deleteNode(head.Next, 4)

	ListPrint(result)
}
