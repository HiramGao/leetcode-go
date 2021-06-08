package main

import (
	"fmt"
	"math"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordIdMap := map[string]int{}
	graph := [][]int{}

	addWord := func(word string) int {
		if id, ok := wordIdMap[word]; ok {
			return id
		} else {
			id = len(wordIdMap)
			wordIdMap[word] = id
			graph = append(graph, []int{})
			return id
		}
	}

	addEdge := func(word string) int {
		wordId := addWord(word)
		wordByte := []byte(word)
		for i, b := range wordByte {
			wordByte[i] = '*'
			wordAndStarId := addWord(string(wordByte))
			graph[wordId] = append(graph[wordId], wordAndStarId)
			graph[wordAndStarId] = append(graph[wordAndStarId], wordId)
			wordByte[i] = b
		}
		return wordId
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginWordId := addEdge(beginWord)
	endWordId, ok := wordIdMap[endWord]
	if !ok {
		return 0
	}
	dist := make([]int, len(wordIdMap))
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	q := []int{beginWordId}
	dist[beginWordId] = 0
	for len(q) > 0 {
		wordId := q[0]
		q = q[1:]
		if wordId == endWordId {
			return dist[endWordId]/2 + 1
		}
		for _, nextWordId := range graph[wordId] {
			if dist[nextWordId] == math.MaxInt32 {
				dist[nextWordId] = dist[wordId] + 1
				q = append(q, nextWordId)
			}
		}
	}

	return 0
}

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
