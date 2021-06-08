package main

import (
	. "../TreeNode"
	"fmt"
	"math"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left:=lowestCommonAncestor(root.Left,p,q)
	right:=lowestCommonAncestor(root.Right,p,q)
	if left != nil && right!=nil{
		return root
	}
	if left==nil{
		return right
	}
	return left

}
func main() {
	root := BuildTreeNode([]int{3, 5, 1, 6, 2, 0, 8, math.MaxInt64, math.MaxInt64, 7, 4})
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right))
}
