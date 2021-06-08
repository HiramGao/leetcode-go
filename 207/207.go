package main

import "fmt"

func canFinish2(numCourses int, prerequisites [][]int) bool {

	var (
		edges  = make([][]int, numCourses)
		indeg  = make([]int, numCourses)
		result = []int{}
	)

	for _, prerequisite := range prerequisites {
		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
		indeg[prerequisite[0]]++
	}
	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}

func canFinish1(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		valid   = true
		dfs     func(u int)
	)
	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
	}
	for _, prerequisite := range prerequisites {
		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}
func main() {
	prerequisites := [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}
	fmt.Println(canFinish1(5, prerequisites))
	fmt.Println(canFinish2(5, prerequisites))
}
