package main

import (
	"fmt"
	"sort"
)

type UnionFindSet struct {
	parent, rank []int
}

func NewUnionFindSet(n int) *UnionFindSet {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = i
	}
	return &UnionFindSet{parent: parent, rank: rank}
}
func (u *UnionFindSet) Find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

func (u *UnionFindSet) Union(x, y int) {
	x, y = u.Find(x), u.Find(y)
	if x == y {
		return
	}
	if u.rank[x] < u.rank[y] {
		x, y = y, x
	}
	u.parent[y] = x
	if u.rank[x] == u.rank[y] {
		u.rank[x] += 1
	}
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)

	set := NewUnionFindSet(n)

	for _, pair := range pairs {
		set.Union(pair[0], pair[1])
	}
	groups := map[int][]byte{}

	for i := range s {
		f := set.Find(i)
		groups[f] = append(groups[f], s[i])
	}

	for _, bytes := range groups {
		sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	}
	res := make([]byte, n)
	for i := range res {
		f := set.Find(i)

		res[i] = groups[f][0]
		groups[f] = groups[f][1:]

	}

	return string(res)
}
func main() {
	fmt.Println(smallestStringWithSwaps("vbjjxgdfnru", [][]int{{8, 6}, {3, 4}, {5, 2}, {5, 5}, {3, 5}, {7, 10}, {6, 0}, {10, 0}, {7, 1}, {4, 8}, {6, 2}}))
}

//"dbjjxgvfnru"
//"bdfgjjnuvrx"
