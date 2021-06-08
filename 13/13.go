package main

import "fmt"

func romanToInt(s string) int {
	valueSymbols := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	res := 0
	b := []byte(s)

	for i := 0; i < len(b); i++ {
		currStr := string(b[i])
		if i+1 < len(b) {
			currStr += string(b[i+1])
			if v, ok := valueSymbols[currStr]; ok {
				res += v
				i++
			} else {
				res += valueSymbols[string(b[i])]
			}
		} else {
			res += valueSymbols[currStr]
		}
	}

	return res
}

func romanToInt1(s string) int {
	valueSymbols := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	res := 0
	n := len(s)

	for i := 0; i < len(s); i++ {
		curr := valueSymbols[s[i]]
		if i+1 < n && curr < valueSymbols[s[i+1]] {
			res -= curr
		} else {
			res += curr
		}
	}

	return res
}
func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))

	fmt.Println(romanToInt1("III"))
	fmt.Println(romanToInt1("IV"))
	fmt.Println(romanToInt1("IX"))
	fmt.Println(romanToInt1("LVIII"))
	fmt.Println(romanToInt1("MCMXCIV"))
}
