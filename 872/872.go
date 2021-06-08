package main

import (
	. "../TreeNode"
	"fmt"
	"math"
)

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	var dfs func(node *TreeNode, source *[]int)

	dfs = func(node *TreeNode, source *[]int){
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			*source = append(*source, node.Val)
		}
		dfs(node.Left, source)
		dfs(node.Right, source)

	}
	var roo1Leaf []int
	var roo2Leaf []int

	dfs(root1, &roo1Leaf)
	dfs(root2, &roo2Leaf)

	if len(roo1Leaf)!= len(roo2Leaf){
		return false
	}

	for i, v := range roo1Leaf {
		if roo2Leaf[i] != v {
			return false
		}
	}
	return true
}

func main() {
	root1 := BuildTreeNode([]int{3, 5, 1, 99, 2, 9, 8, math.MaxInt64, math.MaxInt64, 7, 4})
	root2 := BuildTreeNode([]int{3, 5, 1, 6, 7, 4, 2, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, 9, 8})
	fmt.Println(leafSimilar(root1, root2))
}
