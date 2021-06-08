package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var stack []string

	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if isNumber(cur) {
			digits := getDigits(s, &ptr)
			stack = append(stack, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stack = append(stack, string(cur))
			ptr++
		} else {
			ptr++
			var sub string
			//取出']'前 所有字符
			for stack[len(stack)-1] != "[" {
				sub = stack[len(stack)-1] + sub
				stack = stack[:len(stack)-1]
			}
			//去掉 '['
			stack = stack[:len(stack)-1]
			//取出重复次数
			repTime, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			t := strings.Repeat(sub, repTime)
			stack = append(stack, t)
		}
	}

	return arr2String(stack)
}

func arr2String(stack []string) string {
	str := ""
	for _, c := range stack {
		str += c
	}
	return str
}

func getDigits(s string, ptr *int) string {
	res := ""
	for ; isNumber(s[*ptr]); *ptr++ {
		res += string(s[*ptr])
	}
	return res
}

func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
func main() {
	//                          "accaccacc"
	fmt.Println(decodeString("abc3[cd]xyz"))
}
