package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	negative := (dividend ^ divisor) < 0
	dividend, divisor = abs(dividend), abs(divisor)
	if divisor > dividend {
		return 0
	}
	result := 0
	for i := 31; i >= 0; i-- {
		if (dividend >> i) >= divisor {
			result += 1 << i
			dividend -= divisor << i
		}
	}
	if negative {
		return -result
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	fmt.Println(divide(100, 3))
	fmt.Println(divide(10, 3))
	fmt.Println(divide(10, -3))
	fmt.Println(divide(7, 3))
	fmt.Println(divide(-7, -3))
}
