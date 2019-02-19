package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	if len(ratings) == 1 {
		return 1
	}
	res := make([]int, len(ratings))
	for startIdx := 1; startIdx < len(ratings); {
		if ratings[startIdx] < ratings[startIdx-1] {
			i := startIdx
			for i < len(ratings) && ratings[i] < ratings[i-1] {
				i++
			}
			if i == len(ratings)-1 && ratings[i] < ratings[i-1] {
				for j := i - 1; j >= startIdx; j-- {
					res[j] = res[j+1] + 1
				}
			} else {
				for j := i - 2; j >= startIdx; j-- {
					res[j] = res[j+1] + 1
				}
			}
			res[startIdx-1] = max(res[startIdx-1], res[startIdx]+1)
			startIdx = i
		} else if ratings[startIdx] == ratings[startIdx-1] {
			startIdx++
		} else {
			res[startIdx] = res[startIdx-1] + 1
			startIdx++
		}
	}
	count := len(ratings)
	for _, v := range res {
		count += v
	}
	return count
}

func main() {
	nums := []int{1, 0, 2}
	fmt.Println(candy(nums))
}
