package main

import (
	"fmt"
	"math"
)

func leastInterval(tasks []byte, n int) int {
	hashMap := map[byte]int{}
	for _, task := range tasks {
		hashMap[task]++
	}
	nextValid := make([]int, 0, len(hashMap))
	rest := make([]int, 0, len(hashMap))
	for _, c := range hashMap {
		nextValid = append(nextValid, 1)
		rest = append(rest, c)
	}
	result := 0
	for range tasks {
		result++
		minNextValid := math.MaxInt64

		for i, r := range rest {
			if r > 0 && nextValid[i] < minNextValid {
				minNextValid = nextValid[i]
			}
		}
		if minNextValid > result {
			result = minNextValid
		}
		best := -1
		for i, r := range rest {
			if r > 0 && nextValid[i] <= result && (best == -1 || r > rest[best]) {
				best = i
			}
		}
		nextValid[best] = result + n + 1
		rest[best]--
	}

	return result

}
func main() {
	tasks := []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'B', 'C', 'D', 'E', 'F', 'G'}
	n := 2
	fmt.Println(leastInterval(tasks, n))
}
