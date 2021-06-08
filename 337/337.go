package main

import (
	. "../TreeNode"
)

func rob(root *TreeNode) int {
	res := dfs(root)

	return max(res[0], res[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	selectRoot := node.Val + l[1] + r[1]
	notSelectRoot := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selectRoot, notSelectRoot}
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
func main() {

}
