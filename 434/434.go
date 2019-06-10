package main

import (
	"fmt"
)

func countSegments(s string) int {
	count := 0
	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i++
		} else {
			j := i + 1
			for j < len(s) && s[j] != ' ' {
				j++
			}
			count++
			i = j
		}
	}
	return count
}

func main() {
	s := "Hello, my name is John"
	fmt.Println(countSegments(s))
}
