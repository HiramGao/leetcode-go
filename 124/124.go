package main

import (
	. "../TreeNode"
	"fmt"
	"math"
)

func maxPathSum(root *TreeNode) int {
	ans:=math.MinInt32
	var dfs func(*TreeNode) int
	
	dfs = func(node *TreeNode) int{
		 if node == nil{
		 	return 0
		 }
		 left,right:=max(dfs(node.Left),0),max(dfs(node.Right),0)
		 ans = max(ans,left+right+node.Val)

		 return node.Val + max(left,right)
	}
	
	dfs(root)
	return ans
}
func max(i,j int)int  {
	if i > j{
		return i
	}
	return j
}
func main() {
	root:=BuildTreeNode([]int{-10,9,20,math.MaxInt64,math.MaxInt64,15,7})
	fmt.Println(maxPathSum(root))
}
