package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func countOfAtoms(formula string) string {
	i, n := 0, len(formula)

	parseAtom := func() string {
		left := i
		i++
		for i < n && unicode.IsLower(rune(formula[i])) {
			i++
		}
		return formula[left:i]
	}
	parseNum := func() (ans int) {
		if i == n || !unicode.IsDigit(rune(formula[i])) {
			return 1
		}
		for ; i < n && unicode.IsDigit(rune(formula[i])); i++ {
			ans = ans*10 + int(formula[i]-'0')
		}
		return
	}
	hm := []map[string]int{{}}
	for i < n {
		if ch := formula[i]; ch == '(' {
			i++
			hm = append(hm, map[string]int{})
		} else if ch == ')' {
			i++
			num := parseNum()
			currentAtom := hm[len(hm)-1]
			hm = hm[:len(hm)-1]
			for k, v := range currentAtom {
				hm[len(hm)-1][k] += num * v
			}
		} else {
			atom := parseAtom()
			num := parseNum()
			hm[len(hm)-1][atom] += num
		}
	}

	type pair struct {
		atom string
		num  int
	}

	pairs := []pair{}

	for k, v := range hm[0] {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].atom < pairs[j].atom })
	var ans strings.Builder

	for _, p := range pairs {
		ans.WriteString(p.atom)
		if p.num > 1 {
			ans.WriteString(strconv.Itoa(p.num))
		}
	}

	return ans.String()
}

func main() {
	fmt.Println(countOfAtoms("H2O"))
	fmt.Println(countOfAtoms("Mg2(OH)2"))
	fmt.Println(countOfAtoms("K4(ON(SO3)2)2"))
	fmt.Println(countOfAtoms("Be32"))
}
