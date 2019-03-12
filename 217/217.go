package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	unique := make(map[int]bool)
	for _, num := range nums {
		if _, ok := unique[num]; ok {
			return true
		} else {
			unique[num] = true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}
