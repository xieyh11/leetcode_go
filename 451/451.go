package main

import (
	"fmt"
	"strings"
)

type Freq struct {
	Char  byte
	Count int
}

func frequencySort(s string) string {
	letterFreq := make([]Freq, 0)
	letterMap := make(map[byte]int)
	for i := range s {
		if v, ok := letterMap[s[i]]; ok {
			letterFreq[v].Count++
			for v > 0 && letterFreq[v].Count > letterFreq[v-1].Count {
				letterMap[letterFreq[v].Char], letterMap[letterFreq[v-1].Char] = v-1, v
				letterFreq[v], letterFreq[v-1] = letterFreq[v-1], letterFreq[v]
				v--
			}
		} else {
			letterFreq = append(letterFreq, Freq{s[i], 1})
			letterMap[s[i]] = len(letterFreq) - 1
		}
	}
	res := ""
	for i := range letterFreq {
		res += strings.Repeat(string(letterFreq[i].Char), letterFreq[i].Count)
	}
	return res
}

func main() {
	s := "tree"
	fmt.Println(frequencySort(s))
}
