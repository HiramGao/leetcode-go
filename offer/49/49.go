package main

func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}
	f := make([]int, n+1)
	f[1] = 1
	p2, p3, p4 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x4 := f[p2]*2, f[p3]*3, f[p4]*5
		f[i] = min(x2, min(x3, x4))
		if f[i] == x2 {
			p2++
		}
		if f[i] == x3 {
			p3++
		}
		if f[i] == x4 {
			p4++
		}
	}
	return f[n]
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}
func main() {
	println(nthUglyNumber(10))
}
