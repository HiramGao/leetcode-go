package main

import "fmt"

func combinationSum(candidates []int, target int) (ans [][]int) {

	n := len(candidates)
	var backtrack func(i, target int)
	sumArr := []int{}

	backtrack = func(index, target int) {
		if index == n {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), sumArr...))
			return
		}
		if target-candidates[index] >= 0 {
			sumArr = append(sumArr, candidates[index])
			target -= candidates[index]
			backtrack(index, target)
			sumArr = sumArr[:len(sumArr)-1]
			target += candidates[index]
		}

		backtrack(index+1, target)
	}

	backtrack(0, target)

	return
}
func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
