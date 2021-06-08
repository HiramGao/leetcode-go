package main

import (
	. "../TreeNode"
	"fmt"
	"math"
)

func levelOrder(root *TreeNode) (ans [][]int) {
	if root==nil{
		return
	}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		var otherQueue []*TreeNode
		var tmp []int
		for len(queue)!=0{
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				otherQueue = append(otherQueue, node.Left)
			}
			if node.Right != nil {
				otherQueue = append(otherQueue, node.Right)
			}
		}
		ans = append(ans, tmp)
		queue = otherQueue
	}

	return
}
func main() {
	fmt.Println(levelOrder(BuildTreeNode([]int{3,9,20,math.MaxInt64,math.MaxInt64,15,7})))
}
