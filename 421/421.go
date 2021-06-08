package main

import "fmt"

const highBit = 30

type trie struct {
	l, r *trie
}

func (t *trie) add(num int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.l == nil {
				cur.l = &trie{}
			}
			cur = cur.l
		} else {
			if cur.r == nil {
				cur.r = &trie{}
			}
			cur = cur.r
		}
	}
}

func (t *trie) check(num int) (x int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			//应该向右边走
			if cur.r != nil {
				cur = cur.r
				x = x*2 + 1
			} else {
				cur = cur.l
				x = x * 2
			}
		} else {
			if cur.l != nil {
				cur = cur.l
				x = x*2 + 1
			} else {
				cur = cur.r
				x = x * 2
			}
		}
	}
	return
}

func findMaximumXOR(nums []int) int {
	res := 0
	root := &trie{}

	for i := 1; i < len(nums); i++ {
		root.add(nums[i-1])
		res = max(res, root.check(nums[i]))

	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {
	nums := []int{3, 10, 5, 25, 2, 8}
	fmt.Println(findMaximumXOR(nums))
}
