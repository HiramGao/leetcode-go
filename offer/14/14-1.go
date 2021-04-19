package main

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	res := 1
	for n > 4 {
		res = res * 3 % 1000000007
		n -= 3
	}
	return n * res % 1000000007
}

func main() {
	println(cuttingRope(120))
}
