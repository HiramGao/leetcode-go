package main

import (
	"fmt"
	"math"
)

func printNumbers(n int) []int {
	max := int(math.Pow10(n)) - 1
	var res []int
	for i := 1; i <= max; i++ {
		res = append(res, i)
	}

	return res
}

func main() {
	fmt.Printf("%v", printNumbers(2))
}
