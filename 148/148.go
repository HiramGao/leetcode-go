package main

import (
	. "../ListNode"
	"fmt"
)

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sortList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummy, dummy.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength &&cur!=nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = merge(head1, head2)
			for prev.Next!=nil{
				prev = prev.Next
			}

			cur = next
		}
	}
	return dummy.Next
}

func sort(head *ListNode, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	return merge(sort(head, slow), sort(slow, tail))
}

func merge(node *ListNode, node2 *ListNode) *ListNode {
	dummy := &ListNode{}
	newhead, head1, head2 := dummy, node, node2

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			newhead.Next = head1
			head1 = head1.Next
		} else {
			newhead.Next = head2
			head2 = head2.Next
		}
		newhead = newhead.Next
	}
	if head1 != nil {
		newhead.Next = head1
	} else if head2 != nil {
		newhead.Next = head2
	}
	return dummy.Next
}
func main() {
	ans := ListPrint(sortList2(BuildList([]int{4, 2, 1, 3})))
	fmt.Println(ans)
}
