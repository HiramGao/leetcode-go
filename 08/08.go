package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	n := len(s)
	res := 0
	sign := 0
	isDigit := func(ch byte) bool {
		return '0' <= ch && ch <= '9'
	}
	var setNumber func(i int)
	setNumber = func(i int) {
		if i < n {
			if isDigit(s[i]) {
				res = res*10 + int(s[i]-'0')
				if sign == 1 {
					if -res < math.MinInt32 {
						res = -math.MinInt32
						return
					}
				} else {
					if res > math.MaxInt32 {
						res = math.MaxInt32
						return
					}
				}
				setNumber(i + 1)
			}
		}
	}
	setSign := func() {
		if n > 0 {
			switch s[0] {
			case '-':
				sign = 1
				setNumber(1)
			case '+':
				setNumber(1)
			default:
				setNumber(0)
			}
		}
	}
	setSign()
	if sign == 1 {
		return -res
	}
	return res
}
func main() {
	fmt.Println(myAtoi("42 1"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("+-47"))
}
