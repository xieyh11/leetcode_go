package main

import (
	"fmt"
)

// func wordSubIndex(s, t string) []bool {
// 	sLen, tLen := len(s), len(t)
// 	if sLen == 0 || sLen > tLen {
// 		return []bool{}
// 	}
// 	if sLen == tLen {
// 		return []bool{s == t}
// 	}
// 	res := make([]bool, tLen)
// 	nowScan := []int{}
// 	tNow := 0
// 	for tNow < tLen {
// 		nextSacn := []int{}
// 		for _, scan := range nowScan {
// 			if s[scan] == t[tNow] {
// 				scan++
// 				if scan == sLen {
// 					res[tNow-sLen+1] = true
// 				} else {
// 					nextSacn = append(nextSacn, scan)
// 				}
// 			}
// 		}
// 		if t[tNow] == s[0] {
// 			if sLen > 1 {
// 				if tLen-tNow >= sLen {
// 					nextSacn = append(nextSacn, 1)
// 				}
// 			} else {
// 				res[tNow] = true
// 			}
// 		}
// 		nowScan = nextSacn
// 		tNow++
// 	}
// 	return res
// }

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

func (this *Trie) recSol(word string, first bool) bool {
	nowAt := this.Root
	res := false
	wordLen := len(word)
	for i := range word {
		idx := word[i] - 'a'
		if nowAt != nil && nowAt.Children[idx] != nil {
			nowAt = nowAt.Children[idx]
			if nowAt.IsWord {
				if i+1 < wordLen {
					res = this.recSol(word[i+1:], false)
				} else if !first {
					res = true
				}
			}
		} else {
			break
		}
		if res {
			break
		}
	}
	return res
}

func findAllConcatenatedWordsInADict(words []string) []string {
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}
	res := []string{}
	for _, word := range words {
		if trie.recSol(word, true) {
			res = append(res, word)
		}
	}
	return res
}

func main() {
	words := []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}
	fmt.Println(findAllConcatenatedWordsInADict(words))

	words = []string{""}
	fmt.Println(findAllConcatenatedWordsInADict(words))
}
