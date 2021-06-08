package main

import "fmt"

func titleToNumber(columnTitle string) int {
	res := 0
	for _, v := range columnTitle {
		res = res*26 + int(v-'A'+1)
	}
	return res
}
func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
}
