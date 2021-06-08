package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	var res bytes.Buffer
	if numerator*denominator < 0 {
		res.WriteByte('-')
	}
	numerator, denominator = abs(numerator), abs(denominator)
	res.WriteString(strconv.Itoa(numerator / denominator))

	remainder := numerator % denominator
	if remainder == 0 {
		return res.String()
	}
	res.WriteByte('.')

	mp := map[int]int{}
	var tmp string
	for remainder != 0 {

		if index, ok := mp[remainder]; ok {
			res.WriteString(tmp[:index])
			res.WriteByte('(')
			res.WriteString(tmp[index:])
			res.WriteByte(')')
			return res.String()
		}

		mp[remainder] = len(tmp)
		remainder *= 10
		tmp += strconv.Itoa(remainder / denominator)
		remainder %= denominator
	}
	return res.String() + tmp
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func main() {
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(2, 3))
	fmt.Println(fractionToDecimal(4, 333))
	fmt.Println(fractionToDecimal(-1, 5))
}
