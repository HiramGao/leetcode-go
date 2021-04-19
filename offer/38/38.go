package main

import (
	"fmt"
)

func permutation(s string) []string {

	res := []string{}
	c := []byte(s)
	n := len(s)

	var order func(index int)
	order = func(index int) {
		if index == n-1 {
			res = append(res, string(c))
			//return
		}
		dict := map[byte]bool{}

		for i := index; i < n; i++ {
			if !dict[c[i]] {
				c[index], c[i] = c[i], c[index]
				dict[c[index]] = true
				order(index + 1)
				c[index], c[i] = c[i], c[index]

			}
		}
	}

	order(0)

	return res
}

func main() {
	fmt.Printf("%v", permutation("ab"))
}
