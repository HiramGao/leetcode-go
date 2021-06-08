package main

import (
	. "../TreeNode"
	"fmt"
	"math"
)

func longestUnivaluePath(root *TreeNode) int {

	ans := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := dfs(node.Left), dfs(node.Right)
		arrowLeft,arrowRight:=0,0
		if node.Left!=nil && node.Left.Val == node.Val{
			arrowLeft+=left+1
		}
		if node.Right != nil && node.Right.Val == node.Val{
			arrowRight+=right+1
		}
		ans = max(ans,arrowLeft+arrowRight)
		return max(arrowLeft,arrowRight)
	}
	dfs(root)
	return ans
}

func max(i int, j int) int {
	if i > j{
		return i
	}
	return j
}
func main() {
	root := BuildTreeNode([]int{5, 4, 5, 1, 1, math.MaxInt64, 5})
	fmt.Println(longestUnivaluePath(root))
}
