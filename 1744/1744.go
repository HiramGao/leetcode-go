package main

func canEat(candiesCount []int, queries [][]int) []bool {
	dp := make([]int, len(candiesCount)+1)
	dp[0] = 0
	for i := 1; i <= len(candiesCount); i++ {
		dp[i] = dp[i-1] + candiesCount[i-1]
	}
	result := make([]bool, len(queries))
	for i, query := range queries {
		favoriteType, favoriteDay, dailyCap := query[0]+1, query[1]+1, query[2]
		low := dp[favoriteType]
		high := dp[favoriteType-1] / dailyCap
		result[i] = favoriteDay <= low && favoriteDay > high
	}
	return result
}
func main() {

}
