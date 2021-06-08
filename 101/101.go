package main

import (
	. "../TreeNode"
)

func isSymmetric(root *TreeNode) bool {

	var dfs func(node,right *TreeNode) bool
	dfs= func(left,right *TreeNode) bool{
		if left == nil && right==nil{
			return true
		}
		if left == nil || right==nil{
			return false
		}
		return left.Val==right.Val && dfs(left.Left,right.Right) && dfs(left.Right,right.Left)

	}
	return dfs(root,root)
}
func main() {
	println(isSymmetric(BuildTreeNode([]int{1,2,2})))
}
