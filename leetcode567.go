package myleetcode

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	m1 := stringToArr(s1)
	rs := []rune(s2)
	m2 := stringToArr(s2[0:len(s1)])
	for j := len(s1); j < len(s2); j++ {
		if compareArr(m1, m2) {
			return true
		}
		m2[rs[j-len(s1)]-'a']--
		m2[rs[j]-'a']++
	}

	return compareArr(m1, m2)
}

func compareArr(m1, m2 []int) bool {
	for i, v := range m1 {
		if m2[i] != v {
			return false
		}
	}
	return true
}

func stringToArr(s1 string) []int {
	m1 := make([]int, 26)
	for _, c := range s1 {
		m1[c-'a']++
	}
	return m1
}

func CheckInclusion(s1 string, s2 string) bool {
	return checkInclusion(s1, s2)
}
