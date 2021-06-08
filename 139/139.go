package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)

	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	fmt.Println(dp)

	return dp[len(s)]
}
func main() {
	s := "leet"
	w := []string{"leet"}
	fmt.Println(wordBreak(s, w))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat", "og"}))
}
