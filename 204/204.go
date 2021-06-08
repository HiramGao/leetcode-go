package main

import "fmt"

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	isPrime := make([]bool, n)
	res := 0
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			res++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return res
}
func main() {
	fmt.Println(countPrimes(10))
	fmt.Println(countPrimes(0))
	fmt.Println(countPrimes(1))
}
