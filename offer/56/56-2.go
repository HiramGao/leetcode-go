package main

func singleNumber1(nums []int) int {
	counts := [32]int{}
	for i := range nums {
		for j := 0; j < 32; j++ {
			counts[j] += nums[i] & 1
			nums[i] >>= 1
		}
	}
	for i := 0; i < 32; i++ {
		counts[i] %= 3
	}
	res := 0
	for i := 1; i <= 32; i++ {
		res <<= 1
		res |= counts[32-i]
	}
	return res
}

func singleNumber2(nums []int) int {
	one, two := 0, 0

	for _, v := range nums {
		one = one ^ v&(^two)
		two = two ^ v&(^one)
	}
	return one
}
func main() {
	println(singleNumber2([]int{3, 4, 3, 3}))
}
