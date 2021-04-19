package main

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	popIndex := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) != 0 && stack[len(stack)-1] == popped[popIndex] {
			stack = stack[:len(stack)-1]
			popIndex++
		}
	}
	return len(stack) == 0
}

func main() {
	println(validateStackSequences(
		[]int{1, 2, 3, 4, 5},
		[]int{4, 5, 3, 2, 1},
	))
}
