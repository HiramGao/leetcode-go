package TreeNode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTreeNode(source []int) *TreeNode {

	var (
		q     = []*TreeNode{{Val: source[0]}}
		index = 1
	)
	root := &q[0]
	for len(q) != 0 && index < len(source) {
		node := &q[0]
		if index < len(source) && source[index] != math.MaxInt64 {
			leftNode := &TreeNode{Val: source[index]}
			(*node).Left = leftNode
			q = append(q, leftNode)
		}
		if index+1 < len(source) && source[index+1] != math.MaxInt64 {
			rightNode := &TreeNode{Val: source[index+1]}
			(*node).Right = rightNode
			q = append(q, rightNode)
		}
		index += 2
		q = q[1:]
	}

	return *root
}
