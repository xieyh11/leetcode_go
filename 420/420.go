package main

import (
	"fmt"
)

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func strongPasswordChecker(s string) int {
	lower := false
	upper := false
	digit := false
	cond2 := 0
	for i := range s {
		if !lower && s[i] >= 'a' && s[i] <= 'z' {
			lower = true
			cond2++
		}
		if !upper && s[i] >= 'A' && s[i] <= 'Z' {
			upper = true
			cond2++
		}
		if !digit && s[i] >= '0' && s[i] <= '9' {
			digit = true
			cond2++
		}
	}
	if len(s) < 6 {
		return getMax(3-cond2, 6-len(s))
	}
	over3 := []int{}
	for i := 0; i < len(s); {
		j := i + 1
		for ; j < len(s) && s[i] == s[j]; j++ {
		}
		if j-i >= 3 {
			over3 = append(over3, j-i)
		}
		i = j
	}
	res := 0
	left := 0
	over := getMax(len(s)-20, 0)
	for k := 1; k < 3; k++ {
		for i := 0; i < len(over3) && over > 0; i++ {
			if over3[i] >= 3 && over3[i]%3 == k-1 {
				over3[i] -= k
				over -= k
				res += k
			}
		}
	}
	for i := 0; i < len(over3); i++ {
		if over3[i] >= 3 && over > 0 {
			temp := over3[i] - 2
			if temp <= over {
				over3[i] = 2
				over -= temp
				res += temp
			} else {
				over3[i] -= over
				res += over
				over = 0
			}
		}
		if over3[i] >= 3 {
			left += over3[i] / 3
		}
	}
	if over > 0 {
		res += over
	}
	res += getMax(3-cond2, left)
	return res
}

func main() {
	s := "dkfj3jjjjjjjjjjjj99llkjkj99999999Ajjjjjjjkkkk"
	fmt.Println(strongPasswordChecker(s))

	s = "1234567890123456Baaaa"
	fmt.Println(strongPasswordChecker(s))
}
