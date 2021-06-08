package main

import "fmt"

func isHappy(n int) bool {
	m := map[int]bool{}
	next := func(n int) int {
		sum := 0
		for n > 0 {
			d := n % 10
			sum += d * d
			n = n / 10
		}

		return sum
	}
	for ; n != 1 && !m[n]; n, m[n] = next(n), true {

	}
	return n == 1
}
func main() {
	fmt.Println(isHappy(19))
}
