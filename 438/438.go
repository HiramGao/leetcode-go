package main

import "fmt"

func findAnagrams(s string, p string) []int {
	sl := len(s)
	pl := len(p)
	if pl > sl {
		return []int{}
	}
	res := []int{}
	hashMap := make(map[byte]int)
	for _, v := range []byte(p) {
		hashMap[v]++
	}
	var prev byte
	fmt.Println(hashMap)
	for l, r := 0, 0; r < sl; l, r = l+1, r+1 {
		if r == 0 {
			for i := 0; i < pl; i++ {
				if _, ok := hashMap[s[i]]; ok {
					hashMap[s[i]]--
				}
			}
			r = pl - 1
			prev = s[0]
		} else {
			if _, ok := hashMap[prev]; ok {
				hashMap[prev]++
			}
			if _, ok := hashMap[s[r]]; ok {
				hashMap[s[r]]--
			}
			prev = s[l]
		}
		fmt.Println(l, r, hashMap)

		flag := 0
		for _, v := range hashMap {
			if v != 0 {
				flag = 1
				break
			}
		}
		if flag == 0 {
			res = append(res, l)
		}

	}
	return res
}

func findAnagrams1(s string, p string) []int {
	sl := len(s)
	pl := len(p)
	if pl > sl {
		return []int{}
	}
	res := []int{}
	sArr := [26]int{}
	pArr := [26]int{}
	for i, v := range []byte(p) {
		pArr[v-'a']++
		sArr[s[i]-'a']++
	}
	l, r := 0, pl-1
	sArr[s[r]-'a']--
	for r < sl {
		sArr[s[r]-'a']++
		if sArr == pArr {
			res = append(res, l)
		}
		sArr[s[l]-'a']--
		l++
		r++
	}

	return res

}
func main() {
	fmt.Println(findAnagrams1("aa", "bb"))
}
