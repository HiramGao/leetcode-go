package main

import "fmt"

func plusOne(digits []int) []int {

	n := len(digits)

	for p := n - 1; p >= 0; p-- {
		digits[p]++
		digits[p] = digits[p] % 10
		if digits[p] != 0 {
			return digits
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}
func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}
