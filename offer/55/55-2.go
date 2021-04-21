package main

import (
	. "../TreeNode"
	"math"
)

func isBalanced(root *TreeNode) bool {

	var isB func(node *TreeNode, beforeDep int) (is bool, dep int)

	isB = func(node *TreeNode, beforeDep int) (is bool, dep int) {
		if node == nil {
			return true, 0
		}
		left, leftDep := isB(node.Left, beforeDep)
		right, rightDep := isB(node.Right, beforeDep)

		if leftDep > rightDep {
			beforeDep = leftDep + 1
		} else {
			beforeDep = rightDep + 1
		}
		if left && right {
			diff := leftDep - rightDep
			if diff >= -1 && diff <= 1 {
				return true, beforeDep
			}
		}

		return false, beforeDep
	}
	res,_ := isB(root, 0)
	return res
}

func main() {
	println(isBalanced(
		BuildTreeNode([]int{1,2,2,3,3,math.MaxInt64,math.MaxInt64,4,4})))
}
