package main

import "strings"

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	sLen := len(s)
	if sLen == 0 {
		return false
	}
	//IsE
	EIndex := strings.IndexByte(s, 'E')
	if EIndex == -1 {
		EIndex = strings.IndexByte(s, 'e')
	}
	if EIndex >= 0 {
		return isE(s, EIndex)
	}

	return isInteger(s, false) || isFloat(s)
}

func isE(s string, EIndex int) bool {
	if EIndex == 0 {
		return false
	}
	if EIndex == len(s)-1 {
		return false
	}
	return (isInteger(s[0:EIndex], false) || isFloat(s[0:EIndex])) && isInteger(s[EIndex+1:], false)
}

func isFloat(s string) bool {
	if len(s) <= 1 {
		return false
	}
	dotIndex := strings.IndexByte(s, '.')
	if dotIndex == 0 {
		return allIsNumber(s[1:])
	}
	if dotIndex == len(s)-1 {
		return isInteger(s[0:dotIndex], false)
	}
	if dotIndex == -1 {
		return false
	} else {
		return isInteger(s[0:dotIndex], true) && allIsNumber(s[dotIndex+1:])
	}
}

func isInteger(s string, isFloor bool) bool {
	signIndex := strings.IndexByte(s, '-')
	if signIndex == -1 {
		signIndex = strings.IndexByte(s, '+')
	}
	if signIndex > 0 {
		return false
	}
	if len(s) == 1 && signIndex == 0 {
		return isFloor
	}
	if signIndex == 0 {
		return allIsNumber(s[1:])
	} else {
		return allIsNumber(s)
	}
}

func allIsNumber(s string) bool {
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

func main() {
	println(isNumber("+.8"))
}
