package main

func findRepeatNumber(nums []int) int {
	for i, _ := range nums {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return 0
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	result := findRepeatNumber(nums)
	println(result)
}
