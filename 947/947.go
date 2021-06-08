package main

import "fmt"

func removeStones(stones [][]int) int {
	n := len(stones)
	set := NewUnionFindSet()
	for _, stone := range stones {
		set.Union(stone[0]+10001, stone[1])
	}
	return n - set.GetCount()
}

type UnionFindSet struct {
	parent map[int]int
	count  int
}

func (s *UnionFindSet) Union(x, y int) {
	x, y = s.Find(x), s.Find(y)
	if x == y {
		return
	}
	s.parent[x] = y
	s.count--

}

func (s *UnionFindSet) Find(x int) int {
	if _, ok := s.parent[x]; !ok {
		s.count++
		s.parent[x] = x
	}

	if s.parent[x] != x {
		s.parent[x] = s.Find(s.parent[x])
	}

	return s.parent[x]
}

func (s *UnionFindSet) GetCount() int {
	return s.count
}

func NewUnionFindSet() *UnionFindSet {
	return &UnionFindSet{map[int]int{}, 0}
}
func main() {
	input := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}
	fmt.Println(removeStones(input))
}
