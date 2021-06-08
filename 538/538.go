package main

import (
	. "../TreeNode"
	"fmt"
	"math"
)

func convertBST(root *TreeNode) *TreeNode {
	if root ==nil{
		return root
	}
	nodeList := []*TreeNode{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		nodeList = append(nodeList, node)
		dfs(node.Left)
	}
	dfs(root)
	n:= len(nodeList)
	if n==1{
		return root
	}
	tmp:=nodeList[0].Val
	for i := 1; i <n; i++ {
		(*nodeList[i]).Val+=tmp
		tmp = (*nodeList[i]).Val
	}

	return root
}
func convertBST2(root *TreeNode) *TreeNode {
	sum:=0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum+=node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)

	return root
}
func main() {
	tree := BuildTreeNode([]int{4, 1, 6, 0, 2, 5, 7, math.MaxInt64, math.MaxInt64, math.MaxInt64, 3, math.MaxInt64, math.MaxInt64, math.MaxInt64, 8})
	ans := convertBST(tree)
	fmt.Println(CovToArr(ans))
	tree = BuildTreeNode([]int{4, 1, 6, 0, 2, 5, 7, math.MaxInt64, math.MaxInt64, math.MaxInt64, 3, math.MaxInt64, math.MaxInt64, math.MaxInt64, 8})
	ans = convertBST2(tree)
	fmt.Println(CovToArr(ans))
}
