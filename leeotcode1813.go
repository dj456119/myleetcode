package myleetcode

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	z1 := strings.Split(sentence1, " ")
	z2 := strings.Split(sentence2, " ")
	if z1[0] != z2[0] && z1[len(z1)-1] != z2[len(z2)-1] {
		return false
	}
	z3 := z2
	z4 := z1
	if len(z1) > len(z2) {
		z3 = z1
		z4 = z2
	}
	count1 := 0
	for i := range z4 {
		if z3[i] != z4[i] {
			break
		}
		count1++
	}
	j := len(z3) - 1
	q := count1
	for i := len(z4) - 1; i >= q; i-- {
		if z3[j] != z4[i] {
			break
		}
		count1++
		j--
	}
	if count1 < len(z4) {
		return false
	}
	if count1 <= len(z3) {
		return true
	}
	return false
}
