package main

import (
	"fmt"
)

func romanToInt(s string) int {
	symValue := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	num := 0
	for i := len(s) - 1; i >= 0; i-- {
		if i == len(s)-1 {
			num = num + symValue[s[i]]
		} else {
			if symValue[s[i]] >= symValue[s[i+1]] {
				num = num + symValue[s[i]]
			} else {
				num = num - symValue[s[i]]
			}
		}
	}
	return num
}

func main() {
	s := "MCMXCIV"

	fmt.Println(romanToInt(s))
}
