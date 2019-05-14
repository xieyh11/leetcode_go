package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	letters := [26]int{}
	for i := range magazine {
		letters[magazine[i]-'a']++
	}
	for i := range ransomNote {
		idx := ransomNote[i] - 'a'
		if letters[idx] == 0 {
			return false
		}
		letters[idx]--
	}
	return true
}

func main() {
	str1 := "a"
	str2 := "b"
	fmt.Println(canConstruct(str1, str2))

	str1 = "aa"
	str2 = "ab"
	fmt.Println(canConstruct(str1, str2))

	str1 = "aa"
	str2 = "aab"
	fmt.Println(canConstruct(str1, str2))

	str1 = "ab"
	str2 = "ba"
	fmt.Println(canConstruct(str1, str2))
}
