package main

import "fmt"

func generateParenthesis(n int) []string {
	var (
		res       []string
		backtrack func(open int, close int)
		tmp       = ""
		max       = n * 2
	)

	backtrack = func(open int, close int) {

		if len(tmp) == max {
			res = append(res, tmp)
			return
		}
		if open < n {
			tmp += "("
			backtrack(open+1, close)
			if len(tmp) > 0 {
				tmp = tmp[:len(tmp)-1]
			}
		}
		if close < open {
			tmp += ")"
			backtrack(open, close+1)
			if len(tmp) > 0 {
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtrack(0, 0)
	return res
}
func main() {
	fmt.Print(generateParenthesis(3))
}
