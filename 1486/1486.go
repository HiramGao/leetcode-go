package main

import "fmt"

func xorOperation(n int, start int) (ans int) {
	ans = start
	for i := 1; i < n; i++ {
		ans ^= (start + 2*i)
	}
	return ans
}
func main() {
	fmt.Print(xorOperation(5, 0))
}
