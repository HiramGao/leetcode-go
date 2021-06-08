package main

import (
	. "../TreeNode"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	index := 0

	for i, v := range inorder {
		if v== preorder[0]{
			index = i
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:index])+1],inorder[:index])
	root.Right = buildTree(preorder[len(inorder[:index])+1:],inorder[index+1:])
	return root
}
func main() {
	PrintTree(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
