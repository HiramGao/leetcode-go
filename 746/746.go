package main

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	pre, cur := 0, 0

	for i := 2; i <= n; i++ {
		pre, cur = cur, min(cur+cost[i-1], pre+cost[i-2])
	}
	return cur
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}
func main() {
	println(minCostClimbingStairs([]int{10, 15, 20}))
}
