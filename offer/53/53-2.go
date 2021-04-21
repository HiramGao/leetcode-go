package main

func missingNumber(nums []int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		m := (start + end) >> 1

		if nums[m] != m {
			if m == 0 || nums[m-1] == m-1 {
				return m
			}
			end = m - 1
		} else {
			start = m + 1
		}
	}
	if start == len(nums) {
		return start
	}
	return 0
}
func missingNumber2(nums []int) int {
	count := len(nums)
	for i := 0; i < len(nums); i++ {
		count ^= i ^ nums[i]
		println(count)
	}
	return count
}
func main() {
	println(missingNumber2([]int{0, 1, 3}))
}
