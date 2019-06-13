package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	res := 0
	left := 0
	for left < len(chars) {
		i := left + 1
		for ; i < len(chars) && chars[i] == chars[left]; i++ {
		}
		count := i - left
		if res < left {
			chars[res] = chars[left]
		}
		res++
		if count > 1 {
			countBytes := []byte(strconv.FormatInt(int64(count), 10))
			copy(chars[res:res+len(countBytes)], countBytes)
			res += len(countBytes)
		}
		left = i
	}
	return res
}

func main() {
	chars := []byte("aabbccc")
	res := compress(chars)
	fmt.Println(res, string(chars[:res]))

	chars = []byte("abbbbbbbbbbbb")
	res = compress(chars)
	fmt.Println(res, string(chars[:res]))
}
