package main

import (
	. "../TreeNode"
	"fmt"
	"math"
)

func hasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(node *TreeNode) bool

	sum := 0

	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		sum += node.Val
		if node.Left == nil && node.Right == nil && sum == targetSum {
			return true
		}
		if node.Left != nil {
			if dfs(node.Left) {
				return true
			}
		}
		if node.Right != nil {
			if dfs(node.Right) {
				return true
			}
		}
		sum -= node.Val
		return false
	}

	return dfs(root)
}
func main() {
	root := BuildTreeNode([]int{5, 4, 8, 11, math.MaxInt64, 13, 4, 7, 2, math.MaxInt64, math.MaxInt64, math.MaxInt64, 1})
	ans := hasPathSum(root, 22)
	fmt.Print(ans)
}
