package main

import (
	"fmt"
)

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	which := num % 90 / 10
	if which == 0 {
		which = 9
	}
	idx := num % 10
	if which+idx < 10 {
		return which + idx
	} else {
		return which + idx - 9
	}
}

func main() {
	fmt.Println(addDigits(38))
}
