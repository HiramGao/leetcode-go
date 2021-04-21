package main

import (
	. "../TreeNode"
	"math"
)

func maxDepth(root *TreeNode) int {
	if root ==nil{
		return 0
	}
	left := maxDepth(root.Left)
	right :=maxDepth(root.Right)
	if left>right{
		return left+1
	}
	return right+1
}

func main() {
	println(maxDepth(
		BuildTreeNode([]int{3,9,20,math.MaxInt64,math.MaxInt64,15,7})))
}
