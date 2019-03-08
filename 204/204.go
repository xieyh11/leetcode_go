package main

import (
	"fmt"
)

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	nums := make([]int, n-2)
	for i := range nums {
		nums[i] = i + 2
	}
	for i, v := range nums {
		if v != 1 {
			for j := i + v; j < len(nums); j += v {
				nums[j] = 1
			}
		}
	}
	count := 0
	for _, num := range nums {
		if num != 1 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countPrimes(10))
}
