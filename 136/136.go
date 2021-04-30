package main

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res = res ^ num
	}
	return res
}
func main() {
	println(singleNumber([]int{4, 1, 2, 1, 2}))
}
