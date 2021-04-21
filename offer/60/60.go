package main

import (
	"fmt"
	"math"
)

func dicesProbability(n int) []float64 {
	pProbabilities := [67]int{}
	for i := 1; i <= 6; i++ {
		pProbabilities[i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 6 * i; j >= i; j-- {
			pProbabilities[j] = 0
			for cur := 1; cur <= 6; cur++ {
				if j-cur < i-1 {
					break
				}
				pProbabilities[j] += pProbabilities[j-cur]
			}
		}
	}
	all := math.Pow(6, float64(n))
	var res []float64

	for i := n; i <= 6*n; i++ {
		res = append(res, float64(pProbabilities[i])/all)
	}
	return res
}
func main() {
	fmt.Printf("%v", dicesProbability(2))
}
