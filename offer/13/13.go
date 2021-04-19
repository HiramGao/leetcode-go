package main

func movingCount(m int, n int, k int) int {
	get := func(x int) int {
		res := 0
		for x != 0 {
			res += x % 10
			x = x / 10
		}
		return res
	}
	ans := 1
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	vis[0][0] = true

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 && j == 0) || get(i)+get(j) > k {
				continue
			}
			if i-1 >= 0 {
				vis[i][j] = vis[i-1][j] || vis[i][j]
			}
			if j-1 >= 0 {
				vis[i][j] = vis[i][j-1] || vis[i][j]
			}
			if vis[i][j] {
				ans += 1
			}
		}
	}
	return ans
}

func main() {
	println(movingCount(1, 2, 1))
}
