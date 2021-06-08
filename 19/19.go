package main

import (
	. "../ListNode"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy.Next

	for fast != nil {
		if n>0{
			n--
		}else {
			slow = slow.Next
		}
		fast = fast.Next

	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
func main() {
	head := BuildList([]int{1})
	head = removeNthFromEnd(head, 1)
	ListPrint(head)
}
