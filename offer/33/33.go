package main

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 0 {
		return true
	}
	n := len(postorder)
	root := postorder[n-1]
	//搜索左子树
	index := 0
	for index < n-1 {
		if postorder[index] > root {
			break
		}
		index++
	}
	for j := index; j < n-1; j++ {
		if postorder[j] < root {
			return false
		}
	}
	left := true
	if index > 0 {
		left = verifyPostorder(postorder[:index])
	}
	right := true
	if index < n-1 {
		right = verifyPostorder(postorder[index : n-1])
	}
	return right && left
}
func main() {
	println(verifyPostorder([]int{4, 8, 6, 12, 16, 14, 10}))
}
