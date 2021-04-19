package main

func majorityElement(nums []int) int {
	var (
		x     = 0
		votes = 0
	)
	for _, v := range nums {
		if votes == 0 {
			x = v
		}
		if x == v {
			votes += 1
		} else {
			votes += -1
		}
	}
	return x
}
func main() {
	println(majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}
