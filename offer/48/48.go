package main

func lengthOfLongestSubstring(s string) int {

	res := 0
	hashmap := make(map[int32]int)
	tmp := 0
	for i, v := range s {
		if i == 0 {
			tmp = 1
			hashmap[v] = 0
		} else if closeIndex, ok := hashmap[v]; !ok {
			tmp++
		} else if tmp < i-closeIndex {
			tmp++
		} else {
			tmp = i - closeIndex
		}
		hashmap[v] = i
		res = max(res, tmp)
	}
	return res
}

func max(res int, tmp int) int {
	if res > tmp {
		return res
	}
	return tmp
}
func main() {
	println(lengthOfLongestSubstring("abcabcd"))
}
