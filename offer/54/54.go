package main

import (
	. "../../TreeNode"
	"math"
)

func kthLargest(root *TreeNode, k int) int {
	var findTarget func(node *TreeNode) *TreeNode

	findTarget = func(node *TreeNode) *TreeNode {
		var target *TreeNode
		if node.Right != nil {
			target = findTarget(node.Right)
		}
		if target == nil {
			if k == 1 {
				return node
			}
			k--
		}
		if target == nil && node.Left != nil {
			target = findTarget(node.Left)
		}
		return target
	}
	target:=findTarget(root)
	return target.Val
}
func main() {
	println(kthLargest(
		BuildTreeNode([]int{5,3,6,2,4,math.MaxInt64,math.MaxInt64,1}),
		3))

}
