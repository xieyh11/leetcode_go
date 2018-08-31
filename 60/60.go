package main

import (
	"fmt"
)

var nfactor = [...]int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
var choice = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func getPermutation(n int, k int) string {
	if n < 1 || n > 9 {
		return ""
	}
	res := make([]byte, 0)
	used := make([]bool, n)
	for i := n; i >= 1; i-- {
		temp := (k - 1) / nfactor[i-1]
		k -= temp * nfactor[i-1]
		count := 0
		j := 0
		for ; j < n; j++ {
			if !used[j] {
				count++
			}
			if count > temp {
				break
			}
		}
		res = append(res, choice[j])
		used[j] = true
	}
	return string(res)
}

func main() {
	for k := 1; k < 7; k++ {
		fmt.Println(getPermutation(3, k))
	}
}
