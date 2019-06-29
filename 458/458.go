package main

import (
	"fmt"
)

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	if buckets == 1 {
		return 0
	}
	times := minutesToTest / minutesToDie
	if times == 0 {
		return 0
	}
	combination := []int{1}
	ans := 0
	solveBuckets := make([][]int, times)
	numOfPigs := 0
	for ans < buckets {
		numOfPigs++
		solveBuckets[0] = append(solveBuckets[0], 1<<uint(numOfPigs))
		for i := len(combination) - 1; i > 0; i-- {
			combination[i] += combination[i-1]
		}
		combination = append(combination, 1)
		for i := 1; i < times; i++ {
			res := 0
			for j := 0; j < numOfPigs; j++ {
				res += solveBuckets[i-1][numOfPigs-1-j] * combination[j]
			}
			res++
			solveBuckets[i] = append(solveBuckets[i], res)
		}
		ans = solveBuckets[times-1][numOfPigs-1]
	}
	return numOfPigs
}

func main() {
	fmt.Println(poorPigs(1000, 15, 30))
	fmt.Println(poorPigs(1, 1, 1))
}
