package main

func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			dp[i] = 0
		} else {
			dp[i] = dp[i-1] + 1
			res = max(res, dp[i])
		}
	}
	return res
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {

}
