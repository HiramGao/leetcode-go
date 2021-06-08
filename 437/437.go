package main

import (
	. "../TreeNode"
	"fmt"
)

func pathSum(root *TreeNode, targetSum int) int {
	var dfs func(node *TreeNode)

	sum := 0
	ans := 0

	hashMap := make(map[int]int)
	hashMap[0]=1

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		sum += node.Val
		ans += hashMap[sum-targetSum]
		hashMap[sum]++

		dfs(node.Left)
		dfs(node.Right)
		hashMap[sum]--
		sum -= node.Val
	}
	dfs(root)
	return ans
}



func main() {
	root := BuildTreeNode([]int{1})
	ans := pathSum(root, 0)
	fmt.Print(ans)
}
