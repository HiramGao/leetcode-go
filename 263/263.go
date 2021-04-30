package main

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	f := []int{2, 3, 5}
	for _, v := range f {
		for n%v == 0 {
			n /= v
		}
	}
	return n == 1
}
func main() {
	println(isUgly(14))
}
