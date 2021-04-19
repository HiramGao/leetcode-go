package main

import "fmt"

//func exchange(nums []int) []int {
//	begin, end := 0, len(nums)-1
//	for begin < end {
//		//奇数
//		for nums[begin]&1 == 1 && begin < end {
//			begin++
//		}
//		//偶数
//		for nums[end]&1 == 0 && begin < end {
//			end--
//		}
//		if begin < end {
//			nums[begin], nums[end] = nums[end], nums[begin]
//			begin++
//			end--
//		}
//	}
//	return nums
//}

func exchange(nums []int) []int {
	begin, end := 0, len(nums)-1
	for begin < end {
		isBeginOdd := isOdd(nums[begin])
		isEndOdd := isOdd(nums[end])
		if !isBeginOdd && !isEndOdd { //i,j偶
			//nums[i], nums[j] = nums[j], nums[i]
			end = end - 1
		} else if !isBeginOdd && isEndOdd { //i 偶， j奇
			nums[begin], nums[end] = nums[end], nums[begin]
			begin, end = begin+1, end-1
		} else if isBeginOdd { //i 奇，j偶
			begin = begin + 1
		}
	}
	return nums
}

func isOdd(num int) bool {
	if num&1 == 0 {
		return false
	}
	return true
}

func main() {
	fmt.Printf("%v", exchange([]int{1, 2, 3, 5}))
}
