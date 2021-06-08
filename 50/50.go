package main

import "fmt"

func myPow(x float64, n int) float64 {

	multi := func(n int) float64 {
		ans := 1.0
		y := x
		for n > 0 {
			if n%2 == 1 {
				ans *= y
			}
			y *= y
			n /= 2
		}
		return ans
	}

	if n > 0 {
		return multi(n)
	}

	return 1 / multi(-n)
}
func main() {
	fmt.Println(myPow(2, 10))
	fmt.Println(myPow(2.1, 3))
	fmt.Println(myPow(2, -2))
}
