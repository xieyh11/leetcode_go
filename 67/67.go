package main

import (
	"fmt"
)

func addThreeBits(a, b, c byte) (res, carry byte) {
	switch {
	case a == '0' && b == '0' && c == '0':
		res = '0'
		carry = '0'
	case a == '0' && b == '0' && c == '1':
		fallthrough
	case a == '0' && b == '1' && c == '0':
		fallthrough
	case a == '1' && b == '0' && c == '0':
		res = '1'
		carry = '0'
	case a == '0' && b == '1' && c == '1':
		fallthrough
	case a == '1' && b == '0' && c == '1':
		fallthrough
	case a == '1' && b == '1' && c == '0':
		res = '0'
		carry = '1'
	case a == '1' && b == '1' && c == '1':
		res = '1'
		carry = '1'
	}
	return
}

func addTwoBits(a, b byte) (res, carry byte) {
	if a == b {
		res = '0'
	} else {
		res = '1'
	}
	if a == '1' && b == '1' {
		carry = '1'
	} else {
		carry = '0'
	}
	return
}

func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	maxL := len(a)
	if maxL < len(b) {
		maxL = len(b)
	}
	res := make([]byte, maxL)
	var carry byte = '0'
	ai, bi, resi := len(a)-1, len(b)-1, maxL-1
	for ai >= 0 && bi >= 0 {
		res[resi], carry = addThreeBits(a[ai], b[bi], carry)
		ai--
		bi--
		resi--
	}
	for ai >= 0 {
		res[resi], carry = addTwoBits(a[ai], carry)
		ai--
		resi--
	}
	for bi >= 0 {
		res[resi], carry = addTwoBits(b[bi], carry)
		bi--
		resi--
	}
	if carry != '0' {
		res = append([]byte{carry}, res...)
	}
	return string(res)
}

func main() {
	a := "1"
	b := "111"
	fmt.Println(addBinary(a, b))
}
