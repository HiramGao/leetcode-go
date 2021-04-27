package main

func numOfWays(n int) int {

	var all [][]int
	for j := 0; j < 3; j++ {
		for k := 0; k < 3; k++ {
			for l := 0; l < 3; l++ {
				if j == k || k == l {
					continue
				}
				all = append(all, []int{j, k, l})
			}
		}
	}
	allLen := len(all)
	dp := [5005][12]int{}
	for i := 0; i < allLen; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < allLen; j++ {
			for k := 0; k < allLen; k++ {
				if !sameColor(all[j], all[k]) {
					dp[i][j] = (dp[i][j] + dp[i-1][k]) % 1000000007
				}
			}
		}
	}
	res := 0
	for i := 0; i < allLen; i++ {
		res = (res + dp[n-1][i]) % 1000000007
	}
	return res
}

func sameColor(a []int, b []int) bool {
	for i := 0; i < 3; i++ {
		if a[i] == b[i] {
			return true
		}
	}
	return false
}
func numOfWays2(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 12
	}
	var aba int64 = 6
	var abc int64 = 6
	mod := int64(1000000007)
	for i := 1; i < n; i++ {
		aba, abc = (3*aba+2*abc)%mod, (2*aba+2*abc)%mod
	}
	return int((aba + abc) % mod)
}
func main() {
	println(numOfWays2(3))
}
