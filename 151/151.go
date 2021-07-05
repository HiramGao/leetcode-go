package main

import (
	"bytes"
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	left, right := 0, len(s)-1

	for left < right && s[left] == ' ' {
		left++
	}
	for left < right && s[right] == ' ' {
		right--
	}
	ans := []string{}

	var word bytes.Buffer

	for left <= right {
		ch := s[left]
		if word.Len() != 0 && ch == ' ' {
			ans = append(ans, word.String())
			word.Reset()
		} else if ch != ' ' {
			word.WriteByte(ch)
		}
		left++
	}
	ans = append(ans, word.String())
	n := len(ans)
	for i := 0; i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}

	return strings.Join(ans, " ")
}
func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))

	s = "  hello world  "
	fmt.Println(reverseWords(s))

	s = "a good   example"
	fmt.Println(reverseWords(s))

	s = "  Bob    Loves  Alice   "
	fmt.Println(reverseWords(s))

	s = "Alice does not even like bob"
	fmt.Println(reverseWords(s))
}
