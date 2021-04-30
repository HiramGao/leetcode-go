package main

func singleNumber1(nums []int) int {
	hashMap := map[int]int{}
	for _, num := range nums {
		hashMap[num]++
	}
	for key, value := range hashMap {
		if value == 1 {
			return key
		}
	}
	return 0
}
func singleNumber2(nums []int) int {
	ans := int32(0)

	for i := 0; i < 32; i++ {
		sum := int32(0)
		for _, num := range nums {
			sum += int32(num) >> i & 1
		}
		if sum%3 > 0 {
			ans |= 1 << i
		}
	}

	return int(ans)
}
func main() {
	println(singleNumber2([]int{0, 1, 0, 1, 0, 1, 99}))
}
