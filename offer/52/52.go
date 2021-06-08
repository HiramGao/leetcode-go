package main

import (
	. "../../ListNode"
)

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	headALength := getListHead(headA)
	headBLength := getListHead(headB)

	longList := headA
	shortList := headB
	diff := headALength - headBLength

	if headALength < headBLength {
		longList = headB
		shortList = headA
		diff = -diff
	}
	for diff != 0 {
		longList = longList.Next
		diff--
	}
	for longList != nil && shortList != nil {
		if longList == shortList {
			return shortList
		}
		longList = longList.Next
		shortList = shortList.Next
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	nodeA := headA
	nodeB := headB
	for nodeA != nodeB {
		if nodeA != nil {
			nodeA = nodeA.Next
		} else {
			nodeA = headB
		}
		if nodeB !=nil{
			nodeB = nodeB.Next
		}else {
			nodeB = headA
		}
	}
	return nodeA
}

func getListHead(listNode *ListNode) int {
	length := 0
	for listNode != nil {
		length++
		listNode = listNode.Next
	}
	return length
}
func main() {

}
