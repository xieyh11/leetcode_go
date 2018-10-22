package main

import (
	"fmt"
	"strconv"
	"strings"
)

func recurSol(s string, remain int) [][]string {
	if remain == 1 {
		if s[0] == '0' {
			if len(s) > 1 {
				return [][]string{}
			}
		} else {
			num, _ := strconv.ParseInt(s, 10, 64)
			if num >= 0 && num <= 255 {
				return [][]string{{s}}
			} else {
				return [][]string{}
			}
		}
	}
	if len(s) < remain {
		return [][]string{}
	}
	if len(s) == remain {
		res := make([]string, remain)
		for i := range s {
			res[i] = string(s[i])
		}
		return [][]string{res}
	}
	res := make([][]string, 0)
	for i := 1; i < len(s); i++ {
		num, _ := strconv.ParseInt(s[:i], 10, 64)
		if num > 255 {
			break
		}
		tempR := recurSol(s[i:], remain-1)
		for _, oneR := range tempR {
			res = append(res, append([]string{s[:i]}, oneR...))
		}
		if num == 0 {
			break
		}
	}
	return res
}

func restoreIpAddresses(s string) []string {
	tempR := recurSol(s, 4)
	res := make([]string, 0)
	for _, oneR := range tempR {
		res = append(res, strings.Join(oneR, "."))
	}
	return res
}

func main() {
	str := "010010"
	fmt.Println(restoreIpAddresses(str))
}
