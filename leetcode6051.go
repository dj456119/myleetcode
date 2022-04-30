package myleetcode

func countPrefixes(words []string, s string) int {
	count := 0
	for i := range words {
		if isPrefix([]byte(words[i]), []byte(s)) {
			count++
		}
	}
	return count
}

func isPrefix(s []byte, s1 []byte) bool {
	for i := range s {
		if i >= len(s1) {
			return false
		}
		if s[i] != s1[i] {
			return false
		}
	}
	return true
}
