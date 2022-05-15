package myleetcode

func removeAnagrams(words []string) []string {
	result := []string{}
	deletearr := make([]bool, len(words))
	for i := range words {
		if i == len(words)-1 {
			continue
		}

		if dubi([]byte(words[i]), []byte(words[i+1])) {
			deletearr[i+1] = true
		}

	}
	for i := range words {
		if deletearr[i] == false {
			result = append(result, words[i])
		}
	}
	return result
}

func dubi(s1, s2 []byte) bool {
	z1 := make([]int, 26)
	z2 := make([]int, 26)
	for _, b1 := range s1 {
		z1[b1-'a']++
	}
	for _, b2 := range s2 {
		z2[b2-'a']++
	}
	for i := range z1 {
		if z1[i] != z2[i] {
			return false
		}
	}
	return true
}
