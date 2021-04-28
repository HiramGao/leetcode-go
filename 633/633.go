package main

import "math"

func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		result := left*left + right*right
		if result > c {
			right--
		} else if result < c {
			left++
		} else {
			return true
		}
	}
	return false
}
func main() {
	println(judgeSquareSum(4))
}
