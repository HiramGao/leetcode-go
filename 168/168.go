package main

import "fmt"

func convertToTitle(columnNumber int) string {
	ans := []byte{}
	for columnNumber > 0 {
		a0 := (columnNumber-1)%26 + 1
		ans = append(ans, 'A'+byte(a0-1))
		columnNumber = (columnNumber - a0) / 26
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return string(ans)
}

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(2147483647))
}
