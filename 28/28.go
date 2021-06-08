package main

import "fmt"

func strStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if n == 0 {
		return 0
	}
	if m == n {
		if haystack == needle {
			return 0
		}
		return -1
	}
	for l := 0; l < m-n+1; l = l + 1 {
		if haystack[l:l+n] == needle {
			return l
		}
	}
	return -1
}

func strStr1(text, reg string) int {
	n, m := len(text), len(reg)
	if m == 0 {
		return 0
	}
	next := make([]int, m)

	for i, j := 1, 0; i < m; i++ {
		for j > 0 && reg[i] != reg[j] {
			j = next[j-1]
		}
		if reg[i] == reg[j] {
			j++
		}
		next[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && text[i] != reg[j] {
			j = next[j-1]
		}
		if text[i] == reg[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("abc", "c"))
	fmt.Println(strStr1("hello", "ll"))
	fmt.Println(strStr1("aaaaa", "bba"))
	fmt.Println(strStr1("abc", "c"))

}
