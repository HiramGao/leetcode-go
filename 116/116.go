package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	var prevNode *Node
	var tmp []*Node
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			continue
		}
		if prevNode != nil {
			prevNode.Next = node
		}
		prevNode = node
		tmp = append(tmp, node.Left)
		tmp = append(tmp, node.Right)
		if len(q) == 0 {
			fmt.Println("==========")
			q = tmp
			tmp = []*Node{}
			prevNode = nil
		}
	}
	return root
}
func connect1(root *Node) *Node {
	if root == nil {
		return nil
	}
	for leftNode := root; leftNode.Left != nil; leftNode = leftNode.Left {
		for node := leftNode; node != nil; node = node.Next {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}
func main() {

}
