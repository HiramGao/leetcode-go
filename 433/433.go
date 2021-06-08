package main

import (
	"fmt"
	"math"
)

var apl = []byte{'A', 'C', 'G', 'T'}

func minMutation(start string, end string, bank []string) int {
	result := math.MaxInt32
	visited := map[string]bool{}
	for _, w := range bank {
		visited[w] = false
	}

	var dfs func(s string, depth int)

	dfs = func(s string, depth int) {
		if _, ok := visited[s]; ok && s == end {
			result = min(result, depth)
			return
		}
		if visited[s] {
			return
		}
		visited[s] = true
		depth++
		byteS := []byte(s)
		for i, oldB := range byteS {
			for _, newB := range apl {
				if newB != oldB {
					byteS[i] = newB
					newS := string(byteS)
					if _, ok := visited[newS]; ok {
						dfs(newS, depth)
					}
					byteS[i] = oldB
				}
			}
		}
		visited[s] = false
	}

	dfs(start, 0)

	if result == math.MaxInt32 {
		return -1
	}
	return result
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

func minMutation1(start string, end string, bank []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, ok := wordId[word]
		if ok {
			return id
		}
		id = len(wordId)
		wordId[word] = id
		graph = append(graph, []int{})
		return id
	}
	addEdge := func(word string) int {
		wId := addWord(word)
		wordByte := []byte(word)
		for i, b := range wordByte {
			wordByte[i] = '*'
			wordStarId := addWord(string(wordByte))
			graph[wId] = append(graph[wId], wordStarId)
			graph[wordStarId] = append(graph[wordStarId], wId)
			wordByte[i] = b
		}
		return wId
	}
	for _, w := range bank {
		addEdge(w)
	}
	beginId := addEdge(start)
	endId, ok := wordId[end]
	if !ok {
		return -1
	}

	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[beginId] = 0
	q := []int{beginId}

	for len(q) > 0 {
		wId := q[0]
		q = q[1:]
		if wId == endId {
			return dist[endId] / 2
		}
		for _, w := range graph[wId] {
			if dist[w] == math.MaxInt32 {
				dist[w] = dist[wId] + 1
				q = append(q, w)
			}
		}
	}

	return -1

}
func main() {
	//fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
	//fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
	//fmt.Println(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}))
	//fmt.Println(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCA"}))
	//
	//fmt.Println(minMutation1("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
	//fmt.Println(minMutation1("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
	//fmt.Println(minMutation1("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}))
	fmt.Println(minMutation1("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCA"}))
}
