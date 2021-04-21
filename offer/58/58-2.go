package main

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
func main() {
	println(reverseLeftWords("abcdef", 2))
}
