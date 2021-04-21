package main

func add(a int, b int) int {
	var sum, carry int
	for {
		sum = a ^ b
		carry = (a & b) << 1
		a = sum
		b = carry
		if carry == 0 {
			break
		}
	}
	return sum
}
func main() {
	println(add(1, 1))
}
