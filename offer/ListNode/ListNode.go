package ListNode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertListNode(head *ListNode, newNode *ListNode) {
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = newNode
}

func ListPrint(head *ListNode) []int {
	if head.Next == nil {
		return []int{}
	}
	temp := head
	var nums []int
	for {
		if temp == nil {
			break
		}
		fmt.Printf("Address:%p, Val:%d, Next:%p \n", temp, temp.Val, temp.Next)
		nums = append(nums, temp.Val)
		temp = temp.Next
	}
	return nums
}
