package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	temp := myPow(x, n/2)
	if n%2 == 0 {
		return temp * temp
	}
	return temp * temp * x
}

func main() {
	println(myPow(2.0, 10))
}
