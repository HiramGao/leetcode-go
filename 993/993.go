package main

import (
	. "../TreeNode"
	"fmt"
	"math"
)

func isCousins(root *TreeNode, x int, y int) bool {
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool

	var dfs func(node *TreeNode, parent *TreeNode, depth int)

	dfs = func(node *TreeNode, parent *TreeNode, depth int) {
		if node == nil {
			return
		}
		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		} else if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}
		if xFound && yFound {
			return
		}
		depth++
		dfs(node.Left, node, depth)
		if xFound && yFound {
			return
		}
		dfs(node.Right, node, depth)
		return
	}
	dfs(root, nil, 0)
	return xDepth == yDepth && xParent != yParent
}
func main() {
	root := BuildTreeNode([]int{1, 2, 3, math.MaxInt64, 4, math.MaxInt64, 5})
	fmt.Println(isCousins(root, 5, 4))
}
