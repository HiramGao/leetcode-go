package main

import (
	. "../TreeNode"
	"math"
)

func maxDepth(root *TreeNode) int {
	if  root==nil{
		return 0
	}
	max := 0
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		depth++
		if node.Left!=nil{
			dfs(node.Left,depth)
		}
		if node.Right!=nil{
			dfs(node.Right,depth)
		}
		if node.Right==nil && node.Left==nil{
			if depth > max{
				max = depth
			}
		}
	}
	dfs(root, 0)
	return max
}
func main() {
	println(maxDepth(BuildTreeNode([]int{3, 9, 20, math.MaxInt64, math.MaxInt64, 15, 7})))
}
