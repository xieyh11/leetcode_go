package main

import (
	"fmt"
)

func recSol(n int, ahead int) []int {
	res := []int{}
	for i := 0; i < 10; i++ {
		temp := ahead*10 + i
		if temp <= n {
			res = append(res, temp)
			res = append(res, recSol(n, temp)...)
		} else {
			break
		}
	}
	return res
}

func lexicalOrder(n int) []int {
	res := []int{}
	for i := 1; i < 10; i++ {
		if i <= n {
			res = append(res, i)
			res = append(res, recSol(n, i)...)
		} else {
			break
		}
	}
	return res
}

func main() {
	fmt.Println(lexicalOrder(20))
}
