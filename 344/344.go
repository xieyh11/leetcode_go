package main

import (
	"fmt"
)

func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}
	i := 0
	j := len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func main() {
	s0 := []byte("hello")
	reverseString(s0)
	fmt.Println(string(s0))
}
