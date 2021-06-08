package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	for i := 0; i < 32 && num > 0; i++ {
		res |= num & 1 << (31 - i)
		num = num >> 1
	}
	return res
}
func main() {
	fmt.Println(reverseBits(00000010100101000001111010011100))
}
