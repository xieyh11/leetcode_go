package main

import (
	"fmt"
)

func halfFind(nums []int, num int) (int, bool) {
	if len(nums) == 0 {
		return -1, false
	}
	if len(nums) == 1 {
		if num < nums[0] {
			return -1, false
		}
		return 0, num == nums[0]
	}
	if len(nums) == 2 {
		if num < nums[0] {
			return -1, false
		} else if num < nums[1] {
			return 0, num == nums[0]
		}
		return 1, num == nums[1]
	}

	mid := len(nums) / 2
	if num == nums[mid] {
		return mid, true
	} else if num < nums[mid] {
		return halfFind(nums[:mid], num)
	} else {
		temp, tempFound := halfFind(nums[mid+1:], num)
		return temp + mid + 1, tempFound
	}
}

func splitInts(nums []int, num int) ([]int, []int) {
	idx, _ := halfFind(nums, num)
	return nums[:idx+1], nums[idx+1:]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func recSol(letters map[byte][]int, k int) int {
	findLess := false
	lessLetter := byte(0)
	count := 0
	for l, v := range letters {
		if len(v) < k {
			lessLetter = l
			findLess = true
			break
		} else {
			count += len(v)
		}
	}
	if !findLess {
		return count
	}
	lessLoc := letters[lessLetter]
	res := 0
	for i := range lessLoc {
		left := make(map[byte][]int)
		right := make(map[byte][]int)
		for k, v := range letters {
			if k != lessLetter {
				splitLeft, splitRight := splitInts(v, lessLoc[i])
				if len(splitLeft) != 0 {
					left[k] = splitLeft
				}
				if len(splitRight) != 0 {
					right[k] = splitRight
				}
			}
		}
		letters = right
		res = getMax(res, recSol(left, k))
	}
	res = getMax(res, recSol(letters, k))
	return res
}

func longestSubstring(s string, k int) int {
	letters := make(map[byte][]int)
	for i := range s {
		idx := s[i] - 'a'
		letters[idx] = append(letters[idx], i)
	}
	return recSol(letters, k)
}

func main() {
	s := "aaabb"
	k := 3
	fmt.Println(longestSubstring(s, k))

	s = "ababbc"
	k = 2
	fmt.Println(longestSubstring(s, k))

	s = ""
	k = 1
	fmt.Println(longestSubstring(s, k))

	s = "abc"
	k = 1
	fmt.Println(longestSubstring(s, k))

	s = "abc"
	k = 2
	fmt.Println(longestSubstring(s, k))
}
