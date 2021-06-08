package main

import "fmt"

func reverseParentheses(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	index := make([]int, n)
	for i := range index {
		index[i] = i
	}
	stack := []int{}

	reverse := func(l, r int) {
		for ; l < r; l, r = l+1, r-1 {
			index[l], index[r] = index[r], index[l]
		}
	}
	for i, b := range s {
		if b == '(' {
			index[i] = -1
			stack = append(stack, i)
		} else if b == ')' {
			index[i] = -1
			leftIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			reverse(leftIndex+1, i-1)
		}
	}

	res := []byte{}
	for _, i := range index {
		if i > -1 {
			res = append(res, s[i])
		}
	}

	return string(res)
}
func main() {
	fmt.Println(reverseParentheses("(abcd)"))
	fmt.Println(reverseParentheses("(u(love)i)"))
	fmt.Println(reverseParentheses("(ed(et(oc))el)"))
	fmt.Println(reverseParentheses("a(bcdefghijkl(mno)p)q"))
}
