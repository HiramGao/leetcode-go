package main

import (
	. "../TreeNode"
	"fmt"
	"math"
)

func isValidBST1(root *TreeNode) bool {
	var stack []*TreeNode
	inorder := math.MinInt64

	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, lower, upper int) bool

	dfs = func(node *TreeNode, lower, upper int) bool {
		if nil == node {
			return true
		}
		if node.Val <= lower || node.Val >= upper {
			return false
		}
		return dfs(node.Left, lower, node.Val) && dfs(node.Right, node.Val, upper)
	}

	return dfs(root, math.MinInt64, math.MaxInt64)
}
func main() {
	ans := isValidBST(BuildTreeNode([]int{5, 4, 6, math.MaxInt64, math.MaxInt64, 3, 7}))
	fmt.Println(ans)
}
