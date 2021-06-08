package main

import (
	"fmt"
	"math/big"
)

func uniquePaths(m int, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
func main() {
	fmt.Print(uniquePaths(3, 7))
}
