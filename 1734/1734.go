package main

import "fmt"

func decode(encoded []int) []int {
	n := len(encoded)
	ans := make([]int, n+1)
	total := 0
	for i := 1; i <= n+1; i++ {
		total ^= i
	}
	odd := 0
	for i := 1; i < n; i += 2 {
		odd ^= encoded[i]
	}
	ans[0] = total ^ odd
	for i, v := range encoded {
		ans[i+1] = ans[i] ^ v
	}
	return ans
}
func main() {
	fmt.Println(decode([]int{7, 5, 6, 11, 14, 15, 11, 6}))
}
