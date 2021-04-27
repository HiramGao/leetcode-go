package main

import (
	"math"
)

func maxNumberOfBalloons(text string) int {
	hashMap := map[byte]int{}

	for _, b := range []byte(text) {
		hashMap[b]++

	}

	hashMap['l'] /= 2
	hashMap['o'] /= 2
	min := math.MaxInt64

	for _, b := range []byte("balloon") {

		if hashMap[b] < min {
			min = hashMap[b]
		}
	}
	return min
}
func main() {
	println(maxNumberOfBalloons("balloonbaoon"))
}
