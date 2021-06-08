package main

import (
	. "../TreeNode"
	"math"
)

func flatten(root *TreeNode) {

	var dfs func(node *TreeNode)
	var nodeList []*TreeNode

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		nodeList = append(nodeList, node)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	for i := 1; i < len(nodeList); i++ {
		prev,node := nodeList[i-1],nodeList[i]
		prev.Left = nil
		prev.Right = node
	}
}
func flatten2(root *TreeNode) {
	if root==nil{
		return
	}
	stack:=[]*TreeNode{root}
	prev:=&TreeNode{}
	for len(stack)!=0{
		curr:=stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev!=nil{
			prev.Left = nil
			prev.Right = curr
		}
		if curr.Right!=nil{
			stack = append(stack,curr.Right)
		}
		if curr.Left!=nil{
			stack = append(stack,curr.Left)
		}
		prev = curr
	}

}
func main() {
	ans := BuildTreeNode([]int{1, 2, 5, 3, 4, math.MaxInt64, 6})
	flatten2(ans)
	PrintTree(ans)
}
