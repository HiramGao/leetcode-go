package main

import (
	. "../TreeNode"
)

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left, right := invertTree(root.Left), invertTree(root.Right)
	root.Right = left
	root.Left = right

	return root
}

func main() {

}
