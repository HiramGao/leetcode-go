package main

import (
	. "../ListNode"
)

func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}
	return pre
}
func reverseList(head *ListNode) *ListNode {
	if  head == nil || head.Next==nil{
		return head
	}
	newhead := reverseList(head.Next)
	head.Next.Next = head
	head.Next=nil
	return newhead
}
func main() {

}
