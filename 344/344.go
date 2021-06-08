package main

import "fmt"

func reverseString(s []byte) {

	for l, r := 0, len(s)-1; l < r; {

		s[l], s[r] = s[r], s[l]

		l++
		r--
	}
}
func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o', '0'}
	reverseString(s)
	fmt.Println(string(s))
}
