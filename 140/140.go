package main

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	wordDictSet := make(map[string]bool)
	n := len(s)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([][][]string, n)
	var backtrack func(i int) [][]string
	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}
		var wordList [][]string
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if _, ok := wordDictSet[word]; ok {
				for _, newWords := range backtrack(i) {
					wordList = append(wordList, append([]string{word}, newWords...))
				}
			}
		}
		word := s[index:]
		if _, ok := wordDictSet[word]; ok {
			wordList = append(wordList, []string{word})
		}
		dp[index] = wordList
		return wordList
	}

	result := []string{}
	for _, words := range backtrack(0) {
		result = append(result, strings.Join(words, " "))
	}

	return result
}
func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(wordBreak(s, wordDict))

	s = "pineapplepenapple"
	wordDict = []string{"apple", "pen", "applepen", "pine", "pineapple"}
	fmt.Println(wordBreak(s, wordDict))

	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))
}
