package main

import "fmt"

var letters = [][]string{[]string{"a", "b", "c"},
	[]string{"d", "e", "f"},
	[]string{"g", "h", "i"},
	[]string{"j", "k", "l"},
	[]string{"m", "n", "o"},
	[]string{"p", "q", "r", "s"},
	[]string{"t", "u", "v"},
	[]string{"w", "x", "y", "z"}}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	lettersOne := letters[digits[0]-'2']
	if len(digits) == 1 {
		return lettersOne
	}
	subRes := letterCombinations(digits[1:])
	lenSubRes := len(subRes)
	res := make([]string, len(lettersOne)*lenSubRes)
	for i, letter := range lettersOne {
		for j, subStr := range subRes {
			res[i*lenSubRes+j] = letter + subStr
		}
	}

	return res
}

func main() {
	str := "23"
	fmt.Println(letterCombinations(str))
}
