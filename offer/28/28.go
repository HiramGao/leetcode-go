package main

import (
	. "../../TreeNode"
)

func isSymmetric(root *TreeNode) bool {

	return isSameTree(root, root)
}

func isSameTree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil ||left.Val != right.Val{
		return false
	}
	return isSameTree(left.Left, right.Right) && isSameTree(left.Right, right.Left)
}

func main() {

}
