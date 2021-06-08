package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	n := len(this.nums)
	tmp := make([]int, n)
	copy(tmp, this.nums)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return tmp
}
func main() {
	solution := Constructor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(solution.Shuffle())
	fmt.Println(solution.Reset())
	fmt.Println(solution.Shuffle())
}
