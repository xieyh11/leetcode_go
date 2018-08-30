package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	res := ""

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		allEqual := true
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || c != strs[j][i] {
				allEqual = false
				break
			}
		}
		if allEqual {
			res = res + string(c)
		} else {
			break
		}
	}
	return res
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
