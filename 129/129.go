package main

import (
	. "../TreeNode"
	"fmt"
)

func sumNumbers(root *TreeNode) int {
	ans := 0
	if root == nil {
		return 0
	}
	sum:=0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left == nil && node.Right==nil {
			sum = sum * 10 + node.Val
			ans+=sum
			sum /=10
			return
		}
		sum = sum * 10 + node.Val
		if node.Left!=nil{dfs(node.Left)}
		if node.Right!=nil{dfs(node.Right)}
		sum /=10
	}
	dfs(root)

	return ans
}

func sumNumbers1(root *TreeNode) int {


	var dfs func(node *TreeNode,preSum int) int
	dfs = func(node *TreeNode,prevSum int) int{
		if node == nil{
			return 0
		}
		sum:=prevSum*10+node.Val
		if node.Left == nil && node.Right == nil{return sum}

		return dfs(node.Left,sum)+dfs(node.Right,sum)
	}

	return dfs(root,0)
}

func main() {
	root := BuildTreeNode([]int{4, 9, 0, 5, 1})
	fmt.Println(sumNumbers(root))
	fmt.Println(sumNumbers1(root))
}
