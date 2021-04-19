package main

import (
	"fmt"
)

//
//func maxSlidingWindow(nums []int, k int) []int {
//	if len(nums) == 0 || k == 0 {
//		return []int{}
//	}
//	res := []int{}
//	q := list.New()
//
//	for i := 0; i < k; i++ {
//		for q.Len() != 0 && nums[i] >= nums[q.Back().Value.(int)] {
//			q.Remove(q.Back())
//		}
//		q.PushBack(i)
//	}
//
//	for i := k; i < len(nums); i++ {
//		res = append(res, nums[q.Front().Value.(int)])
//		for q.Len() != 0 && nums[i] >= nums[q.Back().Value.(int)] {
//			q.Remove(q.Back())
//		}
//		if q.Len() != 0 && q.Front().Value.(int) <= i-k {
//			q.Remove(q.Front())
//		}
//		q.PushBack(i)
//	}
//	res = append(res, nums[q.Front().Value.(int)])
//	return res
//}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	res := []int{}
	n := len(nums)
	q := []int{}

	for i := 0; i < k; i++ {
		for len(q) != 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := k; i < n; i++ {
		res = append(res, nums[q[0]])
		for len(q) != 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		if len(q) != 0 && q[0] <= i-k {
			q = q[1:]
		}
		q = append(q, i)
	}
	res = append(res, nums[q[0]])
	return res
}

func main() {
	fmt.Printf("%v", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))

}
