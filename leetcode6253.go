package myleetcode

import "strings"

func isCircularSentence(sentence string) bool {
	r := strings.Split(sentence, " ")
	for i := range r {
		if i == len(r)-1 {
			last := len(r[i]) - 1
			if r[i][last] != r[0][0] {
				return false
			}
		} else {
			last := len(r[i]) - 1
			if r[i][last] != r[i+1][0] {
				return false
			}
		}
	}
	return true
}
