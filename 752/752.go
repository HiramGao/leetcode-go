package main

import "fmt"

const start = "0000"

type pair struct {
	current string
	step    int
}

func nextStatus(status string) (next []string) {
	s := []byte(status)
	for i, b := range s {
		s[i] = b - 1
		if s[i] < '0' {
			s[i] = '9'
		}
		next = append(next, string(s))

		s[i] = b + 1
		if s[i] > '9' {
			s[i] = '0'
		}
		next = append(next, string(s))
		s[i] = b
	}
	return
}

func openLock(deadends []string, target string) int {
	if target == start {
		return 0
	}
	hm := map[string]bool{}
	for _, deadend := range deadends {
		hm[deadend] = true
	}
	if hm[start] {
		return -1
	}
	visited := map[string]bool{start: true}
	q := []pair{{start, 0}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for _, next := range nextStatus(cur.current) {
			if !visited[next] && !hm[next] {
				if next == target {
					return cur.step + 1
				}
				visited[next] = true
				q = append(q, pair{next, cur.step + 1})
			}
		}

	}

	return -1
}

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"))
}
