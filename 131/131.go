package main

import "fmt"

func partition(s string) [][]string {
	n := len(s)
	var res [][]string
	if n == 0 {
		return res
	}
	//[i][j] i~j是否是回文串
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	var check func(l, r int) int
	// 0 未搜索 1 回文串 -1 不会回文串
	check = func(l, r int) int {
		if l >= r {
			return 1
		}
		if dp[l][r] != 0 {
			return dp[l][r]
		}
		dp[l][r] = -1
		if s[l] == s[r] {
			dp[l][r] = check(l+1, r-1)
		}
		return dp[l][r]

	}

	var dfs func(index int, path []string)
	dfs = func(index int, path []string) {
		if index == n {
			t := make([]string, len(path)) // 新建一个和temp等长的切片
			copy(t, path)
			res = append(res, t)
			return
		}
		for r := index; r < n; r++ {
			if check(index, r) > 0 {
				path = append(path, s[index:r+1])
				dfs(r+1, path)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0, []string{})
	return res
}
func main() {
	//fmt.Println(partition("aab"))
	//fmt.Println(partition("a"))
	//fmt.Println(partition("aaa"))
	fmt.Println(partition("cbbbcc"))
}
