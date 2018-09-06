package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	heap := make([]string, 0)
	eachLoc := strings.Split(path, "/")
	for _, v := range eachLoc {
		switch v {
		case "":
		case ".":
		case "..":
			if len(heap) > 0 {
				heap = heap[:len(heap)-1]
			}
		default:
			heap = append(heap, v)
		}
	}
	if len(heap) == 0 {
		return "/"
	}
	return "/" + strings.Join(heap, "/")
}

func main() {
	str := "/home//foo/"
	fmt.Println(simplifyPath(str))
}
