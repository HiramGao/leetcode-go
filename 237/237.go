package main

import (
	. "../ListNode"
)

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
func main() {
	node:=BuildList([]int{4,5,1,9})
	deleteNode(node.Next)
	ListPrint(node)
}
