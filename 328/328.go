package main

import (
	. "../ListNode"
	"fmt"
)

func oddEvenList(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
	evenHead:=head.Next
	odd:=head
	even:=evenHead
	for even!=nil && even.Next!=nil{
		odd.Next = even.Next
		odd  = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
func main() {
	head := BuildList([]int{1, 2, 3, 4, 5, 6})
	ans := ListPrint(oddEvenList(head))
	fmt.Println(ans)
}
