package main

func fib(n int) int {
	f0, f1 := 0, 1

	for i := 0; i < n; i++ {
		f0, f1 = f1, (f0+f1)%1000000007
	}
	return f0
}

func main() {
	println(fib(5))
}
