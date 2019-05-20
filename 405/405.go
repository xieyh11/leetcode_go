package main

import (
	"fmt"
)

var Hex = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}

func toHex(num int) string {
	res := ""
	for i := uint(0); i < 32; i += 4 {
		res = Hex[num&0xf] + res
		num >>= 4
	}
	start := 0
	for start < len(res)-1 && res[start] == '0' {
		start++
	}
	return res[start:]
}

func main() {
	fmt.Println(toHex(26))
	fmt.Println(toHex(-1))
}
