package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}
	return ans
}
func mySqrt1(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := (r-l)/2 + l
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt1(4))
	fmt.Println(mySqrt1(8))
}
