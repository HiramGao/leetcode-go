package utils

import (
	"strconv"
	"strings"
)

type inputToEveryThing struct {
	input string
}

func NewITE(input string) *inputToEveryThing {
	ite := &inputToEveryThing{input: input}
	return ite
}

func (ite inputToEveryThing) ResolveOne() (result []string) {
	return resolveOneRow(ite.input)
}

func resolveOneRow(input string) (result []string) {
	var tmp strings.Builder
	for l, r := 0, len(input)-1; l <= r; l++ {
		if currByte := input[l]; currByte == '[' {
		} else if currByte == ']' {
			result = append(result, tmp.String())
		} else {
			if currByte == ',' {
				result = append(result, tmp.String())
				tmp.Reset()
			} else if currByte != '"' {
				tmp.WriteByte(currByte)
			}
		}
	}
	return
}

func (ite inputToEveryThing) ToInt() []int {
	return toInt(ite.ResolveOne())
}

func toInt(strArr []string) []int {
	result := make([]int, len(strArr))
	for i, v := range strArr {
		result[i], _ = strconv.Atoi(v)
	}
	return result
}

func (ite inputToEveryThing) ToIntInt() [][]int {

	return toIntInt(ite.ResolveMay())
}

func toIntInt(input [][]string) [][]int {
	m := len(input)
	result := make([][]int, m)
	for i := range result {
		result[i] = toInt(input[i])
	}

	return result
}

func (ite inputToEveryThing) ResolveMay() (result [][]string) {
	input := ite.input
	//"[[1,1],[2,2],[3,3]]"
	for l, r := 1, 1; r < len(input)-1; r++ {
		if input[r] == '[' {
			l = r
		} else if input[r] == ']' {
			result = append(result, resolveOneRow(input[l:r+1]))
		}
	}

	return result
}
