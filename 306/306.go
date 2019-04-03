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

func stringAdd(a, b string) string {
	if len(a) == 0 || len(b) == 0 {
		return ""
	}
	lena, lenb := len(a), len(b)
	res := make([]byte, getMax(lena, lenb))
	ai, bi, resi := lena-1, lenb-1, len(res)-1
	carry := byte(0)
	for ai >= 0 && bi >= 0 {
		res[resi] = a[ai] - '0' + b[bi] + carry
		carry = 0
		if res[resi] > '9' {
			res[resi] -= 10
			carry = 1
		}
		resi--
		ai--
		bi--
	}
	for ai >= 0 {
		res[resi] = a[ai] + carry
		carry = 0
		if res[resi] > '9' {
			res[resi] -= 10
			carry = 1
		}
		resi--
		ai--
	}
	for bi >= 0 {
		res[resi] = b[bi] + carry
		carry = 0
		if res[resi] > '9' {
			res[resi] -= 10
			carry = 1
		}
		resi--
		bi--
	}
	if carry == 0 {
		return string(res)
	} else {
		return string(carry+'0') + string(res)
	}
}

func isRest(a, b, num string) bool {
	for len(num) > 0 {
		temp := stringAdd(a, b)
		if len(num) < len(temp) || num[:len(temp)] != temp {
			return false
		}
		a, b = b, temp
		num = num[len(temp):]
	}
	return true
}

func isAdditiveNumber(num string) bool {
	if len(num) == 0 {
		return false
	}
	if num[0] == '0' {
		for j := 1; j < len(num)-1; j++ {
			if isRest("0", num[1:j+1], num[j+1:]) {
				return true
			}
		}
		return false
	}
	for i := range num {
		if i+1 == len(num) {
			break
		}
		if num[i+1] == '0' {
			if i+2 < len(num) && isRest(num[:i+1], "0", num[i+2:]) {
				return true
			}
		} else {
			for j := i + 1; j < len(num)-1; j++ {
				a := num[:i+1]
				b := num[i+1 : j+1]
				if isRest(a, b, num[j+1:]) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	num := "101"
	fmt.Println(isAdditiveNumber(num))
}
