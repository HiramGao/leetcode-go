package main

import (
	"fmt"
)

func letterCasePermutation(S string) []string {

	var res []string
	for _, v := range []byte(S) {
		n := len(res)
		if (v >= 97 && v <= 122) || (v >= 65 && v <= 90) {
			if n == 0 {
				res = append(res, []string{string(v), string(v ^ 32)}...)
				continue
			}
			for i := 0; i < n; i++ {
				res = append(res, res[i]+string(v^32))
				res[i] = res[i] + string(v)
			}
		} else {
			if n == 0 {
				res = append(res, string(v))
				continue
			}
			for i := 0; i < n; i++ {
				res[i] = res[i] + string(v)
			}
		}
	}
	return res
}
func main() {
	fmt.Printf("%v", letterCasePermutation("32z"))
}
