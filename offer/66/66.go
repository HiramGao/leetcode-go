package main

import "fmt"

func constructArr(a []int) []int {
	n := len(a)
	if n == 0 {
		return []int{}
	}
	b := make([]int, n)
	b[0] = 1
	tmp := 1

	for i := 1; i < n; i++ {
		b[i] = b[i-1] * a[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		tmp *= a[i+1]
		b[i] *= tmp
	}

	return b
}
func main() {
	fmt.Printf("%v", constructArr([]int{1, 2, 3, 4, 5}))
}
