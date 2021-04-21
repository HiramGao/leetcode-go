package main

func searchNumberSameAsIndex(nums []int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		m := (start + end) >> 1

		if nums[m] == m {
			return m
		} else if nums[m] > m {
			end = m - 1
		} else {
			start = m + 1
		}
	}
	return 0
}

func main() {
	println(searchNumberSameAsIndex([]int{-3, -1, 1, 3, 5}))
}
