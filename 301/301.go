package main

import (
	"fmt"
)

func removeOneLeft(s string) map[string]bool {
	res := make(map[string]bool)
	for i := range s {
		if s[i] == '(' {
			tempS := s[:i] + s[i+1:]
			res[tempS] = true
		}
	}
	return res
}

func removeOneLeftMap(sol map[string]bool) map[string]bool {
	res := make(map[string]bool)
	for k := range sol {
		temp := removeOneLeft(k)
		for tempK := range temp {
			res[tempK] = true
		}
	}
	return res
}

func removeOneRight(s string) map[string]bool {
	res := make(map[string]bool)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			tempS := s[:i] + s[i+1:]
			res[tempS] = true
		}
	}
	return res
}

func removeOneRightMap(sol map[string]bool) map[string]bool {
	res := make(map[string]bool)
	for k := range sol {
		temp := removeOneRight(k)
		for tempK := range temp {
			res[tempK] = true
		}
	}
	return res
}

func recReverse(s string, sol map[string]bool) map[string]bool {
	res := make(map[string]bool)
	count := 0
	i := len(s) - 1
	for ; i >= 0; i-- {
		switch s[i] {
		case '(':
			count--
		case ')':
			count++
		}
		if count < 0 {
			break
		}
	}
	if count < 0 {
		suffixRes := removeOneLeft(s[i:])
		if len(sol) == 0 {
			res = suffixRes
		} else {
			for former := range suffixRes {
				for latter := range sol {
					res[former+latter] = true
				}
			}
		}
		suffixRes = removeOneLeftMap(sol)
		for latter := range suffixRes {
			res[s[i:]+latter] = true
		}
		if i > 0 {
			return recReverse(s[:i], res)
		}
	} else {
		if len(sol) == 0 {
			res[s] = true
		} else {
			for latter := range sol {
				res[s+latter] = true
			}
		}
	}
	return res
}

func recSol(s string, sol map[string]bool) map[string]bool {
	res := make(map[string]bool)
	count := 0
	i := 0
	for i = range s {
		switch s[i] {
		case '(':
			count++
		case ')':
			count--
		}
		if count < 0 {
			break
		}
	}
	if count > 0 {
		suffixRes := recReverse(s, res)
		if len(sol) > 0 {
			for former := range sol {
				for latter := range suffixRes {
					res[former+latter] = true
				}
			}
		} else {
			res = suffixRes
		}
	} else if count < 0 {
		prefixRes := removeOneRight(s[:i+1])
		if len(sol) == 0 {
			res = prefixRes
		} else {
			for former := range sol {
				for latter := range prefixRes {
					res[former+latter] = true
				}
			}
		}
		prefixRes = removeOneRightMap(sol)
		for former := range prefixRes {
			res[former+s[:i+1]] = true
		}
		if i < len(s)-1 {
			return recSol(s[i+1:], res)
		}
	} else {
		if len(sol) == 0 {
			res[s] = true
		} else {
			for former := range sol {
				res[former+s] = true
			}
		}
	}
	return res
}

func removeInvalidParentheses(s string) []string {
	sol := make(map[string]bool)
	sol = recSol(s, sol)
	if len(sol) == 0 {
		return []string{""}
	} else {
		res := make([]string, 0)
		for k := range sol {
			res = append(res, k)
		}
		return res
	}
}

func main() {
	s := ")("
	fmt.Println(removeInvalidParentheses(s))
}
