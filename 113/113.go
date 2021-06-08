package main

import (
	. "../TreeNode"
	"fmt"
	"math"
)

func pathSum(root *TreeNode, targetSum int) [][]int {
	var dfs func(node *TreeNode)

	sum := 0
	ans := [][]int{}
	tmp := []int{}
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		sum += node.Val
		tmp = append(tmp, node.Val)
		if node.Left == nil && node.Right == nil && sum == targetSum {
			ans = append(ans, append([]int(nil), tmp...))
		}
		dfs(node.Left)
		dfs(node.Right)
		sum -= node.Val
		tmp = tmp[:len(tmp)-1]
	}
	dfs(root)
	return ans
}
func main() {
	root := BuildTreeNode([]int{5,4,8,11,math.MaxInt64,13,4,7,2,math.MaxInt64,math.MaxInt64,5,1})
	ans := pathSum(root, 22)
	fmt.Print(ans)
}
