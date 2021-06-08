package main

import "fmt"

func calculate(s string) int {
	stack := []int{}
	preSign := '+'
	num := 0
	res := 0
	for i, b := range s {
		isNumber := '0' <= b && b <= '9'
		if isNumber {
			num = num*10 + int(b-'0')
		}
		if !isNumber && b != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			preSign = b
			num = 0
		}
	}

	for _, i := range stack {
		res += i
	}

	return res
}
func main() {
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate(" 3/2 "))
	fmt.Println(calculate(" 3+5 / 2 "))
}
