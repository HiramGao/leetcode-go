package main

import (
	. "../TreeNode"
	"fmt"
	"math"
)

func levelOrder1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		res  []int
		q    = []*TreeNode{root}
		node *TreeNode
	)

	for len(q) != 0 {
		node = q[0]
		q = q[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	return res
}

func main() {
	fmt.Printf("%v",
		levelOrder1(
			BuildTreeNode([]int{1, 2, math.MaxInt64, 4,5,6,7,8,9,10}),
		),
	)
}
