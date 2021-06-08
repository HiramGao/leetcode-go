package main

import (
	"fmt"
)

func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n%3 == 1
}

func main() {
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(4 * 4 * 4))
	fmt.Println(isPowerOfFour(4 * 4 * 4 * 4))
	fmt.Println(isPowerOfFour(4 * 4 * 4 * 4 * 4))
	fmt.Println(isPowerOfFour(5))
}
