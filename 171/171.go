package main

import (
	"fmt"
)

func titleToNumber(s string) int {
	res := 0
	for i := range s {
		res *= 26
		res += int(s[i]-'A') + 1
	}
	return res
}

func main() {
	s := "A"
	fmt.Println(titleToNumber(s))
}
