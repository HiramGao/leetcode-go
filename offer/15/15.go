package main

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num = (num - 1) & num
	}
	return count
}

func main() {
	println(hammingWeight(10000000000000000000000000001111))
}
