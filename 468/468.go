package main

import (
	"fmt"
	"strconv"
)

func isHex(a byte) bool {
	return (a >= '0' && a <= '9') || (a >= 'a' && a <= 'f') || (a >= 'A' && a <= 'F')
}

func ipv6Test(ip string) string {
	count := 0
	left := 0
	ipLen := len(ip)
	if ip[ipLen-1] == ':' {
		return "Neither"
	}
	for left < ipLen {
		count++
		j := left + 1
		for j < ipLen && ip[j] != ':' {
			j++
		}
		if j-left == 0 || j-left > 4 {
			return "Neither"
		}
		for i := left; i < j; i++ {
			if !isHex(ip[i]) {
				return "Neither"
			}
		}
		left = j + 1
	}
	if count != 8 {
		return "Neither"
	}
	return "IPv6"
}

func ipv4Test(ip string) string {
	count := 0
	left := 0
	ipLen := len(ip)
	if ip[ipLen-1] == '.' {
		return "Neither"
	}
	for left < ipLen {
		count++
		j := left + 1
		for j < ipLen && ip[j] != '.' {
			j++
		}
		num, err := strconv.ParseInt(ip[left:j], 10, 64)
		if err != nil || num > 255 || num < 0 {
			return "Neither"
		}
		if num != 0 {
			if ip[left] == '0' {
				return "Neither"
			}
		} else {
			if ip[left] == '-' || j-left != 1 {
				return "Neither"
			}
		}
		left = j + 1
	}
	if count != 4 {
		return "Neither"
	}
	return "IPv4"
}

func validIPAddress(IP string) string {
	if len(IP) < 5 {
		return "Neither"
	}
	for i := 0; i < 5; i++ {
		if IP[i] == '.' {
			return ipv4Test(IP)
		} else if IP[i] == ':' {
			return ipv6Test(IP)
		}
	}
	return "Neither"
}

func main() {
	s := "172.16.254.1"
	fmt.Println(validIPAddress(s))

	s = "2001:0db8:85a3:0:0:8A2E:0370:7334"
	fmt.Println(validIPAddress(s))

	s = "2001:0db8:85a3:0:0:8A2E:0370:7334:"
	fmt.Println(validIPAddress(s))
}
