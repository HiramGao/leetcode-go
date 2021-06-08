package main

func findMaxLength(nums []int) int {
	mp := map[int]int{0: -1}
	counter := 0
	res := 0
	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}
		if prevIndex, ok := mp[counter]; ok {
			res = max(res, i-prevIndex)
		} else {
			mp[counter] = i
		}
	}
	return res
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {

}
