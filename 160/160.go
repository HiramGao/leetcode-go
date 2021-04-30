package main

import (
	. "../offer/ListNode"
)

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB==nil{
		return nil
	}

	nodeA, nodeB := headA, headB

	for nodeA != nodeB {
		if nodeA == nil {
			nodeA = headB
		}else {
			nodeA = nodeA.Next
		}
		if nodeB ==nil{
			nodeB = headA
		}else {
			nodeB = nodeB.Next
		}
	}
	return nodeA
}
