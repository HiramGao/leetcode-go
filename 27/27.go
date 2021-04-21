package main

import "fmt"

func removeElement1(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}
func removeElement2(nums []int, val int) int {
	left, right := 0, len(nums)

	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}

	}

	return left
}
func main() {
	nums := []int{3, 3, 3, 3}
	val := 3
	len := removeElement2(nums, val)
	println(len)
	for i := 0; i < len; i++ {
		fmt.Println(nums[i])
	}
}
