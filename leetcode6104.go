package myleetcode

func countAsterisks(s string) int {
	flag := true
	result := 0
	for i := range s {
		if s[i] == '|' {
			flag = !flag
			continue
		}
		if flag {
			if s[i] == '*' {
				result++
			}
		}
	}
	return result
}
