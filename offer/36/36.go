package main

import (
	. "../../TreeNode"
)

func treeToDoublyList(root *TreeNode) *TreeNode {
	var pre, head *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != nil {
			pre.Right = node
		} else {
			head = node
		}
		node.Left = pre
		pre = node
		dfs(node.Right)
	}

	dfs(root)
	head.Left = pre
	pre.Right = head
	return head
}

func main() {

}
