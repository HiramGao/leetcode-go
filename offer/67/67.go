package main

import (
	"math"
	"strings"
)

func strToInt(str string) int {

	var (
		bndry        = 214748364
		strTrimSpace = strings.TrimSpace(str)
		n            = len(strTrimSpace)
		c            = []byte(strTrimSpace)
		res, sign    = 0, 1
		i            = 1
	)
	if n == 0 {
		return 0
	}
	if c[0] == '-' {
		sign = -1
	} else if c[0] != '+' {
		i = 0
	}
	for ; i < n; i++ {
		if c[i] < '0' || c[i] > '9' {
			break
		}
		if res > bndry || (res == bndry && c[i] > '7') {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		res = res*10 + int(c[i]-'0')
	}
	return sign * res
}
func main() {
	println(strToInt(" -+123211312 asd"))
}
