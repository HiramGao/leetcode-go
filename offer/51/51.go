package main

import "sort"

func reversePairs(nums []int) (result int) {
	var discretNum []int
	var sumArr []int
	sumArr = make([]int, len(nums)+5)
	discretization := func() {
		set := map[int]bool{}
		for _, num := range nums {
			set[num] = true
		}
		discretNum = make([]int, 0, len(set))
		for num := range set {
			discretNum = append(discretNum, num)
		}
		sort.Ints(discretNum)
	}

	getId := func(x int) int {
		return sort.SearchInts(discretNum, x) + 1
	}
	query := func(id int) int {
		ret := 0
		for id > 0 {
			ret += sumArr[id]
			id -= lowBit(id)
		}
		return ret
	}
	update := func(id int) {
		for id < len(sumArr) {
			sumArr[id]++
			id += lowBit(id)
		}
	}
	discretization()
	for i := len(nums) - 1; i >= 0; i-- {
		id := getId(nums[i])
		result += query(id - 1)
		update(id)
	}

	return
}
func lowBit(x int) int {
	return x & (-x)
}
func main() {

}
