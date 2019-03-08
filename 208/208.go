package main

import (
	"fmt"
)

type TrieNode struct {
	IsWord      bool
	ContainWord bool
	Children    [26]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{&TrieNode{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	nowAt := this.Root
	for i := range word {
		nowAt.ContainWord = true
		idx := word[i] - 'a'
		if nowAt.Children[idx] == nil {
			nowAt.Children[idx] = new(TrieNode)
		}
		nowAt = nowAt.Children[idx]
	}
	nowAt.IsWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	nowAt := this.Root
	for i := range word {
		idx := word[i] - 'a'
		if nowAt == nil || nowAt.Children[idx] == nil {
			return false
		}
		nowAt = nowAt.Children[idx]
	}
	return nowAt.IsWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	nowAt := this.Root
	for i := range prefix {
		idx := prefix[i] - 'a'
		if nowAt == nil || nowAt.Children[idx] == nil {
			return false
		}
		nowAt = nowAt.Children[idx]
	}
	return nowAt.ContainWord || nowAt.IsWord
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor()

	trie.Insert("a")
	fmt.Println(trie.Search("a"))     // returns true
	fmt.Println(trie.Search("ab"))    // returns false
	fmt.Println(trie.StartsWith("a")) // returns true
	trie.Insert("ab")
	fmt.Println(trie.Search("ab")) // returns true
}
