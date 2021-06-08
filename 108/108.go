package main

import (
	. "../TreeNode"
	"fmt"
)

func sortedArrayToBST(nums []int) *TreeNode {
	n:= len(nums)
	if n == 0{
		return nil
	}
	mid:=n/2
	node:=&TreeNode{Val: nums[mid]}
	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])

	return node
}
func main() {
	nums:=[]int{1,3}
	fmt.Println(CovToArr(sortedArrayToBST(nums)))

}
