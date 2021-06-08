package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	hashMap := map[*Node]*Node{}
	node := head
	var newHead, newNode *Node

	for node != nil {

		if newHead == nil {
			newHead = &Node{Val: node.Val}
			newNode = newHead
		} else {
			newNode.Next = &Node{Val: node.Val}
			newNode = newNode.Next
		}
		hashMap[node] = newNode

		node = node.Next
	}
	node = head
	newNode = newHead
	for node != nil {

		if rNode, ok := hashMap[node.Random]; ok {
			newNode.Random = rNode
		}

		node = node.Next
		newNode = newNode.Next
	}

	return newHead
}
func main() {

}
