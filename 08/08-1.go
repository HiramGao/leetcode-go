package main

func waysToStep(n int) int {
	if n < 3 {
		return n
	}
	if n == 3 {
		return 4
	}

	//f[0] = 1
	//f[1] = 2
	//f[2] = 4
	a, b, c, d := 1, 2, 4, 0
	for i := 3; i < n; i++ {
		//f[i] = ((f[i-1]+f[i-2])%1000000007 + f[i-3]) % 1000000007
		d = (a + b + c) % 1000000007
		a, b, c = b, c, d
	}
	return d
}
func main() {
	println(waysToStep(5))
}
