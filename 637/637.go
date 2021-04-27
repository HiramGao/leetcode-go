package main

import (
	. "../offer/TreeNode"
	"fmt"
	"math"
)

func averageOfLevels(root *TreeNode) []float64 {
	if root ==nil{
		return []float64{}
	}
	q := []*TreeNode{root}
	res := []float64{}
	for len(q) != 0 {
		n := len(q)
		sum := 0
		for i := 0; i < n; i++ {
			node := q[i]
			sum += node.Val
			if node.Left!=nil{
				q = append(q, node.Left)
			}
			if node.Right!=nil{
				q = append(q, node.Right)
			}
		}
		q= q[n:]
		res = append(res, float64(sum)/float64(n))
	}
	return res
}
func main() {
	fmt.Printf("%v", averageOfLevels(BuildTreeNode([]int{3, 9, 20, math.MaxInt64, math.MaxInt64, 15, 7})))
}
