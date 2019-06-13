package main

import (
	"fmt"
	"strconv"
)

func prefixCount(nStr string, prefixStr string) int {
	digitLeft := len(nStr) - len(prefixStr)
	numCount := 10
	count := 1
	for i := 1; i < digitLeft; i++ {
		count += numCount
		numCount *= 10
	}
	if digitLeft > 0 && nStr[:len(prefixStr)] > prefixStr {
		count += numCount
	} else if digitLeft > 0 && nStr[:len(prefixStr)] == prefixStr {
		nStrLeft := nStr[len(prefixStr):]
		num, _ := strconv.ParseInt(nStrLeft, 10, 64)
		count += int(num) + 1
	}
	return count
}

func recSol(n int, nStr string, k int, prefix int, prefixStr string) int {
	if k == 1 {
		return prefix
	}
	if prefix <= n/10 {
		count := 1
		for i := 0; i < 10; i++ {
			temp := prefix*10 + i
			if temp <= n {
				tempStr := prefixStr + string(byte(i)+'0')
				tempCount := prefixCount(nStr, tempStr)
				if count+tempCount >= k {
					return recSol(n, nStr, k-count, temp, tempStr)
				}
				count += tempCount
			}
		}
	}
	return 0
}

func findKthNumber(n int, k int) int {
	count := 0
	nStr := strconv.FormatInt(int64(n), 10)
	for i := 1; i < 10; i++ {
		if i <= n {
			tempStr := string(byte(i) + '0')
			tempCount := prefixCount(nStr, tempStr)
			if count+tempCount >= k {
				return recSol(n, nStr, k-count, i, tempStr)
			}
			count += tempCount
		}
	}
	return 0
}

func main() {
	fmt.Println(findKthNumber(100, 12))
	n := 100
	for k := 1; k <= n; k++ {
		fmt.Println(findKthNumber(n, k))
	}
}
