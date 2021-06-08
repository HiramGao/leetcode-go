package main

import (
	. "../../TreeNode"
	"fmt"
	"math"
)

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int

	var findPath func(treeNode *TreeNode, subPath []int, target int)

	findPath = func(treeNode *TreeNode, subPath []int, target int) {
		if treeNode == nil {
			return
		}

		subPath = append(subPath, treeNode.Val)
		target -= treeNode.Val
		if target == 0 && treeNode.Left == nil && treeNode.Right == nil {
			res = append(res, append([]int{},subPath...))

			return
		}
		findPath(treeNode.Left, subPath, target)
		findPath(treeNode.Right, subPath, target)
	}

	findPath(root, []int{}, target)

	return res
}

func main() {
	fmt.Printf("%v",
		pathSum(
			BuildTreeNode([]int{5, 4, 8, 11, math.MaxInt64, 13, 4, 7, 2, math.MaxInt64, math.MaxInt64, 5, 1}),
			22,
		),
	)
}
