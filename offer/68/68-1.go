package main

import (
	. "../../TreeNode"
)

func lowestCommonAncestor1(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	res := root
	for{
		if p.Val < root.Val && q.Val < root.Val{
			res = res.Left
		}else if p.Val > root.Val && q.Val > root.Val{
			res = res.Right
		}else {
			break
		}
	}
	return res
}
func main() {

}
