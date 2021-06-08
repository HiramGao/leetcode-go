package main

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this

	for _, char := range word {
		char -= 'a'
		if node.children[char] == nil {
			node.children[char] = &Trie{}
		}
		node = node.children[char]
	}

	node.isEnd = true
}

func (this *Trie) SearchPrefix(word string) *Trie {
	node := this
	for _, char := range word {
		char -= 'a'
		if node.children[char] == nil {
			return nil
		}
		node = node.children[char]
	}
	return node
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	trie.Search("apple")   // 返回 True
	trie.Search("app")     // 返回 False
	trie.StartsWith("app") // 返回 True
	trie.Insert("app")
	trie.Search("app")

}
