package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		result  []int
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
		visited[u] = -1
		result = append(result, u)
	}

	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}
	for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
		result[l], result[r] = result[r], result[l]
	}

	return result
}

func findOrder1(numCourses int, prerequisites [][]int) []int {
	var (
		result []int
		edges  = make([][]int, numCourses)
		indeg  = make([]int, numCourses)
		q      = []int{}
	)
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
		indeg[v[0]]++
	}

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
	if len(result) != numCourses {
		return []int{}
	}

	return result
}
func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	fmt.Println(findOrder1(2, [][]int{{1, 0}}))
	fmt.Println(findOrder1(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
