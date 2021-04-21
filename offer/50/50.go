package main

func firstUniqChar(s string) byte {
	if s == "" {
		return ' '
	}
	var hashmap [26]int

	for _, v := range s {
		hashmap[v-'a'] += 1
	}
	for _, v := range s {
		if hashmap[v-'a'] == 1 {
			return byte(v)
		}
	}
	return ' '
}
func main() {
	println(firstUniqChar("abaccdeff"))
}
