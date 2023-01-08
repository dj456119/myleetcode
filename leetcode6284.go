package myleetcode

func isItPossible(word1 string, word2 string) bool {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	m3 := make(map[byte]int)
	m4 := make(map[byte]int)
	for i := range word1 {
		m1[word1[i]]++
		m3[word1[i]]++
	}
	for i := range word2 {
		m2[word2[i]]++
		m4[word2[i]]++
	}
	for i := range m3 {
		for j := range m4 {
			m1[i]--
			m2[j]--
			if m1[i] <= 0 {
				delete(m1, i)
			}
			if m2[j] <= 0 {
				delete(m2, j)
			}
			m1[j]++
			m2[i]++
			if len(m1) == len(m2) {
				return true
			}
			m1[j]--
			m2[i]--
			if m1[j] <= 0 {
				delete(m1, j)
			}
			if m2[i] <= 0 {
				delete(m2, i)
			}
			m1[i]++
			m2[j]++
		}
	}
	return false
}
