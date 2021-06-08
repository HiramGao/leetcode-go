package main

import (
	. "../../TreeNode"
	"math"
)

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	res := root
	if res == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}

	return left
}
func main() {
	root := BuildTreeNode([]int{3, 5, 1, 6, 2, 0, 8, math.MaxInt32, math.MaxInt32, 7, 4})
	println(lowestCommonAncestor(root, &TreeNode{Val: 5}, &TreeNode{Val: 1}).Val)
}
