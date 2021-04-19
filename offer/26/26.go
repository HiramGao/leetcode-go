package main

import (
	. "../TreeNode"
)

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return hasTree(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
func hasTree(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return hasTree(A.Left, B.Left) && hasTree(A.Right, B.Right)
}

func main() {

}
