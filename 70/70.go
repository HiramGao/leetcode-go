package main

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	f := make([]int, n+1)
	f[1] = 1
	f[2] = 2
	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
func main() {
	println(climbStairs(2))
}
