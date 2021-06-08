package main

import "fmt"

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    map[int]*DLinkedNode{},
		head:     &DLinkedNode{0, 0, nil, nil},
		tail:     &DLinkedNode{0, 0, nil, nil},
	}
	l.head.next = l.tail
	l.tail.prev = l.head

	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := &DLinkedNode{key, value, nil, nil}
		this.cache[key] = node
		this.addTohead(node)
		this.size++
		if this.size > this.capacity {
			node := this.removeTail()
			delete(this.cache, node.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addTohead(node)
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addTohead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func main() {
	ans := 0
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)    // 缓存是 {1=1}
	lRUCache.Put(2, 2)    // 缓存是 {1=1, 2=2}
	ans = lRUCache.Get(1) // 返回 1
	fmt.Println(ans)
	lRUCache.Put(3, 3)    // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	ans = lRUCache.Get(2) // 返回 -1 (未找到)
	fmt.Println(ans)
	lRUCache.Put(4, 4)    // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	ans = lRUCache.Get(1) // 返回 -1 (未找到)
	fmt.Println(ans)
	ans = lRUCache.Get(3) // 返回 3
	fmt.Println(ans)
	ans = lRUCache.Get(4) // 返回 4
	fmt.Println(ans)

}
