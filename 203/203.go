package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

func removeElements1(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}

	for node := dummy; node.Next != nil; {
		if node.Next.Val == val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}

	return dummy.Next
}

func main() {

}
