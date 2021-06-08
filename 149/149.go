package main

import (
	"../utils"
	"fmt"
)

func maxPoints(points [][]int) int {
	n := len(points)
	if n < 3 {
		return n
	}
	res := 0
	for i,point1 := range points {
		hash := make(map[float64]int)
		for j, point2 := range points {
			if i != j {
				hash[lineSlope(point1,point2)]++
			}
		}
		for _, v := range hash {
			if v > res{
				res = v
			}
		}
	}

	return res+1
}

func lineSlope(a, b []int) float64 {
	return float64(a[1]-b[1]) / float64(a[0]-b[0])
}

func main() {
	ite := utils.NewITE("[[1,1],[2,2],[3,3]]")
	fmt.Println(maxPoints(ite.ToIntInt()))

	ite = utils.NewITE("[[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]")
	fmt.Println(maxPoints(ite.ToIntInt()))
}
