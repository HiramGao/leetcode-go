package main

import "math"

func findNthDigit(n int) int {
	digits := 1
	for {
		num := countOfIntegers(digits)
		if n < num*digits {
			return digitAtIndex(n, digits)
		}
		n -= digits * num
		digits++
	}
}

func digitAtIndex(n int, digits int) int {

	number := beginNumber(digits) + n/digits
	indexFromRight := digits - n%digits
	for i := 1; i < indexFromRight; i++ {
		number /= 10
	}
	return number % 10
}

func beginNumber(digits int) int {
	if digits == 1 {
		return 0
	}
	return int(math.Pow10(digits - 1))
}

func countOfIntegers(digits int) int {
	if digits == 1 {
		return 10
	}
	count := math.Pow10(digits - 1)
	return int(count * 9)
}
func main() {
	println(findNthDigit(1001))
}
