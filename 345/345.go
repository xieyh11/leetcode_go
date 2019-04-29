package main

import (
	"fmt"
)

func isVowel(a byte) bool {
	return a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u' || a == 'A' || a == 'E' || a == 'I' || a == 'O' || a == 'U'
}

func reverseVowels(s string) string {
	if len(s) <= 1 {
		return s
	}
	sB := []byte(s)
	i, j := 0, len(sB)-1
	for i < j {
		if !isVowel(sB[i]) {
			i++
		} else if !isVowel(sB[j]) {
			j--
		} else {
			sB[i], sB[j] = sB[j], sB[i]
			i++
			j--
		}
	}
	return string(sB)
}

func main() {
	str0 := "hello"
	fmt.Println(reverseVowels(str0) == "holle")

	str1 := "leetcode"
	fmt.Println(reverseVowels(str1) == "leotcede")

	str2 := "aA"
	fmt.Println(reverseVowels(str2) == "Aa")
}
