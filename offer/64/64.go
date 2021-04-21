package main

func sumNums(n int) int {

	var ans int
	var f func(r *int, n int) bool
	f = func(r *int, n int) bool {
		*r += n
		return n != 0 && f(r, n-1)
	}
	f(&ans, n)
	return ans
}
func main() {
	println(sumNums(10))
}
