package main

import (
	. "../ListNode"
)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil{
		return l2
	}else if l2==nil{
		return l1
	}
	dummy := &ListNode{}

	if l1.Val < l2.Val{
		dummy.Next  = l1
		dummy.Next.Next = mergeTwoLists(l1.Next,l2)
	}else{
		dummy.Next  = l2
		dummy.Next.Next = mergeTwoLists(l1,l2.Next)
	}
	return dummy.Next
}



func main() {
	array1 := []int{1, 2, 3, 4, 5, 6}
	head1 := &ListNode{}
	for _, v := range array1 {
		InsertListNode(head1, &ListNode{Val: v, Next: nil})
	}
	array2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	head2 := &ListNode{}
	for _, v := range array2 {
		InsertListNode(head2, &ListNode{Val: v, Next: nil})
	}
	result := mergeTwoLists(head1.Next, head2.Next)
	ListPrint(result)
}
