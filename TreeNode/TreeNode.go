package TreeNode

import (
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTreeNode(source []int) *TreeNode {

	if len(source) == 0 {
		return nil
	}

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

func PrintTree(root *TreeNode) {
	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			dfs(node.Left)
		}
		println(node.Val)
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)
}

func CovToArr(root *TreeNode) (res []string) {
	if root == nil {
		return []string{}
	}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			res = append(res, "nil")
			continue
		}

		res = append(res, strconv.Itoa(node.Val))
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	return res
}
