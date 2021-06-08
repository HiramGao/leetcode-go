package main

import (
	"fmt"
	"strconv"
)

var hm = map[int]string{}

func countAndSay(n int) string {
	prev := "1"
	resolve := func(x int) string {

		if v, ok := hm[x]; ok {
			return v
		}

		ans := ""
		for i := 0; i < len(prev); {
			count := 1
			for i+1 < len(prev) && prev[i] == prev[i+1] {
				count++
				i++
			}
			ans += strconv.Itoa(count) + string(prev[i])
			i++
		}
		hm[x] = ans
		return ans
	}
	for i := 1; i < n; i++ {

		prev = resolve(i)
	}
	return prev
}
func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
	fmt.Println(countAndSay(6))
	fmt.Println(countAndSay(7))
	fmt.Println(countAndSay(8))
}
