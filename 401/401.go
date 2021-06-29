package main

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {
	res := []string{}
	for i := uint8(0); i < 12; i++ {
		for j := uint8(0); j < 60; j++ {
			if bits.OnesCount8(i)+bits.OnesCount8(j) == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return res
}
