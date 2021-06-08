package main

import (
	. "../TreeNode"
	"fmt"
)

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left+right > res {
			res = left + right
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}
	dfs(root)
	return res
}

func main() {
	fmt.Print(diameterOfBinaryTree(BuildTreeNode([]int{1, 2, 3, 4, 5})))
}
