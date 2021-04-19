package main

import "strings"

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func main() {
	s := "We are happy."
	result := replaceSpace(s)
	println(result)
}
