package main

func isValid(s string) bool {

	m := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	var stack []byte

	for _, v := range []byte(s) {
		if m[v] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != m[v] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}

	}
	return len(stack) == 0
}
func main() {
	println(isValid("()[]{}"))
}
