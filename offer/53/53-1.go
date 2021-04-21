package main

func search1(nums []int, target int) int {
	var findTargetIndex func(target int) int
	findTargetIndex = func(target int) int {
		start, end := 0, len(nums)-1
		for start <= end {
			m := (start + end) / 2
			if nums[m] <= target {
				start = m + 1
			} else {
				end = m - 1
			}
		}
		return start
	}
	return findTargetIndex(target) - findTargetIndex(target-1)
}
func search2(nums []int, target int) int {
	start, end := 0, len(nums)-1
	res := 0
	for start <= end {
		m := (start + end) / 2

		if nums[m] == target {
			res++
			front := m - 1
			last := m + 1
			for front >= start && nums[front] == target {
				res++
				front--
			}
			for last <= end && nums[last] == target {
				res++
				last++
			}
			return res
		} else if nums[m] > target {
			end = m - 1
		} else {
			start = m + 1
		}
	}
	return res
}
func main() {
	println(search2([]int{5, 7, 7, 8, 8, 10}, 8))
}
