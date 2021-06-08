package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		if isNumber(token) {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		} else {
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)

			}
		}
	}
	return stack[0]
}

func isNumber(token string) bool {
	return !(token == "+" || token == "-" || token == "*" || token == "/")
}
func main() {
}
