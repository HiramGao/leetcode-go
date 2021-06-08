package main

import (
	. "../../ListNode"
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var nums []int

	for nil != head {
		nums = append(nums, head.Val)
		head = head.Next
	}
	for i,j := 0, len(nums)-1;i<j;{
		nums[i],nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}

func main() {
	headArray := []int{1, 3, 2}
	head := &ListNode{}
	for _, v := range headArray {
		InsertListNode(head, &ListNode{Val: v, Next: nil})
	}
	result := reversePrint(head.Next)
	fmt.Printf("%v", result)
}
