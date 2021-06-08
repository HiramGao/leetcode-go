package main

import "fmt"

func hammingWeight(num uint32) int {

	res := 0
	for ; num > 0; num &= num - 1 {
		res++
	}
	return res
}
func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}
