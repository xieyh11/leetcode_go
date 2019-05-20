package main

import (
	"fmt"
)

func remainWithLength(num string, k int) string {
	if len(num) == k {
		return num
	}
	res := ""
	start := 0
	for k > 0 {
		min := num[start]
		minIdx := start
		for i := start + 1; i <= len(num)-k; i++ {
			if num[i] < min {
				min = num[i]
				minIdx = i
			}
		}
		res += string(min)
		start = minIdx + 1
		k--
	}
	return res
}

func removeKdigits(num string, k int) string {
	if len(num) <= k {
		return "0"
	}
	if k == 0 {
		return num
	}
	beforeZero := 0
	for i := 0; i < len(num); i++ {
		if num[i] == '0' {
			break
		}
		beforeZero++
	}
	if beforeZero > k || beforeZero == len(num) {
		return remainWithLength(num, len(num)-k)
	} else {
		start := beforeZero
		for start < len(num)-1 && num[start] == '0' {
			start++
		}
		return removeKdigits(num[start:], k-beforeZero)
	}
}

func main() {
	num := "1432219"
	k := 3
	fmt.Println(removeKdigits(num, k))

	num = "10200"
	k = 1
	fmt.Println(removeKdigits(num, k))
}
