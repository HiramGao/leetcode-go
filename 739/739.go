package main

import "fmt"

func dailyTemperatures(T []int) []int {

	n := len(T)
	ans := make([]int, n)
	stack := make([]int, n)
	for i, t := range T {
		for len(stack) > 0 && t > T[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			ans[prevIndex] = i - prevIndex
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))

}
