package main

import (
	. "../TreeNode"
	"fmt"
	"math"
)

func kthSmallest(root *TreeNode, k int) int {

	nums := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		nums = append(nums, node.Val)
		dfs(node.Right)
	}
   dfs(root)
	return nums[k-1]
}
func main() {
	root := BuildTreeNode([]int{3, 1, 4, math.MaxInt64, 2})
	fmt.Println(kthSmallest(root, 3))
}
