package main

import "strings"

func reverseWords(s string) string {
	strList := strings.Split(s, " ")
	var res []string
	for i := len(strList) - 1; i >= 0; i-- {
		s := strings.TrimSpace(strList[i])
		if s != "" {
			res = append(res, s)
		}
	}
	return strings.Join(res, " ")
}
func main() {
	println(reverseWords("  hello world!  "))
}
