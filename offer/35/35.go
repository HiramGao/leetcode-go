package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList1(head *Node) *Node {
	if head == nil {
		return head
	}
	node := head
	hashMap := make(map[*Node]*Node)
	for node != nil {
		cloneNode := &Node{Val: node.Val}
		hashMap[node] = cloneNode
		node = node.Next
	}
	node = head
	for node != nil {
		hashMap[node].Next = hashMap[node.Next]
		hashMap[node].Random = hashMap[node.Random]
		node = node.Next
	}
	return hashMap[head]
}

func main() {

}
