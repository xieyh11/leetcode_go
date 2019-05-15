package main

import (
	"fmt"
)

var Utf8First = []int{0x0, 0xc0, 0xe0, 0xf0}
var Utf8FirstExtra = []int{0x80, 0xe0, 0xf0, 0xf8}

func validOther(data int) bool {
	return (data & 0xc0) == 0x80
}

func getNums(data int) int {
	for i := 0; i < 4; i++ {
		if (data & Utf8FirstExtra[i]) == Utf8First[i] {
			return i + 1
		}
	}
	return -1
}

func validUtf8(data []int) bool {
	for i := 0; i < len(data); {
		nums := getNums(data[i])
		if nums == -1 {
			return false
		}
		j := i + nums
		for k := i + 1; k < j; k++ {
			if k >= len(data) {
				return false
			}
			if !validOther(data[k]) {
				return false
			}
		}
		i = j
	}
	return true
}

func main() {
	data := []int{197, 130, 1}
	fmt.Println(validUtf8(data))

	data = []int{235, 140, 4}
	fmt.Println(validUtf8(data))
}
