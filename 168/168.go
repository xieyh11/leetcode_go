package main

import (
	"fmt"
)

func convertToTitle(n int) string {
	n--
	if n < 0 {
		return ""
	}
	res := ""
	for n > 0 {
		temp := byte('A')
		temp += byte(n % 26)
		res = string(temp) + res
		n /= 26
		n--
	}
	if n == 0 {
		res = "A" + res
	}
	return res
}

func main() {
	fmt.Println(convertToTitle(701))
}
