package main

import (
	. "../TreeNode"
	"container/list"
	"fmt"
)

//func levelOrder(root *TreeNode) [][]int {
//	if root ==nil{
//		return [][]int{}
//	}
//	var (
//		res [][]int
//		q = []*TreeNode{root}
//		node *TreeNode
//		toBePushed = 1
//		nextLevel = 0
//
//		rowLineNode []int
//	)
//	for len(q) !=0{
//		node = q[0]
//		if node.Left!=nil{
//			q = append(q, node.Left)
//			nextLevel++
//		}
//		if node.Right!=nil{
//			q = append(q, node.Right)
//			nextLevel++
//		}
//		q = q[1:]
//
//		toBePushed--
//		rowLineNode = append(rowLineNode,node.Val)
//
//		if toBePushed ==0{
//			res = append(res,rowLineNode)
//			toBePushed = nextLevel
//			nextLevel=0
//			rowLineNode = []int{}
//		}
//
//	}
//	return res
//}
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		q   = list.New()
		res = [][]int{}
	)
	q.PushBack(root)
	for q.Len() != 0 {
		rowLine := []int{}
		for i := q.Len(); i > 0; i-- {
			node := q.Remove(q.Back()).(*TreeNode)
			rowLine = append(rowLine, node.Val)
			if node.Left!=nil{
				q.PushFront(node.Left)
			}
			if node.Right!=nil{
				q.PushFront(node.Right)
			}
		}
		res= append(res, rowLine)
	}
	return res
}
func main() {
	fmt.Printf("%v",
		levelOrder2(
			BuildTreeNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		),
	)
}
