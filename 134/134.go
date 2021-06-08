package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {

	n := len(gas)
	start := -1
	var next func(i, prev int) bool

	next = func(i, prev int) bool {
		i %= n
		if i == start {
			return true
		}
		prev = prev + gas[i] - cost[i]
		if prev < 0 {
			return false
		}

		if next(i+1, prev) {
			return true
		}
		return false
	}

	for i, g := range gas {
		c := cost[i]
		if g >= c {
			start = i
			if next(i+1, g-c) {
				return i
			}
			start = -1
		}

	}
	return start
}

func canCompleteCircuit1(gas []int, cost []int) int {

	for i, n := 0, len(gas); i < n; {
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}
	return -1
}
func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
	fmt.Println(canCompleteCircuit1(gas, cost))
	gas = []int{2, 3, 4}
	cost = []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas, cost))
	fmt.Println(canCompleteCircuit1(gas, cost))
}
