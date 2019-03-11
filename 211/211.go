package main

import (
	"fmt"
)

type TrieNode struct {
	IsWord      bool
	ContainWord bool
	Children    [26]*TrieNode
}

type WordDictionary struct {
	Root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{&TrieNode{}}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
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

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.Root.Search([]byte(word))
}

func (this *TrieNode) Search(word []byte) bool {
	if len(word) == 0 {
		return this.IsWord
	}
	if word[0] != '.' {
		idx := word[0] - 'a'
		if this.Children[idx] == nil {
			return false
		}
		return this.Children[idx].Search(word[1:])
	} else if this.ContainWord {
		isWord := false
		for i := 0; i < 26 && !isWord; i++ {
			if this.Children[i] != nil {
				isWord = this.Children[i].Search(word[1:])
			}
		}
		return isWord
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func main() {
	dic := Constructor()
	dic.AddWord("bad")
	dic.AddWord("dad")
	dic.AddWord("mad")
	fmt.Println(dic.Search("pad"))
	fmt.Println(dic.Search("bad"))
	fmt.Println(dic.Search(".ad"))
	fmt.Println(dic.Search("b.."))
}
