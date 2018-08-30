package main

import (
	"fmt"
	"sort"
)

type bytes []byte

func (a bytes) Len() int           { return len(a) }
func (a bytes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bytes) Less(i, j int) bool { return a[i] < a[j] }

func sortString(s string) string {
	sbytes := bytes(s)
	sort.Sort(sbytes)
	return string(sbytes)
}

func groupAnagrams(strs []string) [][]string {
	sToIdx := make(map[string]int)
	res := make([][]string, 0)
	for _, str := range strs {
		soStr := sortString(str)
		if i, ok := sToIdx[soStr]; ok {
			res[i] = append(res[i], str)
		} else {
			sToIdx[soStr] = len(res)
			res = append(res, []string{str})
		}
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
