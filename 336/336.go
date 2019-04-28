package main

import (
	"fmt"
)

type TrieNode struct {
	IsWord      bool
	Index       int
	ContainWord bool
	Children    [26]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

func Constructor() Trie {
	return Trie{&TrieNode{}}
}

func (this *Trie) Insert(word string, idx int) {
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
	nowAt.Index = idx
}

func (this *Trie) InsertReverse(word string, idx int) {
	nowAt := this.Root
	for i := len(word) - 1; i >= 0; i-- {
		nowAt.ContainWord = true
		idx := word[i] - 'a'
		if nowAt.Children[idx] == nil {
			nowAt.Children[idx] = new(TrieNode)
		}
		nowAt = nowAt.Children[idx]
	}
	nowAt.IsWord = true
	nowAt.Index = idx
}

func getSuffix(node *TrieNode) []SuffixIndex {
	if node == nil || !node.ContainWord {
		return []SuffixIndex{}
	}
	res := make([]SuffixIndex, 0)
	for i := byte(0); i < 26; i++ {
		prefix := string(i + 'a')
		if node.Children[i] != nil && node.Children[i].IsWord {
			res = append(res, SuffixIndex{prefix, node.Children[i].Index})
		}
		tempRes := getSuffix(node.Children[i])
		for _, s := range tempRes {
			s.S = prefix + s.S
			res = append(res, s)
		}
	}
	return res
}

type SuffixIndex struct {
	S     string
	Index int
}

func isPalin(s string) bool {
	if len(s) <= 1 {
		return true
	}
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func recSol(o, r *TrieNode) ([][]int, []SuffixIndex, []SuffixIndex) {
	res := make([][]int, 0)
	oSuffix := make([]SuffixIndex, 0)
	rSuffix := make([]SuffixIndex, 0)
	if o != nil && o.ContainWord && r != nil && r.ContainWord {
		for i := byte(0); i < 26; i++ {
			tempRes, oTemp, rTemp := recSol(o.Children[i], r.Children[i])
			res = append(res, tempRes...)
			prefix := string(i + 'a')
			if o.Children[i] != nil && o.Children[i].IsWord {
				oSuffix = append(oSuffix, SuffixIndex{prefix, o.Children[i].Index})
			}
			if r.Children[i] != nil && r.Children[i].IsWord {
				rSuffix = append(rSuffix, SuffixIndex{prefix, r.Children[i].Index})
			}
			for _, s := range oTemp {
				s.S = prefix + s.S
				oSuffix = append(oSuffix, s)
			}
			for _, s := range rTemp {
				s.S = prefix + s.S
				rSuffix = append(rSuffix, s)
			}
		}
	} else if o != nil && o.ContainWord {
		oSuffix = append(oSuffix, getSuffix(o)...)
	} else if r != nil && r.ContainWord {
		rSuffix = append(rSuffix, getSuffix(r)...)
	}

	if o != nil && o.IsWord && r != nil && r.IsWord {
		if o.Index != r.Index {
			res = append(res, []int{o.Index, r.Index})
		}
	}
	if o != nil && o.IsWord {
		for _, rS := range rSuffix {
			if rS.Index != o.Index && isPalin(rS.S) {
				res = append(res, []int{o.Index, rS.Index})
			}
		}
	}
	if r != nil && r.IsWord {
		for _, oS := range oSuffix {
			if oS.Index != r.Index && isPalin(oS.S) {
				res = append(res, []int{oS.Index, r.Index})
			}
		}
	}
	return res, oSuffix, rSuffix
}

func palindromeSearch(origin, reverse Trie) [][]int {
	res, _, _ := recSol(origin.Root, reverse.Root)
	return res
}

func palindromePairs(words []string) [][]int {
	origin := Constructor()
	reverse := Constructor()
	for i, word := range words {
		origin.Insert(word, i)
		reverse.InsertReverse(word, i)
	}
	return palindromeSearch(origin, reverse)
}

func main() {
	strs0 := []string{"abcd", "dcba", "lls", "s", "sssll"}
	fmt.Println(palindromePairs(strs0))

	strs1 := []string{"bat", "tab", "cat"}
	fmt.Println(palindromePairs(strs1))
}
