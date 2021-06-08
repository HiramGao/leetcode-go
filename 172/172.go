package main

import "fmt"

func trailingZeroes(n int) int {
	zeroCount := 0
	currentMultiple := 5
	for n >= currentMultiple {
		zeroCount += n / currentMultiple
		currentMultiple *= 5
	}
	return zeroCount

}
func main() {
	fmt.Println(trailingZeroes(3))
	fmt.Println(trailingZeroes(5))
}
