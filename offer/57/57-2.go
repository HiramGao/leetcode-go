package main

import (
	"fmt"
)

func FindSequence(sum int) [][]int {
	ahead, behind := 1, 2
	mid := sum / 2
	var res [][]int
	currentSum := ahead + behind
	for ahead < mid {
		if currentSum == sum {
			res = append(res, getList(ahead, behind))
		}
		for currentSum > sum && ahead < mid {
			currentSum -= ahead
			ahead++
			if currentSum == sum {
				res = append(res, getList(ahead, behind))
			}
		}
		behind++
		currentSum += behind
	}
	return res
}

func getList(ahead int, behind int) []int {
	tmp := []int{}
	for i := ahead; i <= behind; i++ {
		tmp = append(tmp, i)
	}
	return tmp
}
func main() {
	fmt.Printf("%v", FindSequence(15))
}
