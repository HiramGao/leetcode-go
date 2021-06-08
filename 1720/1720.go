package main

import "fmt"

func decode(encoded []int, first int) []int {
	n := len(encoded)
	res := make([]int, n+1)
	res[0] = first
	for i := 1; i < n+1; i++ {
		res[i] = res[i-1] ^ encoded[i-1]
	}
	return res
}
func main() {
	fmt.Printf("%v", decode([]int{1, 2, 3}, 1))
}
