package main

import "fmt"

func subsets1(nums []int) [][]int {
	var res = [][]int{}
	n := len(nums)

	for i := 0; i < 1<<n; i++ {
		tmp := []int{}
		for j, num := range nums {
			if i>>j&1 > 0 {
				tmp = append(tmp, num)
			}
		}
		res = append(res, append([]int(nil), tmp...))
	}
	return res
}

func subsets(nums []int) [][]int {
	var res = [][]int{}
	set := []int{}
	n := len(nums)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			res = append(res, append([]int(nil), set...))
			return
		}
		set = append(set, nums[i])
		dfs(i + 1)
		set = set[:len(set)-1]
		dfs(i + 1)
	}
	dfs(0)
	return res
}
func main() {
	fmt.Print(subsets([]int{1, 2, 3}))
}
