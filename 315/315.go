package main

import (
	"../utils"
	"fmt"
	"sort"
)

func countSmaller(nums []int) (result []int) {
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
		result = append(result, query(id-1))
		update(id)
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}


	return
}
func lowBit(x int) int {
	return x & (-x)
}

func main() {
	ite := utils.NewITE(`[5,2,6,1]`)
	fmt.Println(countSmaller(ite.ToInt()))
	fmt.Println(lowBit(1))
	fmt.Println(lowBit(2))
	fmt.Println(lowBit(3))
	fmt.Println(lowBit(4))
	fmt.Println(lowBit(5))
}
