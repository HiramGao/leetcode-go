package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {

	var res [][]int

	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})

	for _, person := range people {
		pIndex := person[1]
		res = append(res[:pIndex], append([][]int{person}, res[pIndex:]...)...)
	}

	return res
}
func main() {
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
}
