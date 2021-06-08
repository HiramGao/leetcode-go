package main

import (
	. "../../TreeNode"
	"fmt"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootTreeNode := &TreeNode{Val: preorder[0]}
	count := 0
	for _, v := range inorder {
		if v == preorder[0] {
			break
		}
		count++
	}
	rootTreeNode.Left = buildTree(preorder[1:count+1], inorder[:count])
	rootTreeNode.Right = buildTree(preorder[count+1:], inorder[count+1:])
	return rootTreeNode
}

func main() {

	result := buildTree(
		[]int{3, 9, 20, 15, 7},
		[]int{9, 3, 15, 20, 7},
	)
	fmt.Printf("%v", result)
}
