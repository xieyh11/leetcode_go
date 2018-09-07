package main

import (
	"fmt"
)

func firstWindow(s string, tMap map[byte]int) (start, length int, currentMap map[byte]int) {
	currentMap = make(map[byte]int)
	for k := range tMap {
		currentMap[k] = 0
	}
	count := 0
	size := len(tMap)
	i := 0
	for ; i < len(s); i++ {
		if _, ok := currentMap[s[i]]; ok {
			break
		}
	}
	start = i
	for ; i < len(s) && count < size; i++ {
		if _, ok := currentMap[s[i]]; ok {
			currentMap[s[i]]++
			if currentMap[s[i]] == tMap[s[i]] {
				count++
			}
		}
	}
	if count == size {
		for ; start < i; start++ {
			if _, ok := tMap[s[start]]; ok {
				if currentMap[s[start]] > tMap[s[start]] {
					currentMap[s[start]]--
				} else {
					break
				}
			}
		}
		length = i - start
	} else {
		length = 0
	}
	return
}

func minWindow(s string, t string) string {
	numOfChar := make(map[byte]int)
	for i := range t {
		if _, ok := numOfChar[t[i]]; ok {
			numOfChar[t[i]]++
		} else {
			numOfChar[t[i]] = 1
		}
	}
	minI, minLen, currentMap := firstWindow(s, numOfChar)
	if minLen == 0 {
		return ""
	}
	tempI, tempLen := minI, minLen
	for i := minI + minLen; i < len(s); i++ {
		if _, ok := numOfChar[s[i]]; ok {
			currentMap[s[i]]++
			for ; tempI <= i; tempI++ {
				if _, ok := numOfChar[s[tempI]]; ok {
					if currentMap[s[tempI]] > numOfChar[s[tempI]] {
						currentMap[s[tempI]]--
					} else {
						break
					}
				}
			}
			tempLen = i - tempI + 1
			if tempLen < minLen {
				minLen = tempLen
				minI = tempI
			}
		}
	}
	return s[minI:(minI + minLen)]
}

func main() {
	s := "abcabdebac"
	t := "cda"

	fmt.Println(minWindow(s, t))
}
