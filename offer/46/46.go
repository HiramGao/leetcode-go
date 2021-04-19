package main

import (
	"strconv"
)

func translateNum(num int) int {
	str := strconv.Itoa(num)
	n := len(str)
	counts := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		count := 0
		if i == n-1 {
			count = 1
		} else {
			count = counts[i+1]

			digit1 := str[i] - '0'
			digit2 := str[i+1] - '0'
			cov := digit1*10 + digit2
			if cov >= 10 && cov <= 25 {
				if i != n-2 {
					count += counts[i+2]
				} else {
					count += 1
				}

			}

		}
		counts[i] = count
	}
	return counts[0]
}
func main() {
	println(translateNum(12258))
}
