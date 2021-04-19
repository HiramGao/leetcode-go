package main

import (
	. "../TreeNode"
	"container/list"
	"fmt"
)

func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		q   = list.New()
		res = [][]int{}
		sw = true
	)
	q.PushBack(root)
	for q.Len() != 0 {
		rowLine := []int{}
		for i := q.Len(); i > 0; i-- {
			node := q.Remove(q.Back()).(*TreeNode)
			if sw {
				rowLine = append(rowLine, node.Val)
			}else {
				rowLine = append([]int{node.Val}, rowLine...)
			}

			if node.Left!=nil{
				q.PushFront(node.Left)
			}
			if node.Right!=nil{
				q.PushFront(node.Right)
			}
		}
		res= append(res, rowLine)
		sw = !sw
	}
	return res
}
func main() {
	fmt.Printf("%v",
		levelOrder3(
			BuildTreeNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		),
	)
}
