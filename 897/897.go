package main

import (
	. "../offer/TreeNode"
)

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var dfs func(node *TreeNode)
	var arr []*TreeNode
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Left)
			arr = append(arr, node)
			dfs(node.Right)
		}
	}
	dfs(root)

	for i, node := range arr {
		if i < len(arr)-1 {
			node.Right = arr[i+1]
		} else {
			node.Right = nil
		}
		node.Left = nil
	}

	return arr[0]
}
func increasingBST2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var dfs func(node *TreeNode)
	dummy := &TreeNode{}
	resNode := dummy
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Left)

			resNode.Right = node
			node.Left = nil
			resNode = node

			dfs(node.Right)
		}
	}
	dfs(root)

	return dummy.Right
}

func main() {
	PrintTree(increasingBST2(BuildTreeNode([]int{5, 1, 7})))
}
