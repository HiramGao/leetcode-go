package main

import (
	. "../offer/TreeNode"
	"math"
)

func rangeSumBST(root *TreeNode, low int, high int) int {
	var dfs func(node *TreeNode)
	sum := 0
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if node.Val >= low && node.Val <= high{
			sum+=node.Val
		}
		dfs(node.Right)
	}
	dfs(root)
	return sum
}
func rangeSumBST2(root *TreeNode, low int, high int) int {
	if root ==nil{
		return 0
	}
	if root.Val < low{
		return rangeSumBST(root.Right,low,high)
	}
	if root.Val > high{
		return rangeSumBST(root.Left,low,high)
	}
	return root.Val + rangeSumBST(root.Left,low,high)+rangeSumBST(root.Right,low,high)
}
func rangeSumBST3(root *TreeNode, low int, high int) int {
	sum:=0
	q:= []*TreeNode{root}
	for len(q)!=0{
		node:=q[0]
		q = q[1:]
		if node ==nil{
			continue
		}
		if node.Val>high{
			q = append(q, node.Left)
		}else if node.Val < low{
			q = append(q,node.Right)
		}else {
			sum+=node.Val
			q = append(q, node.Left)
			q = append(q,node.Right)
		}
	}
	return sum
}
func main() {
	println(rangeSumBST3(BuildTreeNode([]int{10, 5, 15, 3, 7, math.MaxInt64, 18}), 7, 15))
}
