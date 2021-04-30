package main

import "fmt"

func hammingDistance(x int, y int) int {
	res := 0
	xor := x ^ y
	println(xor)

	for i := 0; i < 31; i++ {
		if xor&(1<<i) > 0 {
			res++
		}
	}

	return res
}
func main() {
	fmt.Println(hammingDistance(1, 4))
}
