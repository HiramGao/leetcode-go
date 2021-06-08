package main

import (
	. "../TreeNode"
	"fmt"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := []*TreeNode{root}
	res := [][]int{}
	re := false
	var tmp []*TreeNode
	var tmpArr []int
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		tmpArr = append(tmpArr, node.Val)

		if node.Left!=nil{
			tmp = append(tmp, node.Left)
		}
		if node.Right!=nil{
			tmp = append(tmp, node.Right)
		}
		if len(q) == 0 {
			q = tmp
			if re{
				for l,r:=0,len(tmpArr)-1;l<r;l,r=l+1,r-1{
					tmpArr[l],tmpArr[r] = tmpArr[r],tmpArr[l]
				}
			}
			re=!re
			res = append(res, tmpArr)
			tmp = []*TreeNode{}
			tmpArr = []int{}
		}
	}
	return res
}
func main() {
	root := BuildTreeNode([]int{1, 2})
	fmt.Println(zigzagLevelOrder(root))
}
