package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)

	for s := 0; s < len(words); {
		i := s
		count := 0
		for ; i < len(words); i++ {
			count += len(words[i])
			if i != s {
				count++
			}
			if count > maxWidth {
				break
			}
		}
		if i < len(words) {
			numOfWords := i - s
			if numOfWords > 1 {
				count -= len(words[i])
				count--
				intervals := numOfWords - 1
				extraSpace := maxWidth - count
				eachSpace := extraSpace / intervals
				extraSpace %= intervals
				spacesInterval := strings.Repeat(" ", eachSpace+1)
				tempRes := ""
				for j := s; j < i; j++ {
					if j != s {
						tempRes += spacesInterval
						if extraSpace > 0 {
							tempRes += " "
							extraSpace--
						}
					}
					tempRes += words[j]
				}
				res = append(res, tempRes)
			} else {
				spacesInterval := strings.Repeat(" ", maxWidth-len(words[s]))
				res = append(res, words[s]+spacesInterval)
			}
		} else {
			tempRes := ""
			for j := s; j < i; j++ {
				if j != s {
					tempRes += " "
				}
				tempRes += words[j]
			}
			spacesInterval := strings.Repeat(" ", maxWidth-len(tempRes))
			tempRes += spacesInterval
			res = append(res, tempRes)
		}

		s = i
	}

	return res
}

func main() {
	words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
		"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth := 20
	for _, v := range fullJustify(words, maxWidth) {
		fmt.Println(v)
	}
}
