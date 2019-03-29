package main

import (
	"fmt"
)

func canWinNim(n int) bool {
	n %= 4
	if n == 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canWinNim(5))
}
