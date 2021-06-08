package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type wordDict struct {
	word  string
	count int
}
type hp []wordDict

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	x, y := h[i], h[j]
	return x.count < y.count || x.count == y.count && x.word > y.word
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(wordDict)) }
func (h *hp) Pop() (v interface{}) {
	a := *h
	*h, v = a[:len(a)-1], a[len(a)-1]
	return
}
func topKFrequent(words []string, k int) []string {
	hashMap := map[string]int{}

	for _, word := range words {
		hashMap[word]++
	}
	h := &hp{}
	heap.Init(h)
	for word, count := range hashMap {
		heap.Push(h, wordDict{word, count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {

		res[i] = heap.Pop(h).(wordDict).word
	}
	return res
}

func topKFrequent1(words []string, k int) []string {
	hashMap := map[string]int{}

	for _, word := range words {
		hashMap[word]++
	}
	uniqueWords := make([]string, 0, len(hashMap))
	for word := range hashMap {
		uniqueWords = append(uniqueWords, word)
	}
	sort.Slice(uniqueWords, func(i, j int) bool {
		w1, w2 := uniqueWords[i], uniqueWords[j]
		return hashMap[w1] > hashMap[w2] || hashMap[w1] == hashMap[w2] && w1 < w2
	})
	return uniqueWords[:k]
}
func main() {
	words := []string{"rmrypv", "zgsedk", "jlmetsplg", "wnfbo", "hnnftqf", "bxlr", "sviavwoxss", "jdbgvc", "zddeno", "rxgw", "hnnftqf", "hdmvplne", "rlmdt", "jlmetsplg", "ous", "rmrypv", "fwxulnpit", "dmhuq", "rxgw", "ledleb", "bxlr", "indbvb", "ckqqibnx", "cub", "ijww", "ehd", "hfhlfqzkcn", "sviavwoxss", "rxgw", "bxjxpfp", "mgqj", "oic", "ptk", "fwxulnpit", "ijww", "sviavwoxss", "bgfvfa", "zfkgsudxq", "alkq", "dmhuq", "zddeno", "rxgw", "zgsedk", "amarxpg", "bgfvfa", "wnfbo", "sviavwoxss", "sviavwoxss", "alkq", "nmclxk", "zgsedk", "bwowfvira", "ous", "bxlr", "zddeno", "rxgw", "ous", "wnfbo", "rmrypv", "sviavwoxss", "ehd", "zgsedk", "jlmetsplg", "abxvhyehv", "ledleb", "wnfbo", "bgfvfa", "bwowfvira", "amarxpg", "wnfbo", "bwowfvira", "dmhuq", "ouy", "bxlr", "rxgw", "oic", "hnnftqf", "ledleb", "rlmdt", "oldainprua", "ous", "ckqqibnx", "dmhuq", "hnnftqf", "oic", "jlmetsplg", "ppn", "amarxpg", "jlgfgwb", "bxlr", "bwowfvira", "hdmvplne", "oic", "ledleb", "bwowfvira", "oic", "ptk", "dmhuq", "abxvhyehv", "ckqqibnx", "indbvb", "ypzfk", "rmrypv", "bxjxpfp", "amarxpg", "dmhuq", "sviavwoxss", "bwowfvira", "zfkgsudxq", "wnfbo", "rxgw", "jxkvrmajri", "cub", "abxvhyehv", "bwowfvira", "indbvb", "ehd", "ckqqibnx", "oic", "ptk", "hnnftqf", "ouy", "oic", "zgsedk", "abxvhyehv", "ypzfk", "ptk", "sviavwoxss", "rmrypv", "oic", "ous", "abxvhyehv", "hnnftqf", "hfhlfqzkcn", "ledleb", "cub", "ppn", "zddeno", "indbvb", "oic", "jlmetsplg", "ouy", "bwowfvira", "bklij", "gucayxp", "zfkgsudxq", "hfhlfqzkcn", "zddeno", "ledleb", "zfkgsudxq", "hnnftqf", "bgfvfa", "jlmetsplg", "indbvb", "ehd", "wnfbo", "hnnftqf", "rlmdt", "bxlr", "indbvb", "jdbgvc", "jlmetsplg", "cub", "jlgfgwb", "ypzfk", "indbvb", "dmhuq", "jlmetsplg", "zgsedk", "rmrypv", "cub", "rxgw", "ledleb", "rxgw", "sviavwoxss", "ckqqibnx", "hdmvplne", "dmhuq", "wnfbo", "jlmetsplg", "bxlr", "zfkgsudxq", "bxjxpfp", "ledleb", "indbvb", "ckqqibnx", "ous", "ckqqibnx", "cub", "ous", "abxvhyehv", "bxlr", "hfhlfqzkcn", "hfhlfqzkcn", "oic", "ten", "amarxpg", "indbvb", "cub", "alkq", "alkq", "sviavwoxss", "indbvb", "bwowfvira", "ledleb"}
	k := 41
	fmt.Println(topKFrequent(words, k))
	fmt.Println(topKFrequent1(words, k))
}
