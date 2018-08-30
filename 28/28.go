package main

import (
	"container/list"
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	idxList := list.New()
	for i := range haystack {
		removeIdx := make([]*list.Element, 0)
		for j := idxList.Front(); j != nil; j = j.Next() {
			idxValue := j.Value.(int)
			needleIdx := i - idxValue
			if needleIdx < len(needle) && haystack[i] != needle[needleIdx] {
				removeIdx = append(removeIdx, j)
			} else if needleIdx >= len(needle) {
				return idxValue
			}
		}
		for _, temp := range removeIdx {
			idxList.Remove(temp)
		}
		if needle[0] == haystack[i] {
			idxList.PushBack(i)
		}
	}
	if idxList.Len() > 0 {
		res := idxList.Front().Value.(int)
		if len(haystack)-res == len(needle) {
			return res
		} else {
			return -1
		}
	} else {
		return -1
	}
}

func main() {
	haystack := "aa"
	needle := "aaa"
	fmt.Println(strStr(haystack, needle))
}
