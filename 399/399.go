package main

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	n := len(equations)
	set := NewUnionFindSet(n * 2)
	//字母和id的对应关系
	hashMap := map[string]int{}
	id := 0
	for i, equation := range equations {
		var1, var2 := equation[0], equation[1]
		if _, ok := hashMap[var1]; !ok {
			hashMap[var1] = id
			id++
		}
		if _, ok := hashMap[var2]; !ok {
			hashMap[var2] = id
			id++
		}
		set.Union(hashMap[var1], hashMap[var2], values[i])
	}
	res := make([]float64, len(queries))
	for i, query := range queries {
		q1, ok1 := hashMap[query[0]]
		q2, ok2 := hashMap[query[1]]
		if ok1 && ok2 {
			res[i] = set.IsConnected(q1, q2)
		} else {
			res[i] = -1
		}
	}

	return res
}

type UnionFindSet struct {
	parent []int
	weight []float64
}

func (s UnionFindSet) Union(x int, y int, v float64) {
	rootX, rootY := s.Find(x), s.Find(y)
	if rootX == rootY {
		return
	}
	s.parent[rootX] = rootY
	s.weight[rootX] = s.weight[y] * v / s.weight[x]
	fmt.Println("parent", s.parent)
	fmt.Println("weight", s.weight)
}

func (s UnionFindSet) Find(x int) int {
	if s.parent[x] != x {
		org := s.parent[x]
		s.parent[x] = s.Find(s.parent[x])
		s.weight[x] = s.weight[x] * s.weight[org]
	}
	return s.parent[x]
}

func (s UnionFindSet) IsConnected(x int, y int) float64 {
	rootX, rootY := s.Find(x), s.Find(y)
	if rootX == rootY {
		return s.weight[x] / s.weight[y]
	}
	return -1
}

func NewUnionFindSet(n int) *UnionFindSet {
	parent := make([]int, n)
	weight := make([]float64, n)
	for i := range parent {
		parent[i] = i
		weight[i] = 1.0
	}
	return &UnionFindSet{parent, weight}
}
func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println(calcEquation(equations, values, queries))
}
