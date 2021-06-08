package main

import (
	. "../TreeNode"
	"fmt"
	"math"
)

func inorderTraversal(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}

	dfs(root)

	return
}
func main() {
	ans := inorderTraversal(BuildTreeNode([]int{1, math.MaxInt64, 2, 3}))
	fmt.Println(ans)
}
